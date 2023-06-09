package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
)

type HugoHighwayType struct {
	ID         uint           `gorm:"column:id;primaryKey"`
	Name       string         `gorm:"column:highway_type_name"`
	Slug       string         `gorm:"column:highway_type_slug"`
	Sort       int            `gorm:"column:sort" `
	ImageCount int            `gorm:"column:imagecount"`
	Featured   types.ImageID  `gorm:"column:imageid" `
	Highways   pq.StringArray `gorm:"column:highways;type:[]string"`
}

func (HugoHighwayType) TableName() string {
	return "vwhugohighwaytype"
}

func (ht HugoHighwayType) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayTypeDto{
		Name:            ht.Name,
		Slug:            ht.Slug,
		Sort:            ht.Sort,
		ImageCount:      ht.ImageCount,
		Featured:        ht.Featured.String(),
		HighwayTaxomomy: ht.Highways,
	}

	return highwayDto
}

type HighwayType struct {
	ID             uint          `gorm:"primaryKey" yaml:"-"`
	Name           string        `gorm:"column:highway_type_name" yaml:"name"`
	Slug           string        `gorm:"column:slug" yaml:"slug"`
	Sort           int           `gorm:"column:sort" yaml:"sort"`
	ImageCount     int           `gorm:"column:image_count" yaml:"imageCount"`
	Highways       []Highway     `gorm:"foreignKey:HighwayTypeId" yaml:"-"`
	DisplayImageId int           `gorm:"column:display_image_id" yaml:"-"`
	CountryId      int           `gorm:"column:admin_area_country_id" yaml:"-"`
	Featured       types.ImageID `gorm:"-" yaml:"featured,omitempty"`
}

func (ht *HighwayType) TableName() string {
	return "highway_type"
}

func (ht *HighwayType) AfterFind(tx *gorm.DB) (err error) {
	var sign []HighwaySign
	err = tx.Debug().Find(&sign, ht.DisplayImageId).Error
	if err != nil {
		return err
	}

	if len(sign) > 0 {
		ht.Featured = sign[0].ImageId
	}

	return nil
}

func (ht HighwayType) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayTypeDto{
		Name:       ht.Name,
		Slug:       ht.Slug,
		Sort:       ht.Sort,
		ImageCount: ht.ImageCount,
		Featured:   ht.Featured.String(),
	}

	hwyTaxonomy := make([]string, len(ht.Highways))
	for i, _ := range ht.Highways {
		hwyTaxonomy[i] = ht.Highways[i].Slug
	}

	highwayDto.HighwayTaxomomy = hwyTaxonomy

	return highwayDto
}
