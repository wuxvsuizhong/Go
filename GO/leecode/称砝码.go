/*
现有n种砝码，重量互不相等，分别为 m1,m2,m3…mn ；
每种砝码对应的数量为 x1,x2,x3...xn 。现在要用这些砝码去称物体的重量(放在同一侧)，问能称出多少种不同的重量。

注：

称重重量包括 0

数据范围：每组输入数据满足 1≤n≤10,1≤mi≤2000,1≤xi≤10
输入描述：
对于每组测试数据：
第一行：n --- 砝码的种数(范围[1,10])
第二行：m1 m2 m3 ... mn --- 每种砝码的重量(范围[1,2000])
第三行：x1 x2 x3 .... xn --- 每种砝码对应的数量(范围[1,10])
输出描述：
利用给定的砝码可以称出的不同的重量数

示例1
输入：
2
1 2
2 1
输出：
5
说明：
可以表示出0，1，2，3，4五种重量。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var typeno int
	fmt.Scanln(&typeno)
	famas := make([]int, typeno)
	weights := make([]int, typeno)
	for i := 0; i < typeno; i++ {
		fmt.Scan(&famas[i])
	}

	for i := 0; i < typeno; i++ {
		fmt.Scan(&weights[i])
	}
	fmt.Println(famas, weights)

	w := []int{}
	for i := 0; i < len(famas); i++ {
		for j := 0; j < weights[i]; j++ {
			w = append(w, famas[i])
		}
	}
	//fmt.Println(w)
	sort.Ints(w)
	fmt.Println(w)

	rec := make([]bool, len(w))
	res := []int{}
	//res := make(map[int]interface{})

	var dfs func(int, int, int)
	dfs = func(sumval int, pos int, steps int) {
		res = append(res, sumval)
		//res[sumval] = true
		//if steps == len(w) {
		if steps == 2 {
			return
		}

		for i := pos; i < len(w); i++ {
			if rec[i] == false {
				if i > 0 && w[i] == w[i-1] && rec[i-1] == false {
					continue
				}
				rec[i] = true
				sumval += w[i]
				dfs(sumval, i+1, steps+1)
				rec[i] = false
				sumval -= w[i]
			}
		}
	}

	dfs(0, 0, 0)
	fmt.Println(res)
	//fmt.Println(len(res))
}
