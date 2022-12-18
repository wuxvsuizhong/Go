package main

import "fmt"

func is_primer(n int) bool {
	//	检测一个数是否素数
	//fmt.Printf("检测%d是否素数 ", n)
	if n <= 2 && n > 0 {
		return true
	}
	for m := 2; m <= n/2; m++ {
		if n%m == 0 {
			//fmt.Println("no")
			return false
		}
	}
	//fmt.Println("yes")
	return true
}

func main() {
	quantity := 0
	fmt.Scanln(&quantity)
	nums := make([]int, quantity)
	for i := 0; i < len(nums); i++ {
		fmt.Scanf("%d", &nums[i])
	}
	//fmt.Println(nums)
	odds := []int{}
	evens := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evens = append(evens, nums[i])
		} else {
			odds = append(odds, nums[i])
		}
	}

	// 创建二维矩阵，纵向对应偶数，横向对应奇数
	arr := make([][]bool, len(evens))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]bool, len(odds))
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = is_primer(evens[i] + odds[j])
		}
	}

	//fmt.Println(arr)

	var match func(oi int) bool

	evens_visited := make([]bool, len(evens))

	evens_match := make([]int, len(evens))
	for i := 0; i < len(evens_match); i++ {
		evens_match[i] = -1
	}

	match = func(oi int) bool {
		//迭代每个偶数检查是否能够和传入的奇数进行匹配
		for j := 0; j < len(evens); j++ {
			if arr[j][oi] == true && evens_visited[j] == false {
				evens_visited[j] = true
				if evens_match[j] == -1 || match(evens_match[j]) == true {
					evens_match[j] = oi
					return true
				}
			}
		}
		return false
	}

	cnt := 0
	for i := 0; i < len(odds); i++ {
		if match(i) == true {
			cnt += 1
		}

		for j := 0; j < len(evens_visited); j++ {
			evens_visited[j] = false
		}
	}

	fmt.Println(cnt)

}
