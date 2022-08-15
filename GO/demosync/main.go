package main

import (
	"fmt"
	"strconv"
	"sync"
)

//golang 自带的map不是并发安全的
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, val int) {
	m[key] = val
}

func main1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		//超过cpu最大运行线程会容易报错fatal error: concurrent map writes
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=%v,v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// ==============================

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)       //向sync.map中写入键值对
			val, _ := m2.Load(key) //从sync.map中获取键值对
			fmt.Printf("k:%v,v:%v\n", key, val)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
