// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var snumSlice = []string{}
	for scanner.Scan() {
		if scanner.Text() == "\r" {
			break
		}
		snumSlice = append(snumSlice, (strings.Split(strings.TrimSpace(scanner.Text()), " "))...)

	}

	fmt.Printf("%#v\n", snumSlice)
}

func toInt(s []string) (inums []int) {
	results := []int{}
	for _, val := range s {
		n, _ := strconv.Atoi(val)
		results = append(results, n)
	}
	return results
}

func getMin(nums *[]int) int {
	if len(*nums) == 0 {
		return 0
	}
	var min int = (*nums)[0]
	for _, v := range *nums {
		if v < min {
			min = v
		}
	}
	return min
}

func coinsProcess(total int, coins *map[int]interface{}) int {
	// countArr := make([]int, total+1)
	countArr := make(map[int]int)
	coinList := []int{}
	for v, _ := range *coins {
		countArr[v] = 1
		coinList = append(coinList, v)
	}

	// hascoin := func(k int) bool {
	// 	if _, ok := (*coins)[k]; ok {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }

	for i := 1; i <= total; i++ {
		if i < getMin(&coinList) {
			countArr[i] = 0
			continue
		}

		arr := []int{}
		for v, _ := range *coins {
			ret := i - v
			if (ret == 0) || (ret > 0 && countArr[ret] != 0) { //可细分
				arr = append(arr, countArr[ret]+1)
			}
		}
		fmt.Println("i", i, "arr", arr, "getMin:", getMin(&arr))
		// fmt.Println(arr)
		countArr[i] = getMin(&arr)
	}
	fmt.Println(countArr)
	fmt.Println(countArr[total])
	if countArr[total] == 0 {
		return -1
	} else {
		return countArr[total]
	}

}

func main() {
	ireader := bufio.NewReader(os.Stdin)
	bytes, _, _ := ireader.ReadLine()
	// fmt.Printf("%#v\n", string(bytes))
	nums := toInt(strings.Split(strings.TrimSpace(string(bytes)), " "))
	fmt.Printf("%#v\n", nums)
	coins := make(map[int]interface{}, len(nums))
	var mem interface{}
	for _, v := range nums {
		coins[v] = mem
	}
	bytes, _, _ = ireader.ReadLine()
	total := toInt(strings.Split(strings.TrimSpace(string(bytes)), " "))[0]
	fmt.Printf("%#v\n", total)

	fmt.Println(coinsProcess(total, &coins))
}
