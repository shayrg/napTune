package main

import (
	"database/sql"
	//"fmt"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func test() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)"+
		"/napTune?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select * from songs")
	checkErr(err)
	for rows.Next() {
		var id string
		var name string
		var artist string
		var length string
		var location string
		err = rows.Scan(&id, &name, &artist, &length, &location)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(artist)
		fmt.Println(length)
		fmt.Println(location)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
