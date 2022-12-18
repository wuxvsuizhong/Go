/*
2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点
需要注意的是各种进位问题。
极端情况，例如
Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1
*/

package main

import "fmt"

type Listnode struct {
	Val  int
	Next *Listnode
}

func (l *Listnode) Add(root *Listnode, node *Listnode) {
	head := root
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
}

func (l *Listnode) Travel(root *Listnode) {
	head := root
	for head.Next != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}

func main() {
	l1 := &Listnode{9, nil}
	l2 := &Listnode{1, nil}
	for i := 0; i < 4; i++ {
		l1.Add(l1, &Listnode{9, nil})
	}
	l1.Travel(l1)
	l2.Travel(l2)

	res := addTwoNumbers(l1, l2)
	fmt.Println("----------")
	res.Travel(res)
}

func addTwoNumbers(l1 *Listnode, l2 *Listnode) *Listnode {
	if l1 == nil || l2 == nil {
		return nil
	}
	head := &Listnode{0, nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		current.Next = &Listnode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		carry = (x + y + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		current.Next = &Listnode{Val: carry % 10, Next: nil}
	}
	return head.Next
}
