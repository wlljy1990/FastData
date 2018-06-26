package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db,err :=sql.Open("mysql","root:123456@tcp(127.0.0.1)/?charset=utf8")
	checkErr(err)
	db.Query("drop database if exists tmpdb")
	db.Query("create database tmpdb")
	db.Query("use tmpdb")
	db.Query("create table tmpdb.tmptab(c1 int,c2 varchar(20),c3 varchar(20))")
	db.Query("insert into tmpdb.tmptab values(100,'test','what'),(100,'test1','what1')")

	query,err := db.Query("SELECT * from tmpdb.tmpdb")
	checkErr(err)
	v := reflect.ValueOf(query)
	fmt.Println(v)

	db.Close()
}