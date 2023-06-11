package goroutineErrProcess

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello...")
		time.Sleep(time.Second)
	}
}

func test() {
	/*
		如果在主routine中捕获异常，那么发生panic后会导致整个程序崩溃，
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误：", err)
		}
	}()

	/*
		  defer+recover组合对于异常的捕获是以routinr为单位的，也就是如果某个goroutine里面发生了panic，那么直接在routine中捕获处理该异常
			而不影响其他的routine
	*/

	var mymap map[int]string
	//手动创建一个panic错误
	mymap[0] = "golng" //没有make就使用map引发panic
}

func ErrorExample() {
	go sayHello()
	go test()

	for i := 0; i < 20; i++ {
		fmt.Println("主线程中循环等待...")
		time.Sleep(time.Second)
	}
}
