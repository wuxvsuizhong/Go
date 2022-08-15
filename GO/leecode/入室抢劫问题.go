package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 不能连续抢劫连续的两个房子
// 每个房子里能得到的物品价值随机
// 求能最多抢到多少价值的物品

func map2Int(strs []string) (items []int) {
	arr := []int{}
	for _, v := range strs {
		ret, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		arr = append(arr, ret)
	}

	return arr
}

func getCalc(arr []int) int {
	recordMap := map[string]int{
		"00": 0,
	}
	for i, v := range arr {
		for j := 0; j < 2; j++ {
			if j == 0 { //不抢劫当前房子i
				// recordMap[i] = recordMap[i]
				if recordMap[string(i-1)+"0"] > recordMap[string(i-1)+"1"] {
					recordMap[string(i)+"0"] = recordMap[string(i-1)+"0"]
				} else {
					recordMap[string(i)+"0"] = recordMap[string(i-1)+"1"]
				}
			} else { //抢劫当前房子
				if i-1 >= 0 {
					recordMap[string(i)+"1"] = recordMap[string(i-1)+"0"] + v
				} else {
					recordMap[string(i)+"1"] = v
				}

			}
		}
	}

	max := recordMap[string(len(arr)-1)+"1"]
	if recordMap[string(len(arr)-1)+"0"] > max {
		max = recordMap[string(len(arr)-1)+"0"]
	}
	fmt.Printf("%#v\n", recordMap)
	fmt.Printf("%#v\n", max)
	return max
}

func getCalc2(arr []int) int {
	//不能连续抢劫相邻的房间
	//在所有的房间之间是环形连接的情况下，如果抢劫了首个房间那么就不能抢劫和它衔接的尾部
	robrange := func(nums []int, start int, end int) (ret int) {
		if len(nums) == 0 {
			return 0
		} else if len(nums) == 1 {
			return nums[0]
		}

		// do...while
		yes, no := nums[start], 0
		for i := start + 1; i < end; i++ {
			t_no := no
			if yes > no {
				t_no = yes
			}

			t_yes := nums[i] + no
			yes, no = t_yes, t_no
			fmt.Println(yes, no)
		}

		if yes > no {
			return yes
		} else {
			return no
		}
	}

	r1 := robrange(arr, 0, len(arr)-1)
	r2 := robrange(arr, 1, len(arr))
	if r1 > r2 {
		fmt.Printf("%#v\n", r1)
		return r1
	} else {
		fmt.Printf("%#v\n", r2)
		return r2
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	hvals := []int{}
	for {
		bytes, _, _ := reader.ReadLine()
		if strings.TrimSpace(string(bytes)) == "" {
			break
		}
		ns := map2Int(strings.Split(strings.TrimSpace(string(bytes)), " "))
		hvals = append(hvals, ns...)
		// getCalc(hvals)
		getCalc2(hvals)
	}

	fmt.Printf("%#v\n", hvals)
}
