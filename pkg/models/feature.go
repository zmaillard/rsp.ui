package models

import (
	"github.com/lib/pq"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
	"math"
	"strconv"
)

type Feature struct {
	ID          uint           `gorm:"column:id;primaryKey"`
	Point       Point          `gorm:"column:point"`
	Name        string         `gorm:"column:name"`
	Next        []FeatureLink  `gorm:"foreignKey:FromFeatureId"`
	Prev        []FeatureLink  `gorm:"foreignKey:ToFeatureId"`
	Signs       pq.StringArray `gorm:"column:signs;type:text[]"`
	StateName   string         `gorm:"column:state_name"`
	StateSlug   string         `gorm:"column:state_slug"`
	CountryName string         `gorm:"column:country_name"`
	CountrySlug string         `gorm:"column:country_slug"`
}

func (Feature) TableName() string {
	return "sign.vwhugofeature"
}

func (f Feature) X() float64 {
	return f.Point.X
}
func (f Feature) Y() float64 {
	return f.Point.Y
}
func (f Feature) Id() string {
	return strconv.Itoa(int(f.ID))
}

func (f Feature) IsEmpty() bool {
	return f.ID == 0 && f.Name == "" && len(f.Next) == 0 && len(f.Prev) == 0 && len(f.Signs) == 0
}

func (f Feature) ConvertToDto() generator.Generator {
	feature := dto.FeatureDto{
		ID:   f.ID,
		Name: f.Name,
		Next: util.SliceMap(f.Next, func(link FeatureLink) dto.FeatureLinkDto {
			azimuth := azimuth(f.Point, link.ToPoint)
			return dto.FeatureLinkDto{
				ID:            link.ID,
				RoadName:      link.RoadName,
				FromFeatureId: link.FromFeatureId,
				ToFeatureId:   link.ToFeatureId,
				Highways:      link.Highways,
				Direction:     azimuth,
				Bearing:       bearing(azimuth),
			}
		}),
		Prev: util.SliceMap(f.Prev, func(link FeatureLink) dto.FeatureLinkDto {
			azimuth := azimuth(f.Point, link.FromPoint)
			return dto.FeatureLinkDto{
				ID:            link.ID,
				RoadName:      link.RoadName,
				FromFeatureId: link.FromFeatureId,
				ToFeatureId:   link.ToFeatureId,
				Highways:      link.Highways,
				Direction:     azimuth,
				Bearing:       bearing(azimuth),
			}
		}),
		Signs: f.Signs,
		State: struct {
			Name string
			Slug string
		}{
			Name: f.StateName,
			Slug: f.StateSlug,
		},
		Country: struct {
			Name string
			Slug string
		}{
			Name: f.CountryName,
			Slug: f.CountrySlug,
		},
	}

	return feature
}

func azimuth(a Point, b Point) float64 {
	bearing := math.Atan2(b.Y-a.Y, b.X-a.X) * (180 / math.Pi)

	if bearing < 0 {
		bearing += 360
	}

	return bearing
}

func bearing(angle float64) string {
	switch {
	case angle >= 337.5 || angle < 22.5:
		return "E"
	case angle >= 22.5 && angle < 67.5:
		return "NE"
	case angle >= 67.5 && angle < 112.5:
		return "N"
	case angle >= 112.5 && angle < 157.5:
		return "NW"
	case angle >= 157.5 && angle < 202.5:
		return "W"
	case angle >= 202.5 && angle < 247.5:
		return "SW"
	case angle >= 247.5 && angle < 292.5:
		return "S"
	case angle >= 292.5 && angle < 337.5:
		return "SE"
	}
	return ""
}
