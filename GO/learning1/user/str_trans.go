package user

import (
	"fmt"
	//"strconv"
)

func Strtrans() {
	var s1 string = "aaa\nbbb"
	fmt.Println(s1)
	var s2 string = "aaa\bbbb"
	fmt.Println(s2)
	var s3 string = "aaaaaa\rbbb"
	fmt.Println(s3)
	var s4 string = "aaa\tbbb"
	fmt.Println(s4)
	var s5 string = "aaaaa\tbbb"
	fmt.Println(s5)

	var n1 int = 10
	var n2 float32 = 5.34
	var ss1 string = fmt.Sprintf("%d", n1)
	var ss2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("ss1类型是%T,值是%v\n", ss1, ss1)
	fmt.Printf("ss2的类型是%T,值是%v\n", ss2, ss2)
	fmt.Printf("ss2的类型是%T,值是%q\n", ss2, ss2)

	var n3 bool = false
	var ss3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("ss3的类型是%T,值是%q\n", ss3, ss3)
}

/*func ConvTest() {
        var n1 int = 10
        var s1 string = strconv.FormatInt(int64(n1), 10)
        fmt.Printf("s1 的类型是%T,值是%q\n", s1, s1)

        var n2 float64 = 3.33
        var s2 string = strconv.FormatFloat(n2, 'f', 9, 64)
        fmt.Printf("s2的类型%T,值%q\n", s2, s2)
}
*/
