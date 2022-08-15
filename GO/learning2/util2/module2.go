package util2
import "fmt"


var U2_num1 int = 200
var U2_str1 string = "hello , this is util2"

func init(){
	fmt.Println("util2 中的init 被执行...")
	U2_num1 += 30
}
