package main

import (
	"fmt"
)

func main() {
	var row, col, col2 int
	fmt.Scanln(&row)
	fmt.Scanln(&col)
	fmt.Scanln(&col2)

	m1 := make([][]int, row)
	m2 := make([][]int, col)

	for i := 0; i < row; i++ {
		m1[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Scan(&m1[i][j])
		}
	}

	for i := 0; i < col; i++ {
		m2[i] = make([]int, col2)
		for j := 0; j < col2; j++ {
			fmt.Scan(&m2[i][j])
		}
	}

	fmt.Println(m1)
	fmt.Println(m2)

	res_matrix := [][]int{}
	for i := 0; i < row; i++ {
		ret_list := []int{}
		for j := 0; j < col2; j++ {
			calc_col := []int{}
			for k := 0; k < col; k++ {
				calc_col = append(calc_col, m2[k][j])
			}
			ret_list = append(ret_list, calc_matrix(&m1[i], &calc_col))
		}
		res_matrix = append(res_matrix, ret_list)
	}
	fmt.Println(res_matrix)

}

func calc_matrix(m1 *[]int, m2 *[]int) int {
	sum := 0
	for i, val := range *m1 {
		sum += val * (*m2)[i]
	}

	fmt.Println(sum)
	return sum
}
