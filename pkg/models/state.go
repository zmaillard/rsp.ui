package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
)

type State struct {
	ID              uint           `gorm:"column:id;primaryKey"`
	Name            string         `gorm:"column:state_name"`
	Slug            string         `gorm:"column:state_slug"`
	SubdivisionName string         `gorm:"column:subdivision_name"`
	ImageCount      int            `gorm:"column:image_count"`
	Highways        pq.StringArray `gorm:"column:highways;type:text[]"`
	Featured        string         `gorm:"column:featured"`
	CountrySlug     string         `gorm:"column:country_slug"`
	Counties        JSON           `gorm:"column:counties"`
	Places          JSON           `gorm:"column:places"`
}

func (State) TableName() string {
	return "vwhugostate"
}

type adminArea struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (s State) ConvertToDto() generator.Generator {
	var counties []adminArea
	_ = json.Unmarshal(s.Counties, &counties)

	var places []adminArea
	_ = json.Unmarshal(s.Places, &places)
	stateDto := dto.AdminAreaStateDto{
		Name:            s.Name,
		Slug:            s.Slug,
		SubdivisionName: s.SubdivisionName,
		ImageCount:      s.ImageCount,
		Layout:          "state",
		Highways:        s.Highways,
		CountrySlug:     s.CountrySlug,
		Featured:        s.Featured,
		Counties: util.SliceMap(counties, func(county adminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: county.Name,
				Slug: county.Slug,
			}
		}),
		Places: util.SliceMap(places, func(place adminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: place.Name,
				Slug: place.Slug,
			}
		}),
	}

	return stateDto
}
