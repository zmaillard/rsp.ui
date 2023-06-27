package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"github.com/mmcloughlin/geohash"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"time"
)

type HighwaySigns []HighwaySign

func (HighwaySigns) OutLookupFiles() []string {
	return []string{"netlify/edge-functions/common/images.json", "data/images.json"}
}

func (hs HighwaySigns) GetLookup() ([]byte, error) {
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{
			"images":     []string{},
			"mostRecent": "",
			"imageCount": 0,
		})

	}
	var images []string

	dateTaken := hs[0].DateTaken

	for _, v := range hs {
		if dateTaken.Before(v.DateTaken) {
			dateTaken = v.DateTaken
		}
		images = append(images, v.ImageId.String())

	}
	res := map[string]interface{}{
		"images":     images,
		"mostRecent": dateTaken,
		"imageCount": len(images),
	}

	return json.Marshal(res)
}

type HighwaySign struct {
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

func (HighwaySign) TableName() string {
	return "vwhugohighwaysign"
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
