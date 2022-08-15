package usestruct

import "fmt"

type person struct{
	string
	int
	//int
	//匿名变量不能同时存在多个
}
//匿名变量在结构体中只能单一的存在


func TestAnonymous(){
	p1 := person{
		"张三",
		23,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
}

type address struct{
	provice string
	city string
}

type worker struct{
	name string
	age int
	addr address
}

type compny struct{
	name string
	addr address
}

type worker2 struct{
	name string
	age int
	address
	//匿名嵌套结构体
}

type workaddr struct{
	provice string
	city string
}

type worker3 struct{
	name string
	age int
	address
	workaddr
	//匿名嵌套的结构体address 和 workaddr 含有相同的属性provice 和 city
}

func TestInnerStruct(){
	w1 := worker{
		name:"李四",
		age:30,
		addr:address{
			provice:"陕西",
			city:"汉中",
		},
	}

	c1 := compny{
		name:"天地银行",
		addr:address{
			provice:"东胜神洲",
			city:"大唐",
		},
	}

	fmt.Println(w1)
	fmt.Println(w1.addr.provice)
	fmt.Println(c1)

	w2 := worker2{
		name:"王五",
		age:31,
		address:address{
			provice:"陕西",
			city:"西安",
		},
	}
	//匿名嵌套结构体，可以直接跨越嵌套层级使用.访问里层属性(前提是被嵌套的结构体中不能含有同名属性)
	fmt.Println(w2.city)

	w3 := worker3{
		name:"老六",
		age:32,
		address:address{
			provice:"上海",
			city:"上海市",
		},
		workaddr:workaddr{
			provice:"陕西",
			city:"西安",
		},
	}

	fmt.Println(w3)
	//如果匿名嵌套的结构体含有相同的属性名，那么既不能直接跨层级访问属性,任然需要层层递进
	fmt.Println(w3.address.city)
	fmt.Println(w3.workaddr.city)
}
