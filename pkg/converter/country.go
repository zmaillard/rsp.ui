package converter

import (
	"context"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"highway-sign-portal-builder/pkg/util"
	"iter"
)

type countryConverter struct {
	countries *[]db.SignVwhugocountry
}

func NewCountryConverter(ctx context.Context, queries db.Querier) (Converter, error) {
	countries, err := queries.GetHugoCountries(ctx)
	if err != nil {
		return nil, err
	}
	return &countryConverter{&countries}, nil
}
func (c countryConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, country := range *c.countries {
			countryDto := dto.AdminAreaCountryDto{
				Name:            country.CountryName.String,
				Slug:            country.CountrySlug.String,
				SubdivisionName: country.SubdivisionName.String,
				ImageCount:      int(country.ImageCount),
				Featured:        country.Featured.String(),
				States: util.SliceMap(country.States, func(state types.AdminArea) dto.AdminAreaSlimDto {
					return dto.AdminAreaSlimDto{
						Name: state.Name,
						Slug: state.Slug,
					}
				}),
				HighwayTypes: util.SliceMap(country.States, func(ht types.AdminArea) dto.AdminAreaSlimDto {
					return dto.AdminAreaSlimDto{
						Name: ht.Name,
						Slug: ht.Slug,
					}
				}),
			}
			yield(countryDto)
		}
	}
}
