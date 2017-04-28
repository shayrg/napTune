package main

type SongStruct struct {
	Name     string `json:"name"`
	Length   string `json:"length"`
	Location string `json:"location"`
	Order    string `json:"order"`
}



type SongsStruct []SongStruct
