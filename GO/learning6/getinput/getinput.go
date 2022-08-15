package getinput
import(
	"fmt"
	"bufio"
	"os"
)

func GetInput(){
	var s string
	fmt.Print("请输入：")
	fmt.Scanln(&s)
	fmt.Printf("输入的是:%s\n",s)
	//Scanln 无法处理输入的时候有空格的情况，它遇到空格就停止读取了
}

func GetInputByBufio(){
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入：")
	//读取数据，遇到\n就结束读取
	s,_ = reader.ReadString('\n')
	fmt.Printf("输入的是:%s\n",s)
	//bufio 可以连续读取知道遇到设置的字符才停止
}


