package main

import (
	"demogoroutine/DemoMAXPROCS"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Printf("hello!--%d\n", i)
}

func Test1() {
	for i := 0; i < 1000; i++ {
		go hello(i)
	}
}

func Test2() {
	for i := 0; i < 50; i++ {
		go func() { //匿名函数，闭包 ——使用goroutine 并发
			fmt.Println("hello,test2", i)
			// 会出现很多的重复输出，因为打印取的是匿名函数外部的变量；
			// 在输出时可能外循环已经不是刚启动goroutine的时候的外部变量的值了
			// 并行启动打印时候，有可能是几乎同一时间访问外部循环变量，所以会有很多的形同打印
		}()
	}

	fmt.Println("_________________")
	for i := 0; i < 50; i++ {
		go func(j int) { //匿名函数，闭包 ——使用goroutine 并发
			fmt.Println("hello,test2", j)
		}(i) //实时传值进入匿名函数就不会输出重复了，这样就不会访问外部变量
	}
}

func worker(i int) {
	defer wg.Done()                                       //计数减一
	time.Sleep(time.Second * time.Duration(rand.Intn(5))) //随机睡眠1~5 秒
	fmt.Println(i)
}

var wg sync.WaitGroup

func routineWait() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数加1
		go worker(i)
	}

}

func main() {
	// Test1()
	// Test2()
	fmt.Println("main func")
	// time.Sleep(time.Second) //主线程添加延时，防止主线程提前退出

	routineWait()

	wg.Add(1)
	go DemoMAXPROCS.Start()
	wg.Done()

	wg.Wait() //等待计数为0再退出main
}
