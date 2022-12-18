package main

import (
	"fmt"
	"sync"
	"time"
)

func print_num(wg *sync.WaitGroup, lock1 *sync.Mutex, lock2 *sync.Mutex) {
	defer wg.Done()
	i := 0
	for {
		lock1.Lock()
		fmt.Printf("-----%d-----\n", i)
		time.Sleep(time.Second * 1)
		lock2.Unlock()
	}
}

func print_alpha(wg *sync.WaitGroup, lock1 *sync.Mutex, lock2 *sync.Mutex) {
	defer wg.Done()
	i := 0
	for {
		lock2.Lock()
		fmt.Printf("-----%s-----\n", string(rune('A'+i)))
		time.Sleep(time.Second * 1)
		lock1.Unlock()
	}
}

func main() {
	var m1 sync.Mutex
	var m2 sync.Mutex
	var wg sync.WaitGroup
	m1.Lock()
	//m2.Lock()

	wg.Add(1)
	go print_num(&wg, &m1, &m2)
	wg.Add(1)
	go print_alpha(&wg, &m1, &m2)

	wg.Wait()
}
