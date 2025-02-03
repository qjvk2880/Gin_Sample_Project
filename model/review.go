package model

type Review struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	AlbumID uint   `json:"album_id"`
	Album   Album
}
