package models

import (
	"github.com/lib/pq"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
)

type HighwayType struct {
	ID         uint           `gorm:"column:id;primaryKey"`
	Name       string         `gorm:"column:highway_type_name"`
	Slug       string         `gorm:"column:highway_type_slug"`
	Sort       int            `gorm:"column:sort" `
	ImageCount int            `gorm:"column:imagecount"`
	Featured   types.ImageID  `gorm:"column:imageid" `
	Highways   pq.StringArray `gorm:"column:highways;type:[]string"`
	Country    string         `gorm:"column:country"`
}

func (HighwayType) TableName() string {
	return "vwhugohighwaytype"
}

func (ht HighwayType) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayTypeDto{
		Name:            ht.Name,
		Slug:            ht.Slug,
		Sort:            ht.Sort,
		ImageCount:      ht.ImageCount,
		Featured:        ht.Featured.String(),
		HighwayTaxomomy: ht.Highways,
		Country:         ht.Country,
	}

	return highwayDto
}
