//给定n,用1...n个数字来表达Binary search tree,问有多少种表达方式
//分析：以root节点为界，左子树的组成方式*右子树的组成方式 = 数构成的组合数
//给定n,root 节点可以是从1...n之间每一个数字迭代变化，那么相应的每一次不同的root树的组成方式就是其左右子树组成方式的乘积

package main

import (
	"fmt"
)

func numTress(n int) int {
	if n <= 2 {
		return n
	}
	sol := make([]int, n+1)
	sol[0], sol[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for v := 0; v < i; v++ {
			sol[i] += sol[v] * sol[i-1-v]
		}
	}

	return sol[n]
}

func main() {
	n := []int{1, 3, 5, 6}
	for _, v := range n {
		fmt.Println(numTress(v))
	}
}
