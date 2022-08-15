//在一组数中,找到其中连续子数组中最大的乘积
//因为数组元素有正数和负数，所以每次乘积不仅要记录最大的值，还要记录最小的值，防止虽小是负数但是遇到拎一个负数转为最大正数的情况

package main

import (
	"fmt"
)

func getMax(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func getMin(nums []int) int {
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, min, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max = getMax([]int{max * nums[i], min * nums[i], nums[i]})
		min = getMin([]int{max * nums[i], min * nums[i], nums[i]})
		//每次记录max和min
		result = getMax([]int{result, max}) //和上一次的比较选择较大的
	}

	return result
}

func main() {
	ret := maxProduct([]int{2, 3, -2, 4})
	fmt.Println(ret)
}
