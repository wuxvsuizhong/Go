package main

import (
	"fmt"
	"sort"
)

//深度搜索在n个数的数组中找到m个数，这m个数字的和等于targat
func main() {
	target := 0
	fmt.Scanln(&target)
	nums := []int{}
	tmpn := 0
	for {
		if n, _ := fmt.Scanf("%d", &tmpn); n != 0 {
			nums = append(nums, tmpn)
		} else {
			break
		}
	}

	fmt.Println(nums)
	sort.Ints(nums)
	rec := make([]bool, len(nums))

	var search func(int, int, *[]int) bool
	search = func(target int, pos int, nums *[]int) bool {
		if target == 0 {
			return true
		}
		for i := pos; i < len(*nums); i++ {
			if !rec[i] {
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
	fmt.Println(search(target, 0, &nums))

}
