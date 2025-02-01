package model

type Review struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	AlbumID uint
	Album   Album
}
