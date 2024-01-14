package models

import (
	"github.com/lib/pq"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
)

type Highway struct {
	ID               uint           `gorm:"column:id,primaryKey"`
	Name             string         `gorm:"column:highway_name"`
	Slug             string         `gorm:"column:slug"`
	Sort             int            `gorm:"column:sort_number"`
	Image            string         `gorm:"column:image_name"`
	HighwayTypeSlug  string         `gorm:"column:highway_type_slug"`
	HighwayTypeName  string         `gorm:"column:highway_type_name"`
	States           pq.StringArray `gorm:"column:states;type:text[]"`
	Counties         pq.StringArray `gorm:"column:counties;type:text[]"`
	Places           pq.StringArray `gorm:"column:places;type:text[]"`
	PreviousFeatures pq.Int32Array  `gorm:"column:previous_features;type:int[]"`
	NextFeatures     pq.Int32Array  `gorm:"column:next_features;type:int[]"`
}

func (Highway) TableName() string {
	return "sign.vwhugohighway"
}

func (h Highway) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayDto{
		Name:  h.Name,
		Slug:  h.Slug,
		Image: h.Image,
		Sort:  h.Sort,
		HighwayTypeSlug: dto.AdminAreaSlimDto{
			Name: h.HighwayTypeName,
			Slug: h.HighwayTypeSlug,
		},
		Features: getFromTo(h.PreviousFeatures, h.NextFeatures),
		Places:   h.Places,
		States:   h.States,
		Counties: h.Counties,
	}

	var aliases []string
	for _, v := range h.States {
		aliases = append(aliases, "/statehighway/"+v+"/"+h.Name)
	}

	highwayDto.Aliases = aliases

	return highwayDto
}
