package main

import (
	"database/sql"
	"fmt"
	"mypro/usesqlx"
	_ "mysql"
)

type user struct {
	id   int
	name string
	age  int
}

var db *sql.DB

func initDB() (err error) {
	dsn := "zhww:231024@tcp(192.168.1.120:3306)/Test"
	db, err = sql.Open("mysql", dsn)
	//Open返回一个数据库的连接池db(类型是*sql.DB)
	if err != nil {
		err = fmt.Errorf("尝试打开数据库失败，err:%v", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("Ping数据库失败，err:%v", err)
		return err
	}
	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数为10个连接
	//当连接数超过数据库设置的最大连接数时，新的连接会被阻塞直到有就的连接释放
	db.SetMaxIdleConns(5) //设置最大限制的连接数
	return nil
}

func queryOne(sqlstr *string) {
	row := db.QueryRow(*sqlstr, 1) //QueryRow 返回一条查询记录
	var u1 user
	row.Scan(&u1.id, &u1.name, &u1.age) //对于返回的row对象必须调用scan方法才能获取值
	//需要修改值需要传入保存结果值得变量指针
	//scan会自动在串完成后释放连接，所以必须调用scan以防止连接被占用不释放
	fmt.Println(u1)

	/*
		for i := 0; i < 20; i++ { //在initDB中最大连接数被设置为10，如果查询不调用scan那么建立的连接会被一直占用，超过10个后查询被阻塞
			db.QueryRow(*sqlstr)
			fmt.Printf("第%d次查询完毕.\n", i)
		}
	*/
}

func queryAll(sqlstr *string) {
	rows, err := db.Query(*sqlstr, 0)
	if err != nil {
		fmt.Printf("queryAll <%s> 查询失败，err:%v\n", *sqlstr, err)
		return
	}

	//多行查询一定要关闭rows
	defer rows.Close()
	for rows.Next() {
		//循环取值
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("queryAll scan失败，err:%v\n", sqlstr, err)
			return
		}
		fmt.Println(u1)
	}
}

func insert(sqlstr *string) {
	result, err := db.Exec(*sqlstr)
	if err != nil {
		fmt.Printf("插入数据<%s>失败，err:%v\n", sqlstr, err)
		return
	}
	//插入数据可拿到相应的id
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("未获取到插入id值，err:%v\n", err)
		return
	}
	fmt.Println("插入成功！id:", id)
}

func updateRow(sqlstr *string, args ...interface{}) {
	result, err := db.Exec(*sqlstr, args...) //可变参数
	if err != nil {
		fmt.Printf("更新数据<%s>失败，err:%v\n", sqlstr, err)
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("未获取到更新数据的id值，err:%v\n", err)
		return
	}
	fmt.Printf("更新成功%d行数据！\n", n)

}

func deleteRow(sqlstr *string, args ...interface{}) {
	result, err := db.Exec(*sqlstr, args...) //可变参数
	if err != nil {
		fmt.Printf("删除数据<%s>失败，err:%v\n", sqlstr, err)
		return
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("未获取到删除数据的id值，err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据！\n", n)
}

//prepare 预处理方式和数据库交互
func prepareInsert(sqlstr *string, valmap *map[string]int) {
	stmt, err := db.Prepare(*sqlstr)
	if err != nil {
		fmt.Printf("prepare 预处理失败，err:%v\n", err)
		return
	}
	defer stmt.Close()
	//预处理后只需用stmt去处理数据
	for k, v := range *valmap {
		_, err := stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("未获取预处理返回结果，err:%v\n", err)
			return
		}
	}
}

/*
prepareQurey()
prepareUpdate()
prepareDelete()
...
*/

//事务处理数据
func transactionProcess(sqlstrs *[]string) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("事务Begin失败，err:%v\n", err)
		return
	}
	for i := 0; i < len(*sqlstrs); i++ {
		_, err := tx.Exec((*sqlstrs)[i])
		if err != nil {
			fmt.Printf("执行第%d条sql语句<%s>失败,err:%v.将回滚事务\n", i, (*sqlstrs)[i], err)
			tx.Rollback()
			//事务的回滚是整个事务整体的sql回滚，即使是事务中前n条sql已经执行成功，只要有一句sql执行失败，那么之前所有的成功sql执行都会回滚
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("事务commit提交失败，err:%v\n", err)
		return
	}
	fmt.Println("事务执行成功!")
	return
}

func main1() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}

	sqlstr := "select id,name,age from user where id=?;"
	queryOne(&sqlstr)
	fmt.Println("====================")
	sqlstr = "select id,name,age from user where id>?"
	queryAll(&sqlstr)
	/*	fmt.Println("====================")
		sqlstr = "insert into user(name,age) values('王五',30)"
		insert(&sqlstr)
	*/
	fmt.Println("====================")
	sqlstr = "update user set age=? where id=?"
	updateRow(&sqlstr, 100, 1)
	fmt.Println("====================")
	sqlstr = "delete from user where id=?"
	deleteRow(&sqlstr, 2)
	fmt.Println("====================")
	sqlstr = "insert into user(name,age) values(?,?)"
	valmap := map[string]int{
		"一一": 21,
		"二二": 24,
		"三三": 25,
	}
	prepareInsert(&sqlstr, &valmap)
	fmt.Println("====================")
	sqls := []string{
		"update user set age=age-5 where id=1",
		"update user set age=age+5 where id=3",
	}
	transactionProcess(&sqls)

}

func main() {
	usesqlx.Start()
}
