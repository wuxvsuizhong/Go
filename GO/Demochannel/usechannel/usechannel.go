package usechannel

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Start() {
	//声明一个元素为int类型的通道
	var b chan int
	fmt.Println(b)     //未初始化的通道打印是nil
	b = make(chan int) //不带缓冲区的通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("从通道中取到元素:", x)
	}()

	b <- 100 //往不带缓冲区的通道中发送元素,会死锁

	wg.Wait()

}

func Start2() {
	var c chan int
	c = make(chan int, 1) //带缓冲区的通道初始化
	fmt.Println(c)
	//初始化后得到通道的指针
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-c
		fmt.Println("从通道中取到元素:", x)
	}()

	c <- 11 //往带有缓冲区的通道中发送元素，不会发生死锁
	c <- 22
	// c <- 33  //超过通道缓存容量，会发生死锁
	close(c) //关闭通道
	wg.Wait()
}
