package main

import (
	"flag"
	"fmt"
	"os"
)

func main1() {
	//获取所有的命令行参数
	fmt.Println("命令行的参数有:", os.Args)

	for i, v := range os.Args { //遍历获取到的参数
		fmt.Printf("args[%v]: %v\n", i, v)
	}

}

/*
模拟获取mysql数据库登录命令行参数
-u root -p 123456 -h localhost -port 3306
*/
func main() {
	var user string
	var password string
	var host string
	var port int
	//第一个参数是获取的放置的变量，第二个参数是选项，比如u，那么对应命令行的-u
	//第三个参数是在未获取到选项参数时，选项的默认值，最后一个是对选项参数的说明信息
	flag.StringVar(&user, "u", "root", "用户名，默认为空")
	flag.StringVar(&password, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机地址，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口，默认为3306")

	flag.Parse() //重要，必须调用一下parse才能把参数解析
	//flag.StringVar以及Intvar的好处是，不用在意命令行的参数顺序，只要是 -option val 这种形式的，选项紧跟其对应的值即可，
	//而选项之间不用在意输入的顺序关系
	fmt.Printf("user=%v,password=%v,host=%v,port=%v\n",
		user, password, host, port)
	/*
	   go run .\main.go -u root -p 3340 -h 127.0.0.5————输出user=root,password=3340,host=127.0.0.5,port=3306
	   go run .\main.go ————输出user=root,password=,host=localhost,port=3306  没有在命令行输入的选项参数会解析默认值
	*/
}
