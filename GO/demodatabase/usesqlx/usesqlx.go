package usesqlx

import (
	"fmt"
	"sqlx"
)

func initDB(db **sqlx.DB) (err error) {
	dsn := "zhww:231024@tcp(192.168.1.120:3306)/Test"
	*db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sqlx 连接数据库失败,err:", err)
		return err
	}
	(*db).SetMaxOpenConns(10)
	(*db).SetMaxIdleConns(5)
	return nil
}

//属性名称大写，是为了sqlx在Get查询的时候，通过反射访问到结构体中的属性名
type user struct {
	Id   int
	Name string
	Age  int
}

func Start() {
	var db *sqlx.DB
	err := initDB(&db)
	if err != nil {
		return
	}
	var u user
	sqlstr := "select id,name,age from user where id=1"
	db.Get(&u, sqlstr)
	fmt.Println(u)
	var ulist []user
	sqlstr = "select id,name,age from user;"
	err = db.Select(&ulist, sqlstr)
	if err != nil {
		fmt.Println("sqlx select查询失败,err:", err)
		return
	}
	fmt.Printf("%#v\n", ulist)
}
