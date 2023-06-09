package services

import (
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/models"
)

type featureService struct {
	db *gorm.DB
}

func (s *featureService) GetAllFeatures() ([]models.HugoFeature, error) {
	var f []models.HugoFeature
	err := s.db.Debug().
		Preload("Next").
		Preload("Prev").
		Find(&f).Error

	return f, err
}

func NewFeatureService(db *gorm.DB) FeatureService {
	return &featureService{db: db}
}

/*
type FeatureLinkUpdateOperation int

const (
	AddHighway FeatureLinkUpdateOperation = iota
	RemoveHighway
	ReverseDirection
)

type FeatureLinkUpdate struct {
	Operation     FeatureLinkUpdateOperation
	Highway       uint
	IsDescending  bool
	FeatureLinkId uint
}


func (s *featureService) UpdateFeatureAdminArea(feature models.Feature, state models.AdminAreaState, country models.AdminAreaCountry) (models.Feature, error) {
	return feature, s.db.Transaction(func(tx *gorm.DB) error {
		return s.db.Debug().
			Model(&feature).
			Select("AdminAreaCountryId", "AdminAreaStateId").
			Updates(models.Feature{AdminAreaCountryId: country.ID, AdminAreaStateId: state.ID}).Error
	})
}

func (s *featureService) GetFeatureLink(id uint) (models.FeatureLink, error) {
	var fl models.FeatureLink
	err := s.db.Debug().Preload("Highways").First(&fl, id).Error

	return fl, err
}

func (s *featureService) DeleteFeatureLink(feature models.FeatureLink) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := s.db.Debug().Delete(&models.FeatureLinkHighway{}, "feature_link_id = ?", feature.ID).Error
		if err != nil {
			return err
		}

		err = s.db.Debug().Delete(&feature).Error

		return err
	})
}

func (s *featureService) UpdateFeatureLink(update FeatureLinkUpdate) (models.FeatureLink, error) {
	var fl models.FeatureLink

	fl, err := s.GetFeatureLink(update.FeatureLinkId)

	if err != nil {
		return fl, err
	}

	if update.Operation == AddHighway {
		for _, v := range fl.Highways {
			if v.ID == update.Highway {
				return fl, fmt.Errorf("highway already exists")
			}
		}

		flh := models.FeatureLinkHighway{FeatureLinkId: fl.ID, HighwayId: update.Highway, IsDescending: update.IsDescending}
		err = s.db.Debug().Create(&flh).Error
		if err != nil {
			return fl, err
		}

		fl, err = s.GetFeatureLink(fl.ID)
		return fl, err
	}

	if update.Operation == RemoveHighway {
		for _, v := range fl.Highways {
			if v.ID == update.Highway {
				err = s.db.Debug().Delete(&models.FeatureLinkHighway{}, "highway_id = ? AND feature_link_id = ?", update.Highway, update.FeatureLinkId).Error
				if err != nil {
					return fl, err
				}
				newfl, err := s.GetFeatureLink(fl.ID)
				return newfl, err

			}
		}

		return fl, fmt.Errorf("highway not found in feature link")
	}

	if update.Operation == ReverseDirection {
		for _, v := range fl.Highways {
			if v.ID == update.Highway {
				flh := models.FeatureLinkHighway{}
				err := s.db.Debug().Where("highway_id = ? AND feature_link_id = ?", update.Highway, update.FeatureLinkId).First(&flh).Error

				if err != nil {
					return fl, err
				}

				err = s.db.Debug().Model(&flh).Update("is_descending", !flh.IsDescending).Error
				if err != nil {
					return fl, err
				}
				newfl, err := s.GetFeatureLink(fl.ID)
				return newfl, err

			}
		}

		return fl, fmt.Errorf("highway not found in feature link")
	}

	return fl, fmt.Errorf("incorrect operation")

}

func (s *featureService) GetFeature(featureId uint) (models.Feature, error) {
	f := models.Feature{}
	err := s.db.Debug().
		Preload("Signs").
		Preload("Next.FromFeature").
		Preload("Next.ToFeature").
		Preload("Next.Highways").
		Preload("Prev.FromFeature").
		Preload("Prev.ToFeature").
		Preload("Prev.Highways").
		First(&f, featureId).Error

	return f, err
}



func (d CreateFeatureLinkDetailsDto) IsEmpty() bool {
	return d.Name == "" && len(d.HighwayIds) == 0 && d.Feature.IsEmpty()
}

func (s *featureService) CreateFeature(title string, latitude float64, longitude float64, countryId uint, stateId uint) (models.Feature, error) {

	pt := models.Point{
		X: longitude,
		Y: latitude,
	}

	feat := models.Feature{
		Point:              pt,
		Name:               title,
		AdminAreaCountryId: countryId,
		AdminAreaStateId:   stateId,
	}

	err := s.db.Debug().Create(&feat).Error

	return feat, err
}

func (s *featureService) CreateFeatureLink(fromFeatureId uint, toFeatureId uint) (models.FeatureLink, error) {

	featLink := models.FeatureLink{
		FromFeatureId: fromFeatureId,
		ToFeatureId:   toFeatureId,
	}

	err := s.db.Debug().Create(&featLink).Error

	return featLink, err
}

func (s *featureService) UpdateBeginAndEnd(featureLinkId uint, fromFeature uint, toFeature uint) (models.FeatureLink, error) {
	featLink := models.FeatureLink{ID: featureLinkId}

	err := s.db.Debug().
		Model(&featLink).
		Select("FromFeatureId", "ToFeatureId").
		Updates(models.FeatureLink{FromFeatureId: fromFeature, ToFeatureId: toFeature}).Error

	return featLink, err
}

func (s *featureService) GetFeatureLinkHighways(featureLinkId uint) ([]models.FeatureLinkHighway, error) {
	var flh []models.FeatureLinkHighway
	err := s.db.Where(&models.FeatureLinkHighway{FeatureLinkId: featureLinkId}).Find(&flh).Error

	return flh, err
}

func (s *featureService) AddHighwayToFeatureLink(newFeatureLink uint, highwayId uint, isDescending bool) (models.FeatureLinkHighway, error) {
	flh := models.FeatureLinkHighway{
		FeatureLinkId: newFeatureLink,
		HighwayId:     highwayId,
		IsDescending:  isDescending,
	}
	err := s.db.Create(&flh).Error

	return flh, err
}

func (s *featureService) RemoveHighwayFromFeatureLink(featureLinkId uint, highwayId uint) error {
	return s.db.Delete(&models.FeatureLinkHighway{}, "highway_id = ? AND feature_link_id = ?", highwayId, featureLinkId).Error
}
*/
