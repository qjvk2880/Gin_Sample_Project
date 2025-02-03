package repository

import (
	"example/web-service-gin/config"
	"example/web-service-gin/model"
)

type ReviewRdbRepository struct{}

func NewReviewRdbRepository() *ReviewRdbRepository {
	return &ReviewRdbRepository{}
}

func (r ReviewRdbRepository) GetAllByAlbum(albumId int) ([]model.Review, error) {
	var reviews []model.Review

	result := config.DB.Joins("Album", config.DB.Where(&model.Review{AlbumID: uint(albumId)})).Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}
