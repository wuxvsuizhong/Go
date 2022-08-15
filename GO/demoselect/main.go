package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int, 10)
	//ch缓存容量不为0的时候，每次循环执行那个select的case是随机的不确定的
	ch := make(chan int, 1)
	//当ch的缓冲容量为1的时候,从0开始的循环偶数次放入数据会被选择执行
	//ch缓冲区容量为1时，循环奇数次的读取数据会被执行
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // ch能存储数据的时候随机执行向ch中发送数据，但是如果其他分支都是失败的那么该case会被执行，因为case至少得有一个被选择
			fmt.Println(x)
		case ch <- i: //在ch不能存储数据的时候不会走该case
		}
	}
}
