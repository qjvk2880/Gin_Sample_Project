package logic

import (
	"errors"
	"example/web-service-gin/repository"
)

type Logic struct {
	repo repository.Repository
}

func NewLogic(repo repository.Repository) *Logic {
	return &Logic{repo: repo}
}

func (l *Logic) GetAllAlbums() ([]repository.Album, error) {
	return l.repo.GetAll()
}

func (l *Logic) GetAlbumByID(id string) (*repository.Album, error) {
	return l.repo.GetByID(id)
}

func (l *Logic) SaveAlbum(album repository.Album) error {
	if album.Title == "" {
		return errors.New("album title cannot be empty")
	}

	return l.repo.Save(album)
}
