package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//返回一个chan并发起一个goroutine工作
func gen(ctx context.Context) chan int {
	dst := make(chan int)
	n := 1
	go func() { //发起goroutine 向dst中不断写入
		for {
			select {
			case <-ctx.Done(): //检测到context.Done()有返回后就return 跳出匿名函数(goroutine)
				return
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //cancel() 去通知那些有检测context的goroutine并发程式去退出工作

	nums := gen(ctx)
	//for n := range gen(ctx) {
	for n := range nums {
		//fmt.Printf("channel地址:%p\n", nums)
		fmt.Println(n)
		if n == 10 {
			break //break后，会执行defer的cancel()
		}
	}
}

func main2() {
	d := time.Now().Add(2000 * time.Millisecond)
	// d := time.Now().Add(50 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d) //启用一个context并设置过期时间为d

	defer cancel() //主动发起cancel，及时释放context

	select { //执行一次select，某个case一旦成立就结束select
	case <-time.After(1 * time.Second): //1s后打印
		fmt.Println("working...")
	case <-ctx.Done(): //context的dealine超时时间到了以后Done()后有返回
		fmt.Println(ctx.Err())

	}
}

func connectExample(ctx context.Context, wg *sync.WaitGroup) {
	defer (*wg).Done()
LOOP:
	for {
		fmt.Println("connecting ...")
		time.Sleep(time.Millisecond * 10)
		select {
		case x := <-ctx.Done():
			fmt.Printf("收到结束消息:%v\n", x)
			break LOOP
		default:
		}
	}
	fmt.Println("work done!")
	// wg.Done()
}

func main1() {
	//启动一个context并设置超时时间为100ms
	ctx, cancal := context.WithTimeout(context.Background(), time.Millisecond*100)
	var wg sync.WaitGroup
	wg.Add(1)
	go connectExample(ctx, &wg)
	time.Sleep(time.Second * 5)
	cancal() //通知goroutine结束
	wg.Wait()
	fmt.Println("over")
}
