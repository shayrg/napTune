package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllSongs() SongsStruct {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)"+
		"/napTune?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select * from songs")
	checkErr(err)
	var mysongs SongsStruct
	for rows.Next() {
		var id string
		var name string
		var artist string
		var length string
		var location string
		err = rows.Scan(&id, &name, &artist, &length, &location)
		checkErr(err)
		song := SongStruct{
			Id:       id,
			Name:     name,
			Artist:   artist,
			Length:   length,
			Location: location,
		}
		mysongs = append(mysongs, song)
	}
	return mysongs
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
