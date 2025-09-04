package converter

import (
	"context"
	"fmt"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"highway-sign-portal-builder/pkg/util"
	"iter"
)

type stateConverter struct {
	states *[]db.GetHugoStatesRow
}

func NewStateConverter(ctx context.Context, queries db.Querier) (Converter, error) {
	states, err := queries.GetHugoStates(ctx)
	if err != nil {
		return nil, err
	}
	return &stateConverter{&states}, nil
}

func (c stateConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, state := range *c.states {

			// Filter blank places
			validPlaces := util.SliceFilter(state.Places, func(highway types.AdminArea) bool {
				return highway.Name != "" && highway.Slug != ""
			})

			var featured string
			if state.Featured != nil {
				featured = state.Featured.String()
			}
			stateDto := dto.AdminAreaStateDto{
				Name:            state.StateName.String,
				Slug:            state.StateSlug.String,
				SubdivisionName: state.SubdivisionName.String,
				ImageCount:      int(state.ImageCount),
				Layout:          "state",
				Output:          []string{"html", "list"},
				Highways:        state.Highways,
				CountrySlug:     state.CountrySlug.String,
				Featured:        featured,
				HighwayNames:    state.HighwayNames,
				Counties: util.SliceMap(state.Counties, func(county types.AdminArea) dto.AdminAreaSlimDto {
					return dto.AdminAreaSlimDto{
						Name: county.Name,
						Slug: county.Slug,
					}
				}),
				Places: util.SliceMap(validPlaces, func(place types.AdminArea) dto.AdminAreaSlimDto {
					return dto.AdminAreaSlimDto{
						Name: place.Name,
						Slug: place.Slug,
					}
				}),
			}

			var categories []string
			for _, v := range state.Categories {
				stateCat := fmt.Sprintf("%s_%s", state.StateSlug.String, v)
				categories = append(categories, stateCat)
			}

			stateDto.StateCategories = categories

			yield(stateDto)
		}
	}
}
