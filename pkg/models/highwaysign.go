package models

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mmcloughlin/geohash"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"highway-sign-portal-builder/pkg/util"
	"time"
)

type HighwaySigns []HugoHighwaySign

func (HighwaySigns) OutLookupFile() string {
	return "netlify/edge-functions/common/images.json"
}

func (hs HighwaySigns) GetLookup() ([]byte, error) {
	var images []string

	for _, v := range hs {
		images = append(images, v.ImageId.String())
	}
	res := map[string]interface{}{
		"images": images,
	}

	return json.Marshal(res)
}

type HugoHighwaySign struct {
	ID               uint           `gorm:"column:id;primaryKey"`
	Title            string         `gorm:"column:title"`
	Description      string         `gorm:"column:sign_description"`
	FeatureId        uint           `gorm:"column:feature_id"`
	DateTaken        time.Time      `gorm:"column:date_taken"`
	ImageId          types.ImageID  `gorm:"column:imageid"`
	FlickrId         *string        `gorm:"column:flickrid"`
	Point            Point          `gorm:"column:point"`
	Country          string         `gorm:"column:country_slug"`
	State            string         `gorm:"column:state_slug"`
	Place            *string        `gorm:"column:place_slug"`
	StateSubdivision *string        `gorm:"column:county_slug"`
	Tags             pq.StringArray `gorm:"column:tags;type:text[]"`
	Highways         pq.StringArray `gorm:"column:highways;type:text[]"`
	IsTo             pq.StringArray `gorm:"column:is_to;type:text[]"`
	ImageHeight      int            `gorm:"column:image_height"`
	ImageWidth       int            `gorm:"column:image_width"`
}

func (HugoHighwaySign) TableName() string {
	return "vwhugohighwaysign"
}

func (s HugoHighwaySign) X() float64 {
	return s.Point.X
}
func (s HugoHighwaySign) Y() float64 {
	return s.Point.Y
}
func (s HugoHighwaySign) Id() string {
	return s.ImageId.String()
}

func (s HugoHighwaySign) ConvertToDto() generator.Generator {
	gh := geohash.EncodeWithPrecision(s.Y(), s.X(), 5)
	highwaySignDto := dto.HighwaySignDto{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		FeatureId:   s.FeatureId,
		DateTaken:   s.DateTaken,
		ImageId:     s.ImageId.String(),
		FlickrId:    s.FlickrId,
		Point: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  s.Y(),
			Longitude: s.X(),
		},
		StateSlug:            s.State,
		CountrySlug:          s.Country,
		Recent:               s.DateTaken.Format("2006-01"),
		GeoHash:              gh,
		Tags:                 s.Tags,
		PlaceSlug:            s.Place,
		StateSubdivisionSlug: s.StateSubdivision,
		Highways:             s.Highways,
		ToHighways:           s.IsTo,
		ImageWidth:           s.ImageWidth,
		ImageHeight:          s.ImageHeight,
	}

	return highwaySignDto
}

type HighwaySign struct {
	ID                 uint             `gorm:"primaryKey"`
	Title              string           `gorm:"column:title"`
	Description        string           `gorm:"column:sign_description" yaml:"-"`
	FeatureId          uint             `gorm:"column:feature_id"`
	DateTaken          time.Time        `gorm:"column:date_taken"`
	ImageId            types.ImageID    `gorm:"column:imageid"`
	FlickrId           *string          `gorm:"column:flickrid"`
	ImageHeight        int              `gorm:"column:image_height"`
	ImageWidth         int              `gorm:"column:image_width"`
	Place              *AdminAreaPlace  `gorm:"foreignKey:AdminAreaPlaceId" yaml:",omitempty"`
	Point              Point            `gorm:"column:point"`
	State              AdminAreaState   `gorm:"foreignKey:AdminAreaStateId"`
	Country            AdminAreaCountry `gorm:"foreignKey:AdminAreaCountryId"`
	Highways           []HighwaySorting `gorm:"foreignKey:SignId" yaml:"-"`
	StateSubdivision   *AdminAreaCounty `gorm:"foreignKey:AdminAreaCountyId" yaml:",omitempty"`
	Feature            *Feature         `gorm:"foreignKey:FeatureId" json:",omitempty" yaml:",omitempty"`
	HighwayTaxomomy    []string         `gorm:"-" yaml:"highway"`
	AdminAreaCountryId uint             `gorm:"column:admin_area_country_id" json:"-" yaml:"-"`
	AdminAreaStateId   uint             `gorm:"column:admin_area_state_id"  yaml:"-"`
	AdminAreaCountyId  *uint            `gorm:"column:admin_area_county_id"  yaml:"-"`
	AdminAreaPlaceId   *uint            `gorm:"column:admin_area_place_id" yaml:"-"`
	IsArchived         bool             `gorm:"column:archived" `
	Tags               []HighwaySignTag `gorm:"foreignKey:SignId" yaml:"-"`
	DateAdded          time.Time        `gorm:"column:date_added"`
	LastUpdated        time.Time        `gorm:"column:last_update"`
}

func (HighwaySign) TableName() string {
	return "highwaysign"
}

func (s HighwaySign) X() float64 {
	return s.Point.X
}
func (s HighwaySign) Y() float64 {
	return s.Point.Y
}
func (s HighwaySign) Id() string {
	return s.ImageId.String()
}

func (s HighwaySign) ConvertToDto() generator.Generator {
	gh := geohash.EncodeWithPrecision(s.Y(), s.X(), 5)
	highwaySignDto := dto.HighwaySignDto{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		FeatureId:   s.FeatureId,
		DateTaken:   s.DateTaken,
		ImageId:     s.ImageId.String(),
		FlickrId:    s.FlickrId,
		Point: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  s.Y(),
			Longitude: s.X(),
		},
		StateSlug:   s.State.Slug,
		CountrySlug: s.Country.Slug,
		Recent:      s.DateTaken.Format("2006-01"),
		GeoHash:     gh,
		Tags: util.SliceMap(s.Tags, func(t HighwaySignTag) string {
			return t.Tag.Name
		}),
	}

	if s.Place != nil {
		placeFix := s.Place
		placeSlug := fmt.Sprintf("%s_%s", s.State.Slug, placeFix.Slug)
		highwaySignDto.PlaceSlug = &placeSlug
	}

	if s.StateSubdivision != nil {
		countyFix := s.StateSubdivision
		countySlug := fmt.Sprintf("%s_%s", s.State.Slug, countyFix.Slug)
		highwaySignDto.StateSubdivisionSlug = &countySlug
	}

	hwyTaxonomy := make([]string, len(s.Highways))
	var toHighways []string
	for i, _ := range s.Highways {
		hwyTaxonomy[i] = s.Highways[i].Highway.Slug

		if s.Highways[i].IsTo {
			toHighways = append(toHighways, s.Highways[i].Highway.Slug)
		}
	}
	highwaySignDto.Highways = hwyTaxonomy
	highwaySignDto.ToHighways = toHighways

	highwaySignDto.ImageWidth = s.ImageWidth
	highwaySignDto.ImageHeight = s.ImageHeight

	return highwaySignDto
}
