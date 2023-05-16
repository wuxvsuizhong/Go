package fileop

import (
	"bufio"
	"fmt"
	"io"
	"mypro/getfuncinfo"
	"os"
)

type cntStruct struct {
	letter int
	digit  int
	other  int
}

func CharactersCount() {
	getfuncinfo.PrintFuncName()
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("文件打开失败，err=", err)
		return
	}
	defer file.Close()

	var cnt cntStruct

	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, c := range s {
			//fmt.Println(c)
			/*
				      	switch c 会报错
								此处swotch后不能添加条件c，因为case 是按照值byte码大小来比较的，计算结果是bool类型，而c是byte类型，二者不匹配，编译不通过；
								此处直接用switch相当于转为 if ... else if ...的用法
			*/
			switch {
			case c >= 'a' && c <= 'z':
				//cnt.letter++
				fallthrough
			case c >= 'A' && c <= 'Z':
				cnt.letter++
			case c >= '0' && c <= '9':
				cnt.digit++
			default:
				cnt.other++
			}
		}
	}
	fmt.Printf("字母个数:%v 数字个数:%v 其他字符个数:%v\n", cnt.letter, cnt.digit, cnt.other)
}
