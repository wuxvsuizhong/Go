package usechannel

import (
	"fmt"
	"time"
)

/*
使用channel实现异步的数据读写，并能自动判断读写数据的结束
*/
func WriteCh(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("向管道写入：", i)
	}
	close(ch) //同一个channel，关闭channel并不会影响其他routine对channel的读取
}

func ReadCh(rch chan int, exitch chan bool) {
	for {
		if v, ok := <-rch; ok { //如果注释掉读的动作，只向管道中写而没有读的动作，在写入超过管道容量后，会deadlock
			fmt.Println("读取数据：", v)
			time.Sleep(time.Millisecond * 100)
			//time.Sleep(time.Second) //如果在有写也有读的情况下，及时写入的动作很快，读取的动作很慢也不会影响，只要能边读边写，就没有问题，和读写的速度无关
		} else {
			break //读取到的ok标志位false时退出对ch的读取
		}
	}
	exitch <- true //向exitch中发送读任务结束的标识
	close(exitch)
}

func TwoChanneReadWrite() {
	ch := make(chan int, 10)
	exitch := make(chan bool, 1)

	go WriteCh(ch)
	go ReadCh(ch, exitch)

	for {
		_, ok := <-exitch //阻塞等待任务完成标识
		if !ok {
			fmt.Println("任务完成!")
			break
		}
	}
}
