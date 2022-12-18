package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}

	fmt.Println(dp)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				dp[i][j] = true
			}
		}
	}

	res := []string{}

	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if i-j >= 2 {
					dp[i][j] = dp[i-1][j+1]
				} else {
					dp[i][j] = true
				}
			} else {
				dp[i][j] = false
			}

			if dp[i][j] && s[j:i+1] != "" {
				res = append(res, s[j:i+1])
			}
		}
	}

	fmt.Println(res)

}
