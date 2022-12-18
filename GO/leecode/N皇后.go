package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)

	square := make([][]bool, n)
	for i := 0; i < n; i++ {
		square[i] = make([]bool, n)
	}
	book := make([]int, n)

	cnt := 0
	var dfs func(x int)
	dfs = func(x int) {
		if x+1 > n {
			cnt++
			return
		}

		for j := 0; j < n; j++ {
			if isvalid(x, j, &book) {
				book[x] = j
				dfs(x + 1)
				book[x] = 0
			}
		}
	}

	dfs(0)
	fmt.Println(cnt)
}

func isvalid(x int, y int, book *[]int) bool {
	for i := 0; i < x; i++ {
		if (*book)[i] == y {
			return false
		}
		if x+y == i+(*book)[i] || x-y == i-(*book)[i] {
			return false
		}
	}

	return true
}
