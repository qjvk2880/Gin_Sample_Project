package repository

import (
	"example/web-service-gin/config"
	"example/web-service-gin/model"
	"fmt"
	"gorm.io/gorm"
)

type AlbumRdbRepository struct{}

func NewAlbumRdbRepository() *AlbumRdbRepository {
	return &AlbumRdbRepository{}
}

func (r AlbumRdbRepository) GetAll() ([]model.Album, error) {
	var albums []model.Album
	result := config.DB.Find(&albums) // SELECT * FROM albums;

	if result.Error != nil {
		return nil, result.Error
	}
	return albums, nil
}

func (r AlbumRdbRepository) GetByID(id uint) (*model.Album, error) {
	var album model.Album
	result := config.DB.First(&album, id) // SELECT * FROM albums WHERE id = ? LIMIT 1;

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("album with ID %d not found", id)
		}
		return nil, result.Error
	}
	return &album, nil
}

func (r AlbumRdbRepository) Save(album model.Album) error {
	if config.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	result := config.DB.Create(&album)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
