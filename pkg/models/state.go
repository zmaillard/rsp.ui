package models

import (
	"encoding/json"
	"fmt"
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
	Categories      pq.StringArray `gorm:"column:categories;type:text[]"`
}

func (State) TableName() string {
	return "sign.vwhugostate"
}

type AdminArea struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (s State) ConvertToDto() generator.Generator {
	var counties []AdminArea
	_ = json.Unmarshal(s.Counties, &counties)

	var places []AdminArea
	_ = json.Unmarshal(s.Places, &places)
	stateDto := dto.AdminAreaStateDto{
		Name:            s.Name,
		Slug:            s.Slug,
		SubdivisionName: s.SubdivisionName,
		ImageCount:      s.ImageCount,
		Layout:          "state",
		Output:          []string{"html", "list"},
		Highways:        s.Highways,
		CountrySlug:     s.CountrySlug,
		Featured:        s.Featured,
		Counties: util.SliceMap(counties, func(county AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: county.Name,
				Slug: county.Slug,
			}
		}),
		Places: util.SliceMap(places, func(place AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: place.Name,
				Slug: place.Slug,
			}
		}),
	}

	var categories []string
	for _, v := range s.Categories {
		stateCat := fmt.Sprintf("%s_%s", s.Slug, v)
		categories = append(categories, stateCat)
	}

	stateDto.StateCategories = categories

	return stateDto
}
