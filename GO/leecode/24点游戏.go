/*
给出4个1-10的数字，通过加减乘除运算，得到数字为24就算胜利,除法指实数除法运算,运算符仅允许出现在两个数字之间,本题对数字选取顺序无要求，但每个数字仅允许使用一次，且需考虑括号运算
此题允许数字重复，如3 3 4 4为合法输入，此输入一共有两个3，但是每个数字只允许使用一次，则运算过程中两个3都被选取并进行对应的计算操作。
输入描述：
读入4个[1,10]的整数，数字允许重复，测试用例保证无异常数字。

输出描述：
对于每组案例，输出一行表示能否得到24点，能输出true，不能输出false

示例1
输入：
7 2 1 10
输出：
true
*/
package main

import "fmt"

var ops map[string]bool

func main() {
	//nums := [4]interface{}{}
	// 创建空接口类型的切片，可以存放不同类型的数据
	nums := make([]interface{}, 0)
	n := 0
	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		nums = append(nums, n)
	}
	//fmt.Println(nums)

	ops = map[string]bool{"+": true, "-": true, "*": true, "/": true}
	//fmt.Println(ops)

	options := []interface{}{"+", "-", "*", "/"}
	options = append(options, nums...)

	//fmt.Println(options)

	book := make([]bool, len(options))
	var deep_search func([]interface{}, int)
	res := []interface{}{}

	deep_search = func(calc_list []interface{}, f_quqntity int) {
		if f_quqntity > 3 {
			return
		}
		if len(calc_list) == 7 {
			res = append(res, calc_list)
		}

		for index, val := range options {
			if index >= 0 && index <= 3 {
				//	取符号
				if len(calc_list) <= 2 {
					continue
				}
				new_calc_list := append(calc_list, val)
				deep_search(new_calc_list, f_quqntity+1)
				new_calc_list = new_calc_list[:len(new_calc_list)-1]
			} else { // 取数字
				if !book[index] {
					book[index] = true
					new_calc_list := append(calc_list, val)
					deep_search(new_calc_list, f_quqntity)
					book[index] = false
					new_calc_list = new_calc_list[:len(new_calc_list)-1]
				}
			}
		}
	}

	deep_search([]interface{}{}, 0)
	//fmt.Println(res)

	is_arrive := false
	for _, v := range res {
		is_arrive = false
		if calc_slist(v.([]interface{})) == 24 {
			fmt.Println("true")
			is_arrive = true
			break
		}
	}
	if !is_arrive {
		fmt.Println("false")
	}
}

func calc_slist(slist []interface{}) int {
	stack := []int{}
	for _, v := range slist {
		if iv, ok := v.(int); ok {
			stack = append(stack, iv)
		} else {
			if len(stack) < 2 {
				return 0
			}
			n2 := int(stack[len(stack)-1])
			n1 := int(stack[len(stack)-2])
			stack = stack[:len(stack)-1]

			switch v {
			case "+":
				stack[len(stack)-1] = n1 + n2
				break
			case "-":
				stack[len(stack)-1] = n1 - n2
				break
			case "*":
				stack[len(stack)-1] = n1 * n2
				break
			case "/":
				if n2 == 0 {
					return 0
				}
				stack[len(stack)-1] = n2 / n1
				break
			}
		}
	}
	if len(stack) == 1 {
		return stack[0]
	} else {
		return 0
	}
}
