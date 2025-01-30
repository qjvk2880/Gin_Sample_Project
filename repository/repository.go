package repository

import (
	"example/web-service-gin/model"
)

// album 저장소 인터페이스
type Repository interface {
	GetAll() ([]model.Album, error)
	GetByID(id string) (*model.Album, error)
	Save(album model.Album) error
}
