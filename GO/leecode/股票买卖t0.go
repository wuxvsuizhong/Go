//给定一组数字中代表股票的历史价位，股票不限制买卖次数,求能得到的最大利润
//分析：每次一旦后一个价位高于前一个价位那么就可以计算利润，就是一个相邻数组元素之间差值累加的过程
package main

import (
	"fmt"
)

func getProfit(nums []int) int {
	max_profit := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			max_profit += nums[i] - nums[i-1]
		}
	}
	return max_profit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4, 13}
	fmt.Println(getProfit(nums))
}
