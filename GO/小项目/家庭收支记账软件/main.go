package main

import (
	"fmt"
)

func checkinput(a *float64) {
	for {
		if n, err := fmt.Scanln(a); err != nil || n == 0 {
			fmt.Print("\r格式不正确，请输入一个数字：")
		} else {
			break
		}
	}
}

func main() {
	var op int
	var isloop bool = true
	var balance float64 = 10000.00
	var getmoney float64 = 0.0
	var expenditure float64 = 0.0
	var comment string
	var details string
	for {
		fmt.Println("--------家庭收支记账软件--------")
		fmt.Println("----------1.收支明细-----------")
		fmt.Println("----------2.收入登记-----------")
		fmt.Println("----------3.支出登记-----------")
		fmt.Println("----------4.退   出-----------")
		fmt.Println()
		fmt.Print("          请选择:")
		fmt.Scanln(&op)

		switch op {
		case 1:
			fmt.Println("\n--------当前收支明细记录--------")
			fmt.Println("收支\t账户金额\t收支金额\t\t说明")
			fmt.Println(details)
		case 2:
			fmt.Print("\r本次收入金额:")
			//fmt.Scanln(&getmoney)
			checkinput(&getmoney)
			fmt.Print("\r本次收入说明:")
			fmt.Scanln(&comment)
			balance += getmoney
			details += fmt.Sprintf("收入\t%.2f\t%.2f\t\t%s\n", balance, getmoney, comment)
		case 3:
			fmt.Print("登记支出:")
			//fmt.Scanln(&expenditure)
			checkinput(&expenditure)
			fmt.Print("支出说明:")
			fmt.Scanln(&comment)
			balance -= expenditure
			details += fmt.Sprintf("支出\t%.2f\t%.2f\t\t%s\n", balance, expenditure, comment)
		case 4:
			choice := ""
			for {
				fmt.Println("\n确定要退出吗？y/n")
				fmt.Scanln(&choice)
				if choice == "y" || choice == "Y" {
					isloop = false
					break
				} else if choice == "n" || choice == "N" {
					break
				} else {
					fmt.Println("输入有误，请重新输入！")
					continue
				}
			}
			break
		default:
			fmt.Println("\n\r输入不正确!")
			break
		}
		if !isloop {
			break
		}
	}
	fmt.Println("退出......")
}
