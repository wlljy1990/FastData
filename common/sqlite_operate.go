package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err !=nil {
		panic(err)
	}
}


func main () {
	db,err := sql.Open("sqlite3","./test.db")
	checkErr(err)

	sql_table :=`
	CREATE TABLE IF NOT EXISTS test_user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name VARCHAR(64) NULL,
		depart_name VARCHAR(64) NULL,
		created DATE NULL
		);
		`
	db.Exec(sql_table)

	//insert
	stmt,err:= db.Prepare("INSERT INTO test_user(user_name,depart_name,created) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("test","test","2018-06-22")
	id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
	//query
	rows,err := db.Query("SELECT * FROM test_user")
	checkErr(err)

	var uid int 
	var user_name string
	var depart_name string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid,&user_name,&depart_name,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(user_name)
		fmt.Println(depart_name)
		fmt.Println(created)
	}
	rows.Close()
}