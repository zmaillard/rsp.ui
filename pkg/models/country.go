package models

import (
	"encoding/json"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
)

type HugoCountry struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	Name         string `gorm:"column:country_name"`
	Slug         string `gorm:"column:country_slug"`
	ImageCount   int    `gorm:"column:image_count"`
	States       JSON   `gorm:"column:states"`
	Featured     string `gorm:"column:featured"`
	HighwayTypes JSON   `gorm:"column:highway_types"`
}

func (HugoCountry) TableName() string {
	return "vwhugocountry"
}

func (c HugoCountry) ConvertToDto() generator.Generator {
	var states []adminArea
	_ = json.Unmarshal(c.States, &states)

	var highwayTypes []adminArea
	_ = json.Unmarshal(c.HighwayTypes, &highwayTypes)

	return dto.AdminAreaCountryDto{
		Name:            c.Name,
		Slug:            c.Slug,
		SubdivisionName: "", //TODO::Add support for subdivision type
		ImageCount:      c.ImageCount,
		Featured:        c.Featured,
		States: util.SliceMap(states, func(state adminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: state.Name,
				Slug: state.Slug,
			}
		}),
		HighwayTypes: util.SliceMap(highwayTypes, func(ht adminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: ht.Name,
				Slug: ht.Slug,
			}
		}),
	}
}
