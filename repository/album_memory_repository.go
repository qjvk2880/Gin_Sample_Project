package repository

import (
	"errors"
	"example/web-service-gin/model"
)

// In-Memory 저장소 구현
type AlbumMemoryRepository struct {
	albums []model.Album
}

func NewMemoryRepository() *AlbumMemoryRepository {
	var albums = []model.Album{
		{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &AlbumMemoryRepository{albums: albums}
}

func (r *AlbumMemoryRepository) GetAll() ([]model.Album, error) {
	return r.albums, nil
}

func (r *AlbumMemoryRepository) GetByID(id uint) (*model.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("album not found")
}

func (r *AlbumMemoryRepository) Save(album model.Album) error {
	r.albums = append(r.albums, album)
	return nil
}
