package main

import (
	"fmt"
	"sort"
)

func main() {
	//采用在目标数组中找出n个数字满足target值,还是深度搜索，也容易超时
	var types int
	fmt.Scanln(&types)
	weights := make([]int, types)
	wquantity := make([]int, types)

	for i := 0; i < types; i++ {
		fmt.Scan(&weights[i])
	}
	for i := 0; i < types; i++ {
		fmt.Scan(&wquantity[i])
	}

	fmt.Println(weights, wquantity)
	w := []int{}
	for i := 0; i < len(weights); i++ {
		for j := 0; j < wquantity[i]; j++ {
			w = append(w, weights[i])
		}
	}
	sort.Ints(w)
	fmt.Println(w)

	rec := make([]bool, len(w))

	var search func(int, int, *[]int) bool
	search = func(target int, pos int, nums *[]int) bool {
		if target == 0 {
			return true
		} else if target < 0 {
			return false
		} else if target < (*nums)[0] {
			return false
		}
		for i := pos; i < len(*nums); i++ {
			if !rec[i] {
				if pos > 0 && (*nums)[i] == (*nums)[i-1] && !rec[i-1] {
					continue
				}
				rec[i] = true
				ret := search(target-(*nums)[i], pos+1, nums)
				rec[i] = false
				if ret {
					return true
				}

			}
		}
		return false
	}

	//print(search(186, 0, &w))
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
	}

	cnt := 1
	for i := 1; i <= sum; i++ {
		if search(i, 0, &w) {
			cnt += 1
			//println(i)
		}
	}
	println(cnt)

}
