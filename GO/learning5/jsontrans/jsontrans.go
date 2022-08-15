package jsontrans

import(
	"fmt"
	"mypro/getfuncinfo"
	"encoding/json"
)

type person struct{
	//json模块为了能够解析struct,struct 里的字段key必须大写字母开头
	//``中的字段的作用是为了满足一些特定场景下被要求返回的必须是小写字段名
	//``中可以写多个，如json:"name" --表示在json格式中字段名为小写name；ini:"name" --在ini配置中字段名为小写name
	Name string `json:"name",ini:"name"`
	Age int
}

func TestJsonTrans(){
	getfuncinfo.PrintFuncName()
		p1 := person{
			Name:"张三",
			Age:23,
		}
	b,err := json.Marshal(p1)
	if err != nil{
		fmt.Printf("json.Marshal fail! err %v\n",err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------")
	//反序列化，把json字符串转为struct
	str := `{"Name":"上天和太阳肩并肩","Age":10000}`
	var p2 person
	json.Unmarshal([]byte(str),&p2) //传递指针为了能够直接赋值变量(go是值传递，直接传递变量难以在外部观察到修改结果)
	fmt.Printf("%#v\n",p2)
}

