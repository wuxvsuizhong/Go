package main

import (
	"fmt"
	"math"
)

func process(n int) []int {
	end := int(math.Sqrt(float64(n)))
	rets := []int{}
	for i := 1; i < end+1; i++ {
		if n%i == 0 {
			rets = append(rets, i)
			if v := n / i; v != i {
				rets = append(rets, v)
			}
		}
		if len(rets) > 4 {
			return nil
		}
	}
	if len(rets) == 4 {
		return rets
	}
	return nil
}

func main() {
	var nums = []int{21, 4, 7}
	results := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		ret := process(nums[i])
		if ret != nil {
			results = append(results, ret...)
		}
	}
	sum := 0
	for _, val := range results {
		sum += val
	}
	fmt.Println(results)
	fmt.Println("sum:", sum)
}
