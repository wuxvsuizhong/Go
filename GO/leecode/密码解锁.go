package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytestr, _, _ := reader.ReadLine()
	deadends := make(map[string]struct{})
	for _, s := range strings.Split(string(bytestr), " ") {
		deadends[s] = struct{}{}
	}

	var target string
	fmt.Scanln(&target)

	if target == "0000" {
		fmt.Println("res=0")
		return
	}

	if _, ok := deadends[target]; ok {
		fmt.Println("res=-1")
		return
	}

	num_inc := func(ch rune) rune {
		if ch == rune('9') {
			return '0'
		} else {
			return ch + 1
		}
	}

	num_dec := func(ch rune) rune {
		if ch == rune('0') {
			return '9'
		} else {
			return ch - 1
		}
	}
	type pair struct {
		//Status []rune   //不要再这里定义切片类型，因为go在传递拷贝切片的时候是浅拷贝
		// 也就是说如果pair作为list的节点push的话，如果是通过同类型的实例去复制，那么只会复制结构体的外层
		// 结构体的元素如果是切片类型，那么最终所有的复制的结构体中，只要元素是切片类型的，这个元素的切片内部的值都会指向同一份

		Status string
		Steps  int
	}
	seen := make(map[string]struct{})
	var q list.List = list.List{}

	q.PushBack(pair{"0000", 0})
	for q.Len() != 0 {
		e := q.Remove(q.Front())
		ev := e.(pair)
		status_num, stps := ev.Status, ev.Steps
		_, ok1 := seen[status_num]
		_, ok2 := deadends[status_num]
		if ok1 || ok2 {
			continue
		}

		seen[status_num] = struct{}{}
		if status_num == target {
			fmt.Println("res=", stps)
			return
		}

		for i := 0; i < 4; i++ {
			snum := rune(status_num[i])
			t_slice := []rune(status_num)
			t_slice[i] = num_inc(snum)
			q.PushBack(pair{Status: string(t_slice), Steps: stps + 1})
			t_slice[i] = num_dec(snum)
			q.PushBack(pair{Status: string(t_slice), Steps: stps + 1})
		}
	}
	fmt.Println("end,res=-1")
}
