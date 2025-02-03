package repository

import "example/web-service-gin/model"

type ReviewRepository interface {
	GetAllByAlbum(albumId uint) ([]model.Review, error)
}
