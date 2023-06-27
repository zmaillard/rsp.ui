package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type PlaceService interface {
	GetAllCountries() ([]models.Country, error)
	GetAllStates() ([]models.State, error)
	GetAllCounties() ([]models.County, error)
	GetAllPlaces() ([]models.Place, error)
}

type FeatureService interface {
	GetAllFeatures() ([]models.Feature, error)
}

type HighwayService interface {
	GetAllHighways() ([]models.Highway, error)
	GetAllHighwayTypes() ([]models.HighwayType, error)
}

type SignService interface {
	GetAllSigns() (models.HighwaySigns, error)
}

//go:generate mockgen -destination=mocks/mock_datastore.go -package=mocks . Datastore
type Datastore interface {
	GetSignService() SignService
	GetFeatureService() FeatureService
	GetPlaceService() PlaceService
	GetHighwayService() HighwayService
}

type postgresqlStore struct {
	db *gorm.DB
}

func NewDatastore(db *gorm.DB) Datastore {
	return &postgresqlStore{db}
}

func (s *postgresqlStore) GetFeatureService() FeatureService {
	return NewFeatureService(s.db)
}

func (s *postgresqlStore) GetSignService() SignService {
	return NewSignService(s.db)
}

func (s *postgresqlStore) GetPlaceService() PlaceService {
	return NewPlaceService(s.db)
}

func (s *postgresqlStore) GetHighwayService() HighwayService {
	return NewHighwayService(s.db)
}
