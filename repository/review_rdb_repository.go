package repository

import (
	"example/web-service-gin/config"
	"example/web-service-gin/model"
)

type ReviewRdbRepository struct{}

func NewReviewRdbRepository() *ReviewRdbRepository {
	return &ReviewRdbRepository{}
}

func (r ReviewRdbRepository) GetAllByAlbum(albumId uint) ([]model.Review, error) {
	var reviews []model.Review

	result := config.DB.Preload("Album").Where("album_id = ?", albumId).Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}
