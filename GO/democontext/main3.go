package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//自定义TraceCode 类型
type TraceCode string

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer (*wg).Done()
	key := TraceCode("TRACE_INFO")
	traceCode, ok := ctx.Value(key).(string) //获取context中key值相应的value
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker,trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			fmt.Println("收到结束消息")
			break LOOP
		default:
		}
	}

	fmt.Println("worker done!")
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) //启动一个context并设置超时时间为50ms
	ctx = context.WithValue(ctx, TraceCode("TRACE_INFO"), "123234234")
	wg.Add(1)
	go worker(ctx, &wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
