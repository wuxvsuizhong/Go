package myfuncs

//一个被测试的函数，计算1-n的累加值
func AddUp(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func Sub(n1, n2 int) int {
	return n1 - n2
}
