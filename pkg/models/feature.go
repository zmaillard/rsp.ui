package models

import (
	"github.com/lib/pq"
	dto2 "highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
	"strconv"
)

type Feature struct {
	ID          uint           `gorm:"column:id;primaryKey"`
	Point       Point          `gorm:"column:point"`
	Name        string         `gorm:"column:name"`
	Next        []FeatureLink  `gorm:"foreignKey:FromFeatureId"`
	Prev        []FeatureLink  `gorm:"foreignKey:ToFeatureId"`
	Signs       pq.StringArray `gorm:"column:signs;type:text[]"`
	StateName   string         `gorm:"column:state_name"`
	StateSlug   string         `gorm:"column:state_slug"`
	CountryName string         `gorm:"column:country_name"`
	CountrySlug string         `gorm:"column:country_slug"`
}

func (Feature) TableName() string {
	return "sign.vwhugofeature"
}

func (f Feature) X() float64 {
	return f.Point.X
}
func (f Feature) Y() float64 {
	return f.Point.Y
}
func (f Feature) Id() string {
	return strconv.Itoa(int(f.ID))
}

func (f Feature) IsEmpty() bool {
	return f.ID == 0 && f.Name == "" && len(f.Next) == 0 && len(f.Prev) == 0 && len(f.Signs) == 0
}

func (f Feature) ConvertToDto() generator.Generator {
	feature := dto2.FeatureDto{
		ID:   f.ID,
		Name: f.Name,
		Next: util.SliceMap(f.Next, func(link FeatureLink) dto2.FeatureLinkDto {
			return dto2.FeatureLinkDto{
				ID:            link.ID,
				RoadName:      link.RoadName,
				FromFeatureId: link.FromFeatureId,
				ToFeatureId:   link.ToFeatureId,
				Highways:      link.Highways,
			}
		}),
		Prev: util.SliceMap(f.Prev, func(link FeatureLink) dto2.FeatureLinkDto {
			return dto2.FeatureLinkDto{
				ID:            link.ID,
				RoadName:      link.RoadName,
				FromFeatureId: link.FromFeatureId,
				ToFeatureId:   link.ToFeatureId,
				Highways:      link.Highways,
			}
		}),
		Signs: f.Signs,
		State: struct {
			Name string
			Slug string
		}{
			Name: f.StateName,
			Slug: f.StateSlug,
		},
		Country: struct {
			Name string
			Slug string
		}{
			Name: f.CountryName,
			Slug: f.CountrySlug,
		},
	}

	return feature
}
