package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type featureService struct {
	db *gorm.DB
}

func (s *featureService) GetAllFeatures() ([]models.Feature, error) {
	var f []models.Feature
	err := s.db.Debug().
		Preload("Next").
		Preload("Prev").
		Find(&f).Error

	return f, err
}

func NewFeatureService(db *gorm.DB) FeatureService {
	return &featureService{db: db}
}
