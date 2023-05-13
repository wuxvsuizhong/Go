package main

import "fmt"

/*map是引用类型按照引用传递，调用函数修改会修改源数据*/
func change(m map[string]string, ck string, cv string) {
	m[ck] = cv
}

/*map添加键对时会自动扩容,及时超过容量值也没关系*/
func addone(m map[string]string, ak string, av string) {
	m[ak] = av
}

func main() {
	m := make(map[string]string, 2)
	m["a"] = "10"
	m["b"] = "100"
	fmt.Println("调用修改前:", m)
	change(m, "a", "111")
	fmt.Println("调用修改后:", m)

	addone(m, "c", "22")
	addone(m, "d", "33")
	fmt.Println("超过容量添加key-val后：", m)
}
