package converter

import (
	"context"
	"fmt"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"iter"
)

type highwayConverter struct {
	highways *[]db.GetHugoHighwaysRow
}

func NewHighwayConverter(ctx context.Context, db db.Querier) (Converter, error) {
	hwys, err := db.GetHugoHighways(ctx)
	if err != nil {
		return nil, err
	}
	return &highwayConverter{&hwys}, nil
}

func (h highwayConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, hw := range *h.highways {
			highwayDto := dto.HighwayDto{
				Name:        hw.HighwayName.String,
				DisplayName: hw.DisplayName.String,
				Slug:        hw.Slug.String,
				Image:       hw.ImageName.String,
				Sort:        int(hw.SortNumber.Int32),
				HighwayTypeSlug: dto.AdminAreaSlimDto{
					Name: hw.HighwayTypeName.String,
					Slug: hw.HighwayTypeSlug.String,
				},
				Features: getFromTo(hw.PreviousFeatures, hw.NextFeatures),
				Places:   hw.Places,
				States:   hw.States,
				Counties: hw.Counties,
			}

			var aliases []string
			for _, v := range hw.States {
				aliasUrl := fmt.Sprintf("/statehighway/%s/%s", v, hw.HighwayName.String)
				aliases = append(aliases, aliasUrl)
			}

			highwayDto.Aliases = aliases

			yield(highwayDto)
		}
	}
}
