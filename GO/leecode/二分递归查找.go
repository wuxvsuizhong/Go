package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func BinarySearch(arr *[]int, lindex int, rindex int, findVal int) {
	if rindex < lindex {
		fmt.Println("没找到", findVal)
		return
	}
	mid := (lindex + rindex) / 2
	if (*arr)[mid] == findVal {
		fmt.Printf("找到了，目标位于下标：%v\n", mid)
		return
	} else if findVal > (*arr)[mid] {
		lindex = mid + 1
		BinarySearch(arr, lindex, rindex, findVal)
	} else if findVal < (*arr)[mid] {
		rindex = mid - 1
		BinarySearch(arr, lindex, rindex, findVal)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 10; i++ {
		//fmt.Println(rand.Intn(1000))
		arr = append(arr, rand.Intn(10))
	}
	fmt.Println("初始数组", arr)
	sort.Ints(arr)
	fmt.Println("排序后", arr)
	//findIndex := rand.Intn(len(arr)-1)
	BinarySearch(&arr, 0, len(arr)-1, 10)
}
