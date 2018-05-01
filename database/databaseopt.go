package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var dataSourceName = "root:123456@tcp(localhost:3306)/mytest?charset=utf8"
var database sql.DB

type JDBCObject struct {
	object interface{}
	name string
}

func (JDBCObject) JDBC_connect() {
	println("connect.....")
	db,err := sql.Open("mysql", dataSourceName)
	if err != nil {
		println(err)
	}
	database = *db
	println("exit connection....")
}

func (JDBCObject) JDBC_close() {
	println("closing....")
	database.Close()
	println("exit closing...")
}

func MainApp()  {
	o := JDBCObject{}
	o.name = "mytest"
	o.JDBC_connect()
	o.JDBC_close()
}

func Connect_sql()  {
	db,err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()



	rows,err:=db.Query("SELECT * FROM account")
	if err != nil {
		fmt.Println(err)
		return
	}

	type account struct {
		id int
		name string
		money int
	}

	for rows.Next(){
		var a = account{}
		//rows.Scan(&id,&name,&money)
		rows.Scan(&a.id,&a.name,&a.money)
		//Select(a)
		//fmt.Println(id,"\t",name,"\t",money)
	}
}

func Select(db *sql.DB, talbe JDBCObject,obj interface{}){
	var sql = "SELECT * from "+talbe.name
	rows,err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next(){

	}
}
