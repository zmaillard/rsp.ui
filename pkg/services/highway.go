package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type highwayService struct {
	db *gorm.DB
}

func NewHighwayService(db *gorm.DB) HighwayService {
	return &highwayService{db: db}
}

func (h *highwayService) GetAllHighways() ([]models.Highway, error) {
	var highways []models.Highway
	err := h.db.Debug().Find(&highways).Error
	if err != nil {
		return highways, err
	}
	return highways, nil
}

func (h *highwayService) GetAllHighwayTypes() ([]models.HighwayType, error) {
	var highwayTypes []models.HighwayType
	err := h.db.Debug().
		Find(&highwayTypes).
		Error

	if err != nil {
		return highwayTypes, err
	}

	return highwayTypes, nil
}
