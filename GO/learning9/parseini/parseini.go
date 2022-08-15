package parseini

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type DBConfig struct {
	mysql `ini:"mysql"`
	redis `ini:"redis"`
}

type mysql struct {
	Host     string `ini:"address"`
	Port     int16  `ini:"port"`
	Password string `ini:"password"`
	User     string `ini:"username"`
	Database string `ini:"database"`
}

type redis struct {
	Host     string `ini:"address"`
	Port     int16  `ini:"port"`
	User     string `ini:"username"`
	Database int16  `ini:"database"`
}

func loadIni(fname string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("数据不是指针类型")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("指针指向不是结构体类型")
		return err
	}

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	lines := strings.Split(string(b), "\n")
	v := reflect.ValueOf(data)
	var section string

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if line[0] == '[' && line[len(line)-1] == ']' {
			if secContain := line[1 : len(line)-1]; len(secContain) > 0 {
				//识别到ini节点,查找对应的结构体保存在section
				for i := 0; i < t.Elem().NumField(); i++ {
					if reflect.StructTag(t.Elem().Field(i).Tag.Get("ini")) == reflect.StructTag(secContain) {
						section = t.Elem().Field(i).Name
						// section = reflect.ValueOf(t.Elem().Field(i).Name)
						fmt.Printf("找到ini节点%s对应的结构体%v\n", secContain, section)
						break
					}
				}
			}
		} else if eqPos := strings.Index(line, "="); eqPos != -1 {
			if strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("第%d行,格式不正确:%s", i+1, line)
				return err
			}
			key := strings.TrimSpace(line[:eqPos])
			val := strings.TrimSpace(line[eqPos+1:])
			structObj := v.Elem().FieldByName(section)
			if structObj.Kind() != reflect.Struct {
				err := fmt.Errorf("%v不是一个结构体", structObj)
				return err
			}

			tStructObj := structObj.Type()
			for i := 0; i < tStructObj.NumField(); i++ {
				// fmt.Println(tStructObj.Field(i).Tag.Get("ini"))
				if tStructObj.Field(i).Tag.Get("ini") == key {
					fmt.Printf("%s->%s:%s\n", section, tStructObj.Field(i).Name, val)
					switch tStructObj.Field(i).Type.Kind() {
					case reflect.Int:
						intval, _ := strconv.ParseInt(val, 10, 16)
						structObj.Field(i).SetInt(intval)
					case reflect.String:
						structObj.Field(i).SetString(val)
					}
				}
			}
		}
	}
	return nil
}

func Start() {
	dbcfg := DBConfig{}
	err := loadIni("./db.ini", &dbcfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", dbcfg)
}
