package main

import (
	"fmt"
	"sort"
)

func main() {
	/*演示案例
	  	个key-value 的value是林外一个map
	    如:存放3个学生信息，每个学生有name ,gender,家庭住址等信息
	  	map形式：map[string]map[string]string
	*/

	stumap := make(map[string]map[string]string)
	stumap["stu1"] = make(map[string]string, 3)
	stumap["stu1"]["name"] = "zhangsan"
	stumap["stu1"]["gender"] = "男"
	stumap["stu1"]["addr"] = "中国"

	stumap["stu2"] = make(map[string]string, 3)
	stumap["stu2"]["name"] = "lisi"
	stumap["stu2"]["gender"] = "女"
	stumap["stu2"]["addr"] = "上海"

	fmt.Println(stumap)

	delete(stumap, "stu3") //删除map中的键值对时，及时key不存在也不会报错，只是实际上什么都没做而已

	val, ok := stumap["stu1"] //在map中查key-val
	if ok {
		fmt.Println("存在key:", val)
	} else {
		fmt.Println("不存在key")
	}

	//map遍历
	for k, _ := range stumap {
		fmt.Println("key1:", k)
		for k2, v2 := range stumap[k] {
			//fmt.Println("key2:", k2)
			fmt.Printf("\t %v:%v \n", k2, v2)
		}
	}
	//len（map) 获取键值对个数
	fmt.Printf("stumap有%d对key-val\n", len(stumap))

	mapsort()
}

func slicemap() {
	/*
	   map的切片
	   形如：[]map[string]string 也就是一个列表里面放置了一个个的map
	*/
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	monsters[0] = make(map[string]string)
	monsters[0]["name"] = "牛魔王"
	monsters[0]["age"] = "500"
	monsters[0]["married"] = "yes"

	monsters[1] = make(map[string]string)
	monsters[1]["name"] = "孙悟空"
	monsters[1]["age"] = "600"

	fmt.Printf("%q\n", monsters)

	monsters = append(monsters, map[string]string{"name": "玉面狐狸", "age": "300"})
	fmt.Printf("扩容后 %q\n", monsters)
}

func mapsort() {
	/*
		map是无序的
		如果需要对mao排序，排序其key值，先把key取出来放到切片中中
		然后对切片排序
		遍历切片，然后按照key的顺序输出map的值
	*/
	m := make(map[string]string, 5)
	m["100"] = "aaaa"
	m["90"] = "1"
	m["80"] = "cccc"
	m["70"] = "23"
	m["60"] = "eeee"

	fmt.Printf("未排序前：%q\n", m)

	mkeys := []string{}
	for k, _ := range m {
		mkeys = append(mkeys, k)
	}
	sort.Strings(mkeys)
	fmt.Println("key值排序后", mkeys)
	for _, val := range mkeys {
		fmt.Printf("%q\n", m[val])
	}

}
