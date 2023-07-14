package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type signService struct {
	db *gorm.DB
}

func NewSignService(db *gorm.DB) SignService {
	return &signService{db: db}
}

func (s *signService) GetAllSigns() (models.HighwaySigns, error) {
	var hs []models.HighwaySign
	err := s.db.Debug().
		Find(&hs).Error

	return hs, err
}

func (s *signService) GetAllTags() (models.Tags, error) {
	var tags []models.Tag

	err := s.db.Debug().Find(&tags).Error

	return tags, err
}
