package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, ch *chan bool) {
	defer (*wg).Done()

	//块定义的标识
FORLOOP:
	for {
		fmt.Println("work running...")
		time.Sleep(500 * time.Millisecond)
		select {
		case x := <-(*ch):
			fmt.Printf("work收到结束消息,%v\n", x)
			break FORLOOP //从通道中接收到消息后就break 跳出FORLOOP 块
		default:
		}
	}
}

func work2(wg *sync.WaitGroup, ctx *context.Context) {
	defer (*wg).Done()
	(*wg).Add(1)
	go childwork2(wg, ctx) //在gorotine中发起子goroutine
FORLOOP:
	for {
		fmt.Println("work2 running...")
		time.Sleep(500 * time.Millisecond)
		select {
		case x := <-(*ctx).Done(): //如果context.Done()有返回则退出work
			fmt.Printf("work2收到结束消息,%#v\n", x)
			break FORLOOP
		default:
		}
	}

}

func childwork2(wg *sync.WaitGroup, ctx *context.Context) {
	defer (*wg).Done()
FORLOOP:
	for {
		fmt.Println("childwork2 running...")
		time.Sleep(200 * time.Millisecond)
		select {
		case x := <-(*ctx).Done(): //如果context.Done()有返回则退出work
			fmt.Printf("childwork2收到结束消息,%#v\n", x)
			break FORLOOP
		default:
		}
	}

}

func main1() {
	var wg sync.WaitGroup
	var ch = make(chan bool)
	wg.Add(1)

	go work(&wg, &ch)
	time.Sleep(5 * time.Second) //主gorouine睡眠5s
	ch <- true                  //主gorouine睡眠5s后向通道ch发送消息
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go work2(&wg, &ctx)
	time.Sleep(5 * time.Second) //主gorouine睡眠5s
	cancel()                    //主gorouine睡眠5s后使用cancel()通知子goroutine逐级退出
	wg.Wait()
}
