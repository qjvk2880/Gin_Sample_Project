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

	result := config.DB.Find(&reviews, "album_id = ?", albumId)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func (r ReviewRdbRepository) Save(review model.Review) error {
	result := config.DB.Create(&review)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
