/*
'''
王强决定把年终奖用于购物，他把想买的物品分为两类：主件与附件，附件是从属于某个主件的，下表就是一些主件与附件的例子：
主件	附件
电脑	打印机，扫描仪
书柜	图书
书桌	台灯，文具
工作椅	无
如果要买归类为附件的物品，必须先买该附件所属的主件，且每件物品只能购买一次。
每个主件可以有 0 个、 1 个或 2 个附件。附件不再有从属于自己的附件。
王强查到了每件物品的价格（都是 10 元的整数倍），而他只有 N 元的预算。除此之外，他给每件物品规定了一个重要度，用整数 1 ~ 5 表示。他希望在花费不超过 N 元的前提下，使自己的满意度达到最大。
满意度是指所购买的每件物品的价格与重要度的乘积的总和，假设设第i件物品的价格为v[i]，重要度为w[i]，共选中了k件物品，编号依次为j1,j2,...,jk ，则满意度为：v[j1]*w[j1]+v[j2]*w[j2]+ … +v[jk]*w[jk]（其中 * 为乘号）
请你帮助王强计算可获得的最大的满意度。


输入描述：
输入的第 1 行，为两个正整数N，m，用一个空格隔开：
（其中 N （ N<32000 ）表示总钱数， m （m <60 ）为可购买的物品的个数。）
从第 2 行到第 m+1 行，第 j 行给出了编号为 j-1 的物品的基本数据，每行有 3 个非负整数 v p q
（其中 v 表示该物品的价格（ v<10000 ）， p 表示该物品的重要度（ 1 ~ 5 ）， q 表示该物品是主件还是附件。如果 q=0 ，表示该物品为主件，如果 q>0 ，表示该物品为附件， q 是所属主件的编号）

输出描述：
 输出一个正整数，为张强可以获得的最大的满意度。
示例1
输入：
1000 5
800 2 0
400 5 1
300 5 1
400 3 0
500 2 0
输出：
2200
示例2
输入：
50 5
20 3 5
20 3 5
10 3 0
10 2 0
10 1 0
输出：
130
说明：
由第1行可知总钱数N为50以及希望购买的物品个数m为5；
第2和第3行的q为5，说明它们都是编号为5的物品的附件；
第4~6行的q都为0，说明它们都是主件，它们的编号依次为3~5；
所以物品的价格与重要度乘积的总和的最大值为10*1+20*3+20*3=130
'''
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var monoey, quantity int
	fmt.Scanln(&monoey, &quantity)
	prices := make([]int, quantity)
	weights := make([]int, quantity)
	attch_rec := make([]int, quantity)
	for i := 0; i < quantity; i++ {
		fmt.Scanln(&prices[i], &weights[i], &attch_rec[i])
	}

	fmt.Println("prices:", prices)
	fmt.Println("weights:", weights)
	fmt.Println("attch_rec:", attch_rec)

	var get_combinations func(grp []int, nums *[]int, res *[][]int, pos int, rec *[]bool)
	get_combinations = func(grp []int, nums *[]int, res *[][]int, pos int, rec *[]bool) {
		if len(grp) != 0 {
			*res = append(*res, grp)
		} else if len(grp) == len(*nums) {
			return
		}

		for i := pos; i < len(*nums); i++ {
			if (*rec)[i] == false {
				(*rec)[i] = true
				new_grp := append(grp, (*nums)[i])
				get_combinations(new_grp, nums, res, pos+1, rec)
				(*rec)[i] = false
				new_grp = new_grp[:len(new_grp)-1]
			}
		}
	}

	attachs := map[int][]int{}
	for i, v := range attch_rec {
		if v > 0 {
			attachs[v] = append(attachs[v], i+1)
		}
	}

	fmt.Println(attachs)
	grp_map := map[int]interface{}{}
	for k, _ := range attachs {
		//fmt.Println(attachs[k])
		rec := make([]bool, len(attachs[k]))
		grp := []int{}
		nums := attachs[k]
		res := [][]int{}
		get_combinations(grp, &nums, &res, 0, &rec)
		grp_map[k] = res
	}
	fmt.Println(grp_map)

	for k, v := range grp_map {
		fmt.Println("v:", v)
		for j := 0; j < len(v.([][]int)); j++ {
			v.([][]int)[j] = append(v.([][]int)[j], k)
		}
		v = append(v.([][]int), []int{k})
		grp_map[k] = v
	}
	fmt.Println(grp_map)

	fmt.Println("grp_map->:", grp_map)

	grp_prices := [][]int{}
	grp_weights := [][]int{}
	grp_values := [][]int{}

	for _, v := range grp_map {
		grp_prices_tmp := []int{}
		grp_weights_tmp := []int{}
		grp_values_tmp := []int{}
		for j := 0; j < len(v.([][]int)); j++ {
			sum := 0
			wsum := 0
			vsum := 0

			for k := 0; k < len(v.([][]int)[j]); k++ {
				index := v.([][]int)[j][k] - 1
				sum += prices[index]
				wsum += weights[index]
				vsum += prices[index] * weights[index]
			}
			grp_prices_tmp = append(grp_prices_tmp, sum)
			grp_weights_tmp = append(grp_weights_tmp, wsum)
			grp_values_tmp = append(grp_values_tmp, vsum)

		}
		grp_prices = append(grp_prices, grp_prices_tmp)
		grp_weights = append(grp_weights, grp_weights_tmp)
		grp_values = append(grp_values, grp_values_tmp)
	}
	grp_prices = append([][]int{{0}}, grp_prices...)
	grp_weights = append([][]int{{0}}, grp_weights...)
	grp_values = append([][]int{{0}}, grp_values...)

	//tmp_list := make([][]int, 1)
	for i := 0; i < len(attch_rec); i++ {
		if attch_rec[i] == 0 {
			if _, ok := grp_map[i+1]; !ok {
				tmp_price_list := [][]int{{prices[i]}}
				tmp_weight_list := [][]int{{weights[i]}}
				tmp_value_list := [][]int{{prices[i] * weights[i]}}
				//tmp_list[0] = append(tmp_list[0], i+1)
				grp_prices = append(grp_prices, tmp_price_list...)
				grp_weights = append(grp_weights, tmp_weight_list...)
				grp_values = append(grp_values, tmp_value_list...)

			}
		}
	}

	fmt.Println(grp_prices)
	fmt.Println(grp_weights)
	fmt.Println(grp_values)

	dp := make([][]int, len(grp_prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, monoey+1)
	}

	for i := 1; i < len(grp_prices); i++ {
		for j := 1; j <= monoey; j++ {
			for k := 0; k < len(grp_prices[i]); k++ {
				if grp_prices[i][k] > j {
					tmp_nums := []int{dp[i-1][j], dp[i][j]}
					sort.Ints(tmp_nums)
					dp[i][j] = tmp_nums[len(tmp_nums)-1]
				} else {
					tmp_nums := []int{dp[i-1][j], dp[i-1][j-grp_prices[i][k]] + grp_values[i][k], dp[i][j]}
					sort.Ints(tmp_nums)
					dp[i][j] = tmp_nums[len(tmp_nums)-1]
				}
			}
		}
	}

	fmt.Println(dp[len(grp_prices)-1][monoey])
}
