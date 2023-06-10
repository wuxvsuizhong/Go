package getinput

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var s string
	fmt.Print("请输入：")
	fmt.Scanln(&s)
	fmt.Printf("输入的是:%s\n", s)
	//Scanln 无法处理输入的时候有空格的情况，它遇到空格就停止读取了
	fmt.Print("2请输入：")
	fmt.Scanln(&s)
	fmt.Printf("输入的是:%s\n", s)

	var name string
	var age int8
	var sex int8
	fmt.Println("请输入姓名, 年龄, 姓别")
	fmt.Scan(&name, &age, &sex)
	fmt.Printf("姓名: %v 年龄: %v 性别: %v", name, age, sex)

	var en int
	fmt.Print("请输入一个数字：")
	fmt.Scan(&en)
	fmt.Println("输入的数字是:", en)
	fmt.Scan(&en)

}

func GetInputByBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入：")
	//读取数据，遇到\n就结束读取
	s, _ = reader.ReadString('\n')
	fmt.Printf("输入的是:%s\n", s)
	//bufio 可以连续读取知道遇到设置的字符才停止
}
