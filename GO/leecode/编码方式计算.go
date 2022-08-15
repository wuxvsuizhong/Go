//一段全为数字的短信使用如下方式编码
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一段编码的短信，计算编码方式

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
func encodeString(s string) int {
	//迭代s,如果连续的两位数字值大于26,那么只能采取单独每一位编码一次的方式,此时的编码方式个数取决于前n-2位的编码方式f(n-2)
	//如果连续的两位的值大于10但是小于或者等于26(11-26之间)，那么可以采取每两位编码一次，也可以是每一位单独编码一次，此时的编码方式就是f(n-1)+f(n-2)
	ecodeMap := map[string]string{}
	for i := 1; i <= 26; i++ {
		ss := fmt.Sprintf("%c", i+96)
		ecodeMap[string(i)] = ss
	}
	// fmt.Println(ecodeMap)
	fmt.Printf("len:%v\n", len(s))

	recMap := map[int]int{0: 1}
	for i := 1; i < len(s); i++ {
		iret, err := strconv.Atoi(s[i-1 : i+1])
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return 0
		}
		if iret > 26 {
			_, ok := ecodeMap[s[i-1]]
			_, ok2 := ecodeMap[s[i]]
			if ok && ok2 { //单独编码都能找到(0不能在编码中找到所以需要排除这种情况)
				recMap[i] = recMap[i-1]
			} else {
				recMap[i] = 0 //如果连续的两位大于26但是其中有无法单独编码的字符
			}

		} else if iret < 10 { //连续的两位的前一位是0
			recMap[i] = recMap[i-1]
		} else {
			if i > 1 {
				recMap[i] = recMap[i-1] + recMap[i-2]
			} else if i == 1 {
				recMap[i] = recMap[i-1] + 1
			}

		}

	}

	fmt.Printf("%#v\n", recMap)
	return recMap[len(s)-1]
}
*/

func encodeString2(s string) int {
	if s == "" || s[0] == '0' {
		return 0
	}

	dp := []int{1, 1}
	for i := 2; i < len(s)+1; i++ {
		result := 0
		n, err := strconv.Atoi(s[i-2 : i])
		if err != nil {
			return 0
		}
		if n >= 10 && n <= 26 { //如果可以是两位编码，那么编码方式数目就是f(n-2)
			result = dp[i-2]
		}
		if s[i-1] != '0' { //如果按照每一位单独编码编码方式数就是f(n-1)
			result += dp[i-1]
		}
		dp = append(dp, result)
	}

	fmt.Println(dp[len(s)])
	return dp[len(s)]

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	line := string(bytes)
	fmt.Printf("%#v\n", line)
	// encodeString(line)
	encodeString2(line)
}
