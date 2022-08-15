//一个很长的台阶，每次落下一步都需要花费若干不等的费用
//一次只能走一步或者两步，计算最少花费多少可以走完台阶

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func change2Int(arr []string) (items []int) {
	retItems := []int{}
	for _, v := range arr {
		ret, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		retItems = append(retItems, ret)
	}

	return retItems
}

func walk(items []int) int {
	length := len(items)
	if length == 0 {
		return 0
	} else if length == 1 {
		return items[1]
	}

	recordMap := map[int]int{0: items[0], 1: items[1]}
	for i := 2; i < length; i++ {
		if recordMap[i-1]+items[i] < recordMap[i-2]+items[i] {
			recordMap[i] = recordMap[i-1] + items[i]
		} else {
			recordMap[i] = recordMap[i-2] + items[i]
		}
	}

	fmt.Println(recordMap)
	return recordMap[length-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := []int{}
	for {
		bytes, _, _ := reader.ReadLine()
		if strings.TrimSpace(string(bytes)) == "" {
			break
		}
		t_arr := change2Int(strings.Split(strings.TrimSpace(string(bytes)), " "))
		nums = append(nums, t_arr...)
		walk(nums)
	}

}
