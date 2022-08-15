package progressctl

import "fmt"

func TestFor(){
	var s1 string = "golang 你好"
	for i := 0;i < len(s1);i++ {
		fmt.Printf("%c\n",s1[i])
	}
	//一个字节一个字节的遍历，针对中文由于是UTF8编码会有输出异常
}

func TestFor2(){
	var s2 string = "golang 你好"
	for i,val := range s2 {
			fmt.Printf("下标%d的值:%c\n",i,val)
	}
	//range遍历按照元素遍历，不会分割字符，序号用i接受，下标对应的值用val接收,可以正常输出中文
}

func TestGoto(){
	fmt.Println("line1")
	fmt.Println("line2")
	fmt.Println("line3")
	fmt.Println("line4")
	if 1 == 1{
		goto label //立即跳转到label标号处开始向下执行
	}
	fmt.Println("line5")
	fmt.Println("line6")
	fmt.Println("line7")
	fmt.Println("line8")
	fmt.Println("line9")
	label:
	fmt.Println("line10")
	fmt.Println("line11")
}
