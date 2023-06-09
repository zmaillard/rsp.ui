package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"highway-sign-portal-builder/pkg/util"
)

type AdminAreaCountry struct {
	ID              uint             `gorm:"column:id;primaryKey"`
	Name            string           `gorm:"column:name"`
	Slug            string           `gorm:"column:slug"`
	SubdivisionName string           `gorm:"column:subdivision_name"`
	ImageCount      int              `gorm:"image_count"`
	States          []AdminAreaState `gorm:"foreignKey:CountryId"`
	HighwayTypes    []HighwayType    `gorm:"foreignKey:CountryId"`
	DisplayImageId  *int             `gorm:"column:featured_sign_id"`
	Featured        types.ImageID    `gorm:"-"`
}

func (AdminAreaCountry) TableName() string {
	return "admin_area_country"
}

func (AdminAreaCountry) SchemaName() string {
	return "sign"
}

func (c AdminAreaCountry) ConvertToDto() generator.Generator {
	dto := dto.AdminAreaCountryDto{
		Name:            c.Name,
		Slug:            c.Slug,
		SubdivisionName: c.SubdivisionName,
		ImageCount:      c.ImageCount,
		Featured:        c.Featured.String(),
		States: util.SliceMap(c.States, func(state AdminAreaState) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: state.Name,
				Slug: state.Slug,
			}
		}),
		HighwayTypes: util.SliceMap(c.HighwayTypes, func(ht HighwayType) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: ht.Name,
				Slug: ht.Slug,
			}
		}),
	}

	return dto
}
func (c *AdminAreaCountry) AfterFind(tx *gorm.DB) (err error) {
	var sign HighwaySign

	if c.DisplayImageId != nil {
		err = tx.Debug().First(&sign, c.DisplayImageId).Error
		if err != nil {
			return err
		}

		c.Featured = sign.ImageId
	}

	if len(c.States) == 0 {
		var states []AdminAreaState
		err = tx.Where("adminarea_country_id = ?", c.ID).Find(&states).Error
		if err != nil {
			return err
		}
		c.States = states
	}

	if len(c.HighwayTypes) == 0 {
		var highwayTypes []HighwayType
		err = tx.Where("admin_area_country_id = ?", c.ID).Find(&highwayTypes).Error
		if err != nil {
			return err
		}
		c.HighwayTypes = highwayTypes
	}

	return nil
}

type Countries []AdminAreaCountry

func (c Countries) GetLookup() ([]byte, error) {
	nameSlug := make(map[string]string)

	for _, v := range c {
		nameSlug[v.Name] = v.Slug
	}

	return json.Marshal(nameSlug)
}

func (Countries) OutLookupFile() string {
	return "data/country.json"
}

type AdminAreaState struct {
	ID              uint              `gorm:"primaryKey"`
	Name            string            `gorm:"column:name" `
	Slug            string            `gorm:"column:slug" `
	SubdivisionName string            `gorm:"column:subdivision_name" `
	CountryId       uint              `gorm:"column:adminarea_country_id"`
	ImageCount      int               `gorm:"image_count"`
	Highways        []string          `gorm:"-"`
	Counties        []AdminAreaCounty `gorm:"foreignKey:StateId"`
	Places          []AdminAreaPlace  `gorm:"foreignKey:StateId"`
	DisplayImageId  *int              `gorm:"column:featured_sign_id"`
	Featured        types.ImageID     `gorm:"-"`
	Country         AdminAreaCountry  `gorm:"foreignKey:CountryId"`
}

func (s AdminAreaState) ConvertToDto() generator.Generator {
	stateDto := dto.AdminAreaStateDto{
		Name:            s.Name,
		Slug:            s.Slug,
		SubdivisionName: s.SubdivisionName,
		ImageCount:      s.ImageCount,
		Layout:          "state",
		Highways:        s.Highways,
		CountrySlug:     s.Country.Slug,
		Featured:        s.Featured.String(),
		Counties: util.SliceMap(s.Counties, func(county AdminAreaCounty) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: county.Name,
				Slug: county.Slug,
			}
		}),
		Places: util.SliceMap(s.Places, func(place AdminAreaPlace) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: place.Name,
				Slug: place.Slug,
			}
		}),
	}

	return stateDto
}

func (s *AdminAreaState) AfterFind(tx *gorm.DB) (err error) {
	var sign []HighwaySign

	if s.DisplayImageId != nil {
		err = tx.Debug().Find(&sign, s.DisplayImageId).Error
		if err != nil {
			return err
		}

		if len(sign) > 0 {
			s.Featured = sign[0].ImageId
		}
	}

	var hwyslugs []string
	err = tx.Debug().
		Raw("SELECT DISTINCT h.slug FROM highwaysign_highway hs INNER JOIN highwaysign sign ON hs.highwaysign_id = sign.id INNER JOIN highway h on hs.highway_id = h.id WHERE sign.admin_area_state_id = ? AND hs.is_to = false", s.ID).
		Find(&hwyslugs).
		Error

	s.Highways = hwyslugs
	if err != nil {
		return err
	}

	return nil
}

func (AdminAreaState) TableName() string {
	return "admin_area_state"
}

type AdminAreaPlace struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `gorm:"column:name"`
	Slug       string         `gorm:"column:slug"`
	StateId    uint           `gorm:"column:admin_area_stateid" `
	ImageCount int            `gorm:"image_count"`
	State      AdminAreaState `gorm:"foreignKey:StateId"`
}

func (s AdminAreaPlace) ConvertToDto() generator.Generator {
	placeDto := dto.AdminAreaPlaceDto{
		Name:       s.Name,
		Slug:       s.Slug,
		ImageCount: s.ImageCount,
		StateSlug:  s.State.Slug,
	}

	return placeDto
}

func (AdminAreaPlace) TableName() string {
	return "admin_area_place"
}

type AdminAreaCounty struct {
	ID         uint           `gorm:"primaryKey" yaml:"-"`
	Name       string         `gorm:"column:name" yaml:"name"`
	Slug       string         `gorm:"column:slug" yaml:"slug"`
	StateId    uint           `gorm:"column:admin_area_stateid" yaml:"-"`
	ImageCount int            `gorm:"image_count"`
	State      AdminAreaState `gorm:"foreignKey:StateId"`
}

func (c AdminAreaCounty) ConvertToDto() generator.Generator {
	countyDto := dto.AdminAreaCountyDto{
		Name:       c.Name,
		Slug:       c.Slug,
		ImageCount: c.ImageCount,
		StateSlug:  c.State.Slug,
		Aliases:    []string{fmt.Sprintf("/county/%s/%s", c.State.Slug, c.Slug)},
	}

	return countyDto
}

func (AdminAreaCounty) TableName() string {
	return "admin_area_county"
}
