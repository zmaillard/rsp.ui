package models

import (
	"fmt"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
)

type County struct {
	ID         uint   `gorm:"primaryKey:column:id"`
	Name       string `gorm:"column:county_name"`
	Slug       string `gorm:"column:county_slug"`
	ImageCount int    `gorm:"column:image_count"`
	StateName  string `gorm:"column:state_name"`
	StateSlug  string `gorm:"column:state_slug"`
}

func (County) TableName() string {
	return "sign.vwhugocounty"
}

func (c County) ConvertToDto() generator.Generator {
	countyDto := dto.AdminAreaCountyDto{
		Name:       c.Name,
		Slug:       c.Slug,
		ImageCount: c.ImageCount,
		StateSlug:  c.StateSlug,
		Aliases:    []string{fmt.Sprintf("/county/%s/%s", c.StateSlug, c.Slug)},
	}

	return countyDto
}
