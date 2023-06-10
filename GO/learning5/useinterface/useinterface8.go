package useinterface

import (
	"fmt"
	"math/rand"
	"mypro/getfuncinfo"
	"sort"
)

/*
接口编程的实践——实现对某个struct的排序
通过调用sort包的Interface接口实现，sort包的Interface接口有三个方法Len,Less和Swap
通过把struct放在list切片中，把切片作为一个类型，使用该类型实现这三个方法就可以对该list排序，形从而实现对struct的排序
步骤：
1.定义struct结构体
2.定义切片用于放置一个个的strut
3.切片定义为一个新的类型，然后给该类型绑定实现Len,Less,Swap方法
4.调用sort.Sort()
*/

type Hero struct {
	Name string
	Age  int
}

type HeroList []Hero

func (hl HeroList) Len() int {
	return len(hl)
}

func (hl HeroList) Less(i, j int) bool {
	return hl[i].Age < hl[j].Age
}

func (hl HeroList) Swap(i, j int) {
	hl[i], hl[j] = hl[j], hl[i]
}

//Len,Less,Swap三个方法都实现以后，就意味着实现了Inerface接口

func HeroSort() {
	getfuncinfo.PrintFuncName()
	var heros HeroList
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	//  排序前
	fmt.Printf("排序前:%v\n", heros)
	sort.Sort(heros)
	//排序后
	fmt.Printf("排序后:%v\n", heros)
}

/*
总结：方法的绑定可以是任何的类型type，而go中可以把任意的类型使用type关键字自定义为自己的类型
通过定义后的类型t，可以给类型t绑定方法，而接口实际是把一堆的方法做了集合，当某个类型实现了接口I包含的所有方法，那么该类型t就实现了接口I
然后就可以通过接口I 指代这个类型t，这就相当于多态，只要定义的类型t实现了接口I中包含的所有方法，那么接口I 可以用于指代类型t的任意实例
有了这个多态，我么就可以通过一些通用的方法，针对某个多态的类型，实现通用的操作如排序，格式化输出等等
*/
