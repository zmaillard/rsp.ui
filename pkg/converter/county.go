package converter

import (
	"context"
	"fmt"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"iter"
)

type countyConverter struct {
	counties *[]db.GetHugoCountiesRow
}

func NewCountyConverter(ctx context.Context, queries db.Querier) (Converter, error) {
	counties, err := queries.GetHugoCounties(ctx)
	if err != nil {
		return nil, err
	}
	return &countyConverter{&counties}, nil
}

func (c countyConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, county := range *c.counties {

			countyDto := dto.AdminAreaCountyDto{
				Name:       county.CountyName.String,
				Slug:       county.CountySlug.String,
				ImageCount: int(county.ImageCount),
				StateSlug:  county.StateSlug.String,
				Aliases:    []string{fmt.Sprintf("/county/%s/%s", county.StateSlug.String, county.CountySlug.String)},
			}
			yield(countyDto)
		}
	}
}
