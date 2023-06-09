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

func (h *highwayService) GetAllHighways() ([]models.HugoHighway, error) {
	var highways []models.HugoHighway
	err := h.db.Debug().Find(&highways).Error
	if err != nil {
		return highways, err
	}
	return highways, nil
}

func (h *highwayService) GetAllHighwayTypes() ([]models.HugoHighwayType, error) {
	var highwayTypes []models.HugoHighwayType
	err := h.db.Debug().
		Find(&highwayTypes).
		Error

	if err != nil {
		return highwayTypes, err
	}

	return highwayTypes, nil
}

/*

type CreateHighway struct {
	Name          string
	Slug          string
	HighwayTypeId uint
	ScopeId       uint
	ImageName     string
	SortNumber    int
	CountryId     uint
	StateId       *uint
}


func (h *highwayService) GetScopes() ([]models.Scope, error) {
	var scopes []models.Scope
	err := h.db.Debug().Find(&scopes).Error
	if err != nil {
		return scopes, err
	}
	return scopes, nil
}

func (h *highwayService) GetScope(id uint) (models.Scope, error) {
	var scope models.Scope
	err := h.db.Debug().First(&scope, id).Error
	if err != nil {
		return scope, err
	}
	return scope, nil
}

func (h *highwayService) GetScopeByName(scopeName string) (models.Scope, error) {
	var scope models.Scope
	err := h.db.Debug().Where("scope = ?", scopeName).First(&scope).Error
	if err != nil {
		return scope, err
	}
	return scope, nil
}

func (h *highwayService) GetHighway(id uint) (models.Highway, error) {
	var highway models.Highway
	err := h.db.Debug().Preload("HighwayType").First(&highway, id).Error
	if err != nil {
		return highway, err
	}
	return highway, nil
}



func (h *highwayService) GetHighwayType(id uint) (models.HighwayType, error) {
	var highwayType models.HighwayType
	err := h.db.Debug().Preload("Highways").First(&highwayType, id).Error
	if err != nil {
		return highwayType, err
	}
	return highwayType, nil
}

func (h *highwayService) GetHighwayTypeByName(highwayTypeName string) (models.HighwayType, error) {
	var highwayType models.HighwayType
	err := h.db.Debug().Where("highway_type_name = ?", highwayTypeName).First(&highwayType).Error
	if err != nil {
		return highwayType, err
	}
	return highwayType, nil
}



func (h *highwayService) CreateHighway(highway CreateHighway) (models.Highway, error) {
	newHighway := models.Highway{
		Name:          highway.Name,
		Slug:          highway.Slug,
		Sort:          highway.SortNumber,
		Image:         highway.ImageName,
		HighwayTypeId: highway.HighwayTypeId,
		ScopeId:       highway.ScopeId,
		CountryId:     highway.CountryId,
		StateId:       highway.StateId,
	}

	err := h.db.Debug().
		Create(&newHighway).
		Error

	if err != nil {
		return newHighway, err
	}

	return newHighway, nil
}

type ReplaceHighwaysDto struct {
	FromHighway models.Highway
	ToHighway   models.Highway
	Signs       []models.HighwaySign
}

func (h *highwayService) ReplaceHighways(highways ReplaceHighwaysDto) error {
	return h.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range highways.Signs {
			err := h.db.Debug().Model(&models.HighwaySorting{}).
				Where("highwaysign_id = ? AND highway_id = ?", v.ID, highways.FromHighway.ID).
				Update("highway_id", highways.ToHighway.ID).
				Error

			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *highwayService) InsertHighwaySorting(signId uint, highwayId uint, isTo bool) error {
	hwy := models.HighwaySorting{
		SignId:    signId,
		HighwayId: highwayId,
		IsTo:      isTo,
	}

	return h.db.Debug().Create(&hwy).Error
}

func (h *highwayService) GetHighwaysForStateAndCountry(stateId uint, countryId uint) ([]models.Highway, error) {
	var highways []models.Highway

	err := h.db.Debug().
		Where("admin_area_state_id = ? OR (admin_area_state_id is null AND admin_area_country_id = ?)", stateId, countryId).
		Order("highway_name").
		Find(&highways).Error

	return highways, err
}

*/
