package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	esCli, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth("elastic", "231024"),
	)
	if err != nil {
		fmt.Printf("创建ES连接客户端失败,err:%v\n", err)
		return
	}
	fmt.Println("连接ES成功!")

	p1 := Person{Name: "shangsan", Age: 100}
	put1, err := esCli.Index().Index("stu").BodyJson(p1).Do(context.Background())
	if err != nil {
		fmt.Printf("往ES put数据出错,err:\n", err)
		return
	}
	fmt.Printf("Index User %s to index %s,type %s\n", put1.Id, put1.Index, put1.Type)
}
