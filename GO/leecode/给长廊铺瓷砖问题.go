package main

import (
	"fmt"
)

//长廊长度为n,宽度为2，每块瓷砖大小是长*宽 1*2
//瓷砖可以横铺，也可以竖着,
//计算铺满长廊有多少种铺法

func putSqure(n int) {
	//如果最后一块瓷砖m是竖着铺，那么铺法就是f(m-1)
	//如果最后一块瓷砖m是横着铺，那么m-1块瓷砖必须要横着铺，所以这种情况下能决定铺法的是第m-2 块瓷砖的铺法f(m-2)
	recordMap := map[int]int{0: 0}
	recordMap[1] = 1
	recordMap[2] = 2 //2*2 的面积，两块瓷砖要么全横着铺，要么全竖着铺
	for i := 3; i < n; i++ {
		recordMap[i] = recordMap[i-1] + recordMap[i-2]
	}

	fmt.Printf("%#v\n", recordMap)

}

func main() {
	putSqure(5)
}
