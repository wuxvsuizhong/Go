package main

import (
	"encoding/json"
	"fmt"
)

type AutoGenerated struct {
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Child []int  `json:"child"`
}

func main() {
	jsonStr1 := `{"age":14,"name":"potter","child":[1,2,3]}`
	a := AutoGenerated{}
	json.Unmarshal([]byte(jsonStr1), &a)
	aa := a.Child
	fmt.Println(aa) //打印[1,2,3]
	jsonStr2 := `{"age":12,"name":"potter","child":[3,4,5,6,7,8]}`
	json.Unmarshal([]byte(jsonStr2), &a)
	fmt.Println(aa) // 打印[3,4,5]

	//	对于json从字符串解析到结构体中，当第一次解析生成的结构体实例长度确定后，就不在变化了
	// 如果第二次再重复利用之前的结构体，一定要注意各个字段的长度变化，因为超过初始生成的结构体的字段长度，会被截断
	// 要么每次解析的时候读重新实例化一份结构体避免截断问题
}
