package models

import (
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
)

type Place struct {
	ID         uint   `gorm:"primaryKey:column:id"`
	Name       string `gorm:"column:place_name"`
	Slug       string `gorm:"column:place_slug"`
	ImageCount int    `gorm:"image_count"`
	StateName  string `gorm:"column:state_name"`
	StateSlug  string `gorm:"column:state_slug"`
}

func (Place) TableName() string {
	return "sign.vwhugoplace"
}

func (s Place) ConvertToDto() generator.Generator {
	placeDto := dto.AdminAreaPlaceDto{
		Name:       s.Name,
		Slug:       s.Slug,
		ImageCount: s.ImageCount,
		StateSlug:  s.StateSlug,
	}

	return placeDto
}
