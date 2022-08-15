//给定一组数代表股票每天的价格，计算在这几天买卖各一次能够得到的最大利润
//分析：实时记录和更新遍历过的价格中最小的值，每一天的价格和其相减的值做比较得到的利润，与之前的利润相比取最大值
//这类型问题就是获取一个数组中两个数字差值最大的过程

package main

import (
	"fmt"
)

func getProfect(arr []int) int {
	profect, minPrice := 0, arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < minPrice { //实时记录刷新最小价位
			minPrice = arr[i]
		}

		if arr[i]-minPrice > profect { //如果天的价位与历史最小价位的差值大于之前的profect
			profect = arr[i] - minPrice
		}

	}

	return profect
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(getProfect(nums))
}
