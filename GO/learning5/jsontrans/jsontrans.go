package jsontrans

import (
	"encoding/json"
	"fmt"
	"mypro/getfuncinfo"
)

type person struct {
	//json模块为了能够解析struct,struct 里的字段key必须大写字母开头
	//``中的字段的作用是为了满足一些特定场景下被要求返回的必须是小写字段名
	//``中可以写多个，如json:"name" --表示在json格式中字段名为小写name；ini:"name" --在ini配置中字段名为小写name
	Name string `json:"name",ini:"name"`
	Age  int
}

func TestJsonTrans() {
	getfuncinfo.PrintFuncName()
	p1 := person{
		Name: "张三",
		Age:  23,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal fail! err %v\n", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------")
	//反序列化，把json字符串转为struct
	str := `{"Name":"上天和太阳肩并肩","Age":10000}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传递指针为了能够直接赋值变量(go是值传递，直接传递变量难以在外部观察到修改结果)
	fmt.Printf("%#v\n", p2)
}

//反序列化为map
func UnmarshalMap() {
	str := "{\"addr\":\"北京\",\"ahe\":10,\"name\":\"jerry\"}"
	var a map[string]interface{}
	//反序列化时最终转为的map不需要make,unmarshal里已经做了make

	//需要把要反序列化的字符串转化为byte数组，以及传递最终转为map后，map的变量地址
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化map失败,err=", err)
	}
	fmt.Println("反序列化后:", a)
}

//反序列化为切片
func Unmarshal2Slice() {
	str := "[{\"addr\":\"北京\",\"age\":22,\"name\":\"tom\"},{\"addr\":\"北京\",\"ahe\":10,\"name\":\"jerry\"}]\n"

	//切片的元素类型需要和字符串的json格式保持一致，字符串为一个json数组，数组里是一个个的map，那么徐亚定义切片的时候格式要对应
	var a []map[string]interface{}

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化slice失败,err=", err)
	}
	fmt.Println("反序列化后:", a)
}

//切片做json序列化
func SliceMarshal() {
	getfuncinfo.PrintFuncName()
	//一个元素是map的切片
	var slice []map[string]interface{}
	m1 := make(map[string]interface{}) //使用map先make
	m1["name"] = "tom"
	m1["age"] = 22
	m1["addr"] = "北京"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "jerry"
	m2["ahe"] = 10
	m2["addr"] = "北京"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误，err=", err)
	}
	fmt.Printf("slice 序列化后:%v\n", string(data))
	//	slice 序列化后:[{"addr":"北京","age":22,"name":"tom"},{"addr":"北京","ahe":10,"name":"jerry"}]
}
