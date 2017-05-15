package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

const dbString string = "root:@tcp(localhost:3306)" +
	"/napTune?charset=utf8"

/**
 * Songs table
 */
//Get All
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

//Get
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

//Insert
func InsertSong(song SongObject) string {
	song.Id = getNextId("songs")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into songs values (?, ?, ?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(song.Id, song.Name, song.Artist, song.Length,
		song.Location)
	checkErr(err)
	db.Close()
	return song.Id
}

//Update
func UpdateSong(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update songs set name = ?, artist = ?, " +
		"length = ?, location = ? where id = ? ")
	checkErr(err)
	result, err := stmt.Exec(song.Name, song.Artist, song.Length,
		song.Location, song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

//Delete
func DeleteSong(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from songs where " +
		"id = ?")
	checkErr(err)
	result, err := stmt.Exec(song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

/**
 * Playlists table
 */
//Get All
func GetAllPlaylists() PlaylistsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query("select * from playlists")
	checkErr(err)
	//Build Songs Object
	playlists := buildPlaylistsObject(rows)
	db.Close()
	return playlists
}

//Get
func GetPlaylistById(playlistId string) PlaylistsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select * from playlists where id = ?")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	//Build Songs Object
	playlists := buildPlaylistsObject(rows)
	db.Close()
	return playlists
}

//Insert
func InsertPlaylist(playlist PlaylistObject) string {
	playlist.Id = getNextId("playlists")
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlists values (?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlist.Id, playlist.Name)
	checkErr(err)
	db.Close()
	return playlist.Id
}

//Delete
func DeletePlaylist(playlist PlaylistObject) {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlists where id = ?")
	checkErr(err)
	_, err = stmt.Exec(playlist.Id)
	checkErr(err)
	db.Close()
}

//Update
func UpdatePlaylist(playlist PlaylistObject) {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("update playlists set name = ? where id = ?")
	checkErr(err)
	_, err = stmt.Exec(playlist.Name, playlist.Id)
	checkErr(err)
	db.Close()
}

/**
 * PlaylistSongs Table
 */
//Get
func GetPlaylistSongsById(playlistId string) SongsObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select songs.id, songs.name, artist, length, location from songs join playlistSongs on songs.id = songId join playlists on playlists.id = playlistId where playlistId = ?")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	//Build Songs Object
	songs := buildSongsObject(rows)
	db.Close()
	return songs
}

//Insert
func InsertSongInPlaylist(playlistSong PlaylistSongObject) string {
	playlistSong.SongOrder = getNextPlaylistOrder(playlistSong.PlaylistId)
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("insert into playlistSongs values (?, ?, ?)")
	checkErr(err)
	_, err = stmt.Exec(playlistSong.PlaylistId, playlistSong.SongId,
		playlistSong.SongOrder)
	checkErr(err)
	db.Close()
	return playlistSong.SongOrder
}

//Delete
func DeleteSongInPlaylist(playlistSong PlaylistSongObject) {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"playlistId = ? and songId = ?")
	checkErr(err)
	_, err = stmt.Exec(playlistSong.PlaylistId, playlistSong.SongId)
	checkErr(err)
	db.Close()
}
func DeleteAllSongsInPlaylist(playlist PlaylistObject) {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"playlistId = ?")
	checkErr(err)
	_, err = stmt.Exec(playlist.Id)
	checkErr(err)
	db.Close()
}
func DeleteSongFromAllPlaylists(song SongObject) int64 {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("delete from playlistSongs where " +
		"songId = ?")
	checkErr(err)
	result, err := stmt.Exec(song.Id)
	checkErr(err)
	db.Close()
	rowsAffected, err := result.RowsAffected()
	return rowsAffected
}

/**
 * User Table
 */
//Get
func GetLogin(userLogin LoginObject) LoginObject {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select email, password from users where email = ?")
	checkErr(err)
	rows, err := stmt.Query(userLogin.Email)
	checkErr(err)
	login := buildLogin(rows)
	db.Close()
	return login
}

/**
 * Helper functions
 */
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getNextId(table string) string {
	queryString := "select id from " + table + " order by id desc limit 1"
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	rows, err := db.Query(queryString)
	checkErr(err)
	var id string
	for rows.Next() {
		var rowId string
		rows.Scan(&rowId)
		id = rowId
	}
	incrementString(&id)
	db.Close()
	return id
}
func getNextPlaylistOrder(playlistId string) string {
	db, err := sql.Open("mysql", dbString)
	checkErr(err)
	stmt, err := db.Prepare("select songOrder from playlistSongs where " +
		"playlistId = ? order by songOrder desc limit 1")
	checkErr(err)
	rows, err := stmt.Query(playlistId)
	checkErr(err)
	var songOrder string
	for rows.Next() {
		var rowSongOrder string
		rows.Scan(&rowSongOrder)
		songOrder = rowSongOrder
	}
	incrementString(&songOrder)
	db.Close()
	return songOrder
}
func incrementString(string *string) {
	myInt, _ := strconv.ParseInt(*string, 0, 64)
	myInt++
	*string = fmt.Sprintf("%04d", myInt)
}
