package main

import "fmt"

func test(b byte) byte {
	return b + 1
}

func main() {
	var key byte
	fmt.Println("请输入一个字符 [a,b,c,d,e,f,g]")
	_, err := fmt.Scanf("%c", &key)
	if err != nil {
		fmt.Println(err)
	}

	switch test(key) + 1 { //switch计算的是表达式的最终的值
	//case结束不用加break
	case 'a':
		fmt.Println("捕获到a") //输入单个字符的时候不会被执行到，及时输入最小的a,表达式计算结果都会是'a' + 2 = 'c'
	case 'b':
		fmt.Println("捕获到b")
	case 'c':
		fmt.Println("捕获到c")
	case 'd':
		fmt.Println("捕获到d")
	default:
		fmt.Println("其他...")
	}

	/*
		switch和case的检查条件的类型需要一致
	*/
	var n1 int32 = 20
	var n2 int32 = 20
	//var n2 int64 = 20
	switch n1 {
	//case n2:   // case后的检查值的类型需要和switch的一致
	case n2, 10, 5: //case后可以添加多个条件
		fmt.Println("ok1")
	default:
		fmt.Println("其他...")
	} //输出ok1

	/*
		switch后面也可以不加表达式，相当于 if ... else ...的用法
	*/
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("其他...")
	}

	/*
		case 条件是一个范围
		按照case的顺序从上到下匹配，先到达那个条件就执行哪个条件后的语句，执行完后直接返回
	*/
	var score int
	fmt.Printf("输入分数：")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println(err)
	}
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score > 80:
		fmt.Println("成绩良好")
	case score > 70:
		fmt.Println("再接再厉")
	case score > 60:
		fmt.Println("多多努力")
	}

	/*
		switch 穿透————fallthrough
	*/
	//var num int = 10
	var num int = 20
	switch num {
	case 10:
		fmt.Println("ok1")
	case 20:
		fmt.Println("ok2")
		fallthrough //默认只能穿透1层
	case 30:
		fmt.Println("ok3") //上一个case有fallthrough，当上一个条件满足的时候，这个case也会执行
		fallthrough        //如果要穿透多层那么加多个fallthrough
	case 40:
		fmt.Println("ok4")
	default:
		fmt.Println("其他")
	}

	/*
		switch 判断类型
	*/
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型 %T\n", i)
	case int:
		fmt.Printf("x 是 int类型")
	case float64:
		fmt.Println("x 是float64类型")
	default:
		fmt.Println("其他")
	}
}
