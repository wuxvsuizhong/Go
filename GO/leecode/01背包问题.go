/*
 一个能抓装M公斤的背包，现有n件物品，它们的重量分别是：W1,W2,W3...Wn，它们的价值分别是C1,C2,C3...Cn,求该背包能够装下最大多大价值的物品。
输入：
10 4
2 1
3 3
4 5
7 9

第一行10 4标识背包的容量为10，然后一共有4件物品
接下来的4行分别是每件物品的重量以及价值
*/
package main

import "fmt"

func main() {
	var M, quantity int
	fmt.Scanln(&M, &quantity)

	W := make([]int, quantity+1) //各个物品的重量
	C := make([]int, quantity+1) //各个物品的价值
	for i := 1; i <= quantity; i++ {
		fmt.Scanln(&W[i], &C[i])
	}

	dp := make([][]int, quantity+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, M+1)
	}

	for i := 1; i <= quantity; i++ {
		for j := 1; j <= M; j++ {
			if W[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				if dp[i-1][j] > dp[i-1][j-W[i]]+C[i] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j-W[i]] + C[i]
				}
			}
		}
	}

	fmt.Printf("%#v\n", dp)
	fmt.Println(dp[quantity][M])
}
