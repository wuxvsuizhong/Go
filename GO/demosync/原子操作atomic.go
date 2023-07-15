package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var x1 int64
var ax int64

var wg2 sync.WaitGroup
var lock sync.Mutex

func add1() {
	x1 += 1
	wg2.Done()
}

func add() {
	lock.Lock()
	x += 1
	lock.Unlock()
	wg2.Done()
}

func addAtomic() {
	atomic.AddInt64(&ax, 1)
	wg2.Done()
}

func main() {
	wg2.Add(100000)
	for i := 0; i < 100000; i++ {
		go add1()
	}
	wg2.Wait()
	fmt.Println("x1", x1)

	wg2.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg2.Wait()
	fmt.Println("x", x)

	wg2.Add(100000)
	for i := 0; i < 100000; i++ {
		go addAtomic()
	}
	wg2.Wait()
	fmt.Println("ax", ax)
}
