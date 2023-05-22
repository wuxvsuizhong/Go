package usechannel

import (
	"fmt"
	// "math/rand"
	"sync"
)

// 1.启动一个goroutine 生成100个数字发送到ch1
// 2.启动另外一个goroutine 从ch1中取数计算平方放到ch2中
// 3.打印ch2中的所有元素

var wg2 sync.WaitGroup

func work1(ch1 chan int64) {
	defer wg2.Done()
	for i := 0; i < 100; i++ {
		ch1 <- int64(i)
	}
	close(ch1)
}

func work2(ch1 chan int64, ch2 chan int64) {
	defer wg2.Done()

	for { //循环读取消费ch1
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func Start3() {
	ch1 := make(chan int64, 100)
	ch2 := make(chan int64, 100)
	wg2.Add(2)
	go work1(ch1)
	go work2(ch1, ch2)
	wg2.Wait()

	for val := range ch2 {
		fmt.Println(val)
	}
}

func Start4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	val, ok := <-ch
	fmt.Println(val, ok)
	val, ok = <-ch
	fmt.Println(val, ok)
	val, ok = <-ch
	fmt.Println(val, ok)

	//对已经关闭的通道是可以取值的，在通道中还有值的时候能取到元素，并返回true的ok标志
	//当已经关闭的通道中没有值的时候，取值不会异常，但是ok标志返回false
}

func Start5() {
	ch := make(chan int, 2)
	ch <- 11
	ch <- 22

	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	/*
		当管道close后，遍历可以正常结束，不会发生死锁
	*/
}
