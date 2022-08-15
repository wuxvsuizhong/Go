package workpool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine 和channel 实现计算int64 各个位数的数字的加和的程序
// 1.开启一个goroutine 循环生成int64数字，发送到jobchan
// 2.开启24个goroutine 从jobchn中随机取出数字计算各个位数的加和，并发送到resultchan
// 3.主goroutine从resultchan中取出结果并打印输出

type goods struct {
	x int64
}

type result struct {
	g   *goods
	ret int64
}

func producer(jobchan chan<- *goods) {
	//循环生成int64并发送到jobchan
	wg.Done() //不会执行，因为for死循环
	for {
		jobchan <- &goods{
			x: rand.Int63(),
		}
		time.Sleep(time.Millisecond * 500) //休眠500毫秒
	}
}

func consumer(jobchan <-chan *goods, resultchan chan<- *result) {
	//从jobchan 中不断取出元素并处理计算
	wg.Done() //不会执行，因为for死循环
	for {
		g := <-jobchan
		sum := int64(0)
		n := g.x
		for n > 0 {
			sum += n % 10
			n /= 10
		}

		calcret := &result{
			g:   g,
			ret: sum,
		}

		resultchan <- calcret
	}
}

var wg sync.WaitGroup

func Start() {
	jobchan := make(chan *goods, 100)
	resultchan := make(chan *result, 100)

	wg.Add(1)
	go producer(jobchan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go consumer(jobchan, resultchan)
	}

	for ret := range resultchan {
		fmt.Printf("value:%d,sum:%d\n", ret.g.x, ret.ret)
	}

	wg.Wait()
}
