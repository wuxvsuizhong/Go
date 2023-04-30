package main

import (
	"fmt"
)

func main() {
	var s string = "hello,中国"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i]) //字符串中有中文会出现乱码，因为s[i] 是按照字节取得，而汉字是3个字节，如果按照单个字节取会不完整
	}

	var strs = []rune(s) //转换为rune数组，访问中文就没问题了
	for i := 0; i < len(strs); i++ {
		fmt.Printf("%c\n", strs[i]) //每次取rune数组的元素
	}

	for _, v := range s { //range 即使有中文也没问题
		fmt.Printf("%c\n", v)
	}
}
