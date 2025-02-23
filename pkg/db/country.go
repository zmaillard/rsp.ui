package db

import (
	"fmt"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/models"
	"highway-sign-portal-builder/pkg/util"
)

func (c SignVwhugocountry) ConvertToDto() generator.Generator {

	return dto.AdminAreaCountryDto{
		Name:            c.CountryName.String,
		Slug:            c.CountrySlug.String,
		SubdivisionName: c.SubdivisionName.String,
		ImageCount:      int(c.ImageCount),
		Featured:        c.Featured.String(),
		States: util.SliceMap(c.States, func(state models.AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: state.Name,
				Slug: state.Slug,
			}
		}),
		HighwayTypes: util.SliceMap(c.HighwayTypes, func(ht models.AdminArea) dto.AdminAreaSlimDto {
			return dto.AdminAreaSlimDto{
				Name: ht.Name,
				Slug: ht.Slug,
			}
		}),
	}
}

func (c GetHugoCountiesRow) ConvertToDto() generator.Generator {
	countyDto := dto.AdminAreaCountyDto{
		Name:       c.CountyName.String,
		Slug:       c.CountySlug.String,
		ImageCount: int(c.ImageCount),
		StateSlug:  c.StateSlug.String,
		Aliases:    []string{fmt.Sprintf("/county/%s/%s", c.StateSlug.String, c.CountySlug.String)},
	}

	return countyDto
}

func foo(featurelinks []GetHugoFeatureLinksRow, features []SignVwhugofeature) {
	featureMap := make(map[int32]SignVwhugofeature)
	for _, feature := range features {
		featureMap[feature.ID] = feature
	}

}
