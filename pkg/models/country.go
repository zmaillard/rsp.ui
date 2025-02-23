package models

import (
	"encoding/json"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
)

type Country struct {
	ID              uint   `gorm:"column:id;primaryKey"`
	Name            string `gorm:"column:country_name"`
	Slug            string `gorm:"column:country_slug"`
	SubdivisionName string `gorm:"column:subdivision_name"`
	ImageCount      int    `gorm:"column:image_count"`
	States          JSON   `gorm:"column:states"`
	Featured        string `gorm:"column:featured"`
	HighwayTypes    JSON   `gorm:"column:highway_types"`
}

func (Country) TableName() string {
	return "sign.vwhugocountry"
}

func (c Country) ConvertToDto() generator.Generator {
	var states []AdminArea
	_ = json.Unmarshal(c.States, &states)

	var highwayTypes []AdminArea
	_ = json.Unmarshal(c.HighwayTypes, &highwayTypes)

	return dto.AdminAreaCountryDto{
		Name:            c.Name,
		Slug:            c.Slug,
		SubdivisionName: c.SubdivisionName,
		ImageCount:      c.ImageCount,
		Featured:        c.Featured,
		States: util.SliceMap(states, func(state AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: state.Name,
				Slug: state.Slug,
			}
		}),
		HighwayTypes: util.SliceMap(highwayTypes, func(ht AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: ht.Name,
				Slug: ht.Slug,
			}
		}),
	}
}
