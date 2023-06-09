package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type placeService struct {
	db *gorm.DB
}

func NewPlaceService(db *gorm.DB) PlaceService {
	return &placeService{db: db}
}

func (p *placeService) GetAllCountries() ([]models.HugoCountry, error) {
	var countries []models.HugoCountry
	err := p.db.Debug().Find(&countries).Error

	return countries, err
}

func (p *placeService) GetAllStates() ([]models.HugoState, error) {
	var states []models.HugoState

	query := p.db.Debug()
	err := query.
		Find(&states).
		Error

	if err != nil {
		return states, err
	}

	return states, nil
}

func (p *placeService) GetAllPlaces() ([]models.HugoPlace, error) {
	var places []models.HugoPlace
	err := p.db.Debug().
		Where("image_count > 0").
		Find(&places).
		Error

	if err != nil {
		return places, err
	}

	return places, nil
}

func (p *placeService) GetAllCounties() ([]models.HugoCounty, error) {
	var counties []models.HugoCounty
	err := p.db.Debug().
		Find(&counties).
		Error

	if err != nil {
		return counties, err
	}

	return counties, nil
}
