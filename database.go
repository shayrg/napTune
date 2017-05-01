package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dbString string = "root:@tcp(localhost:3306)" +
	"/napTune?charset=utf8"

func GetAllSongs() SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query("select * from songs")
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}
func GetSongById(songId string) SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select * from songs where id = ?")
	checkErr(err)
	rows, err := stmt.Query(songId)
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}
func buildSongsObject(rows *sql.Rows) SongsObject {
	var songs SongsObject
	for rows.Next() {
		var id string
		var name string
		var artist string
		var length string
		var location string
		err := rows.Scan(&id, &name, &artist, &length, &location)
		checkErr(err)
		song := SongObject{
			Id:       id,
			Name:     name,
			Artist:   artist,
			Length:   length,
			Location: location,
		}
		songs = append(songs, song)
	}
	return songs
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
