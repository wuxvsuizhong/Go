package DemoMAXPROCS

import (
	"fmt"
	"runtime"

	// "runtime"
	"sync"
)

var sWg sync.WaitGroup

func worker1() {
	defer sWg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("worker1:%d\n", i)
	}
}

func worker2() {
	defer sWg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("worker2:%d\n", i)
	}
}

func Start() {
	runtime.GOMAXPROCS(2) //不设置则默认跑满所有CPU核心
	sWg.Add(2)
	go worker1()
	go worker2()
	sWg.Wait()
}
