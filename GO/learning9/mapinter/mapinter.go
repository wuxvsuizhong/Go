package mapinter

import (
	"fmt"
	"reflect"
)

type human struct {
	Name string
	Age  int
}

type db struct {
	Host     string
	Port     int
	User     string
	Database int
}

func Start() {
	var commap map[string]interface{}
	commap = make(map[string]interface{})

	h1 := human{
		Name: "张三",
		Age:  26,
	}

	htype := reflect.TypeOf(h1)
	// fmt.Println(htype.Name())
	commap[htype.Name()] = &h1

	d1 := db{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Database: 0,
	}

	dtype := reflect.TypeOf(d1)
	// fmt.Println(dtype.Name())
	commap[dtype.Name()] = &d1

	fmt.Printf("%#v\n", commap)

	for item := range commap {
		// fmt.Println(item)
		val := reflect.ValueOf(commap[item])
		for i := 0; i < val.Elem().NumField(); i++ {
			if val.Elem().Field(i).Kind() == reflect.Int {
				val.Elem().Field(i).SetInt(100)
			}
		}
	}

	fmt.Println(h1)
	fmt.Println(d1)
}
