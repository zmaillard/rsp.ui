package converter

import (
	"context"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"iter"
)

type placeConverter struct {
	place *[]db.GetHugoPlacesRow
}

func NewPlaceConverter(ctx context.Context, queries db.Querier) (Converter, error) {
	place, err := queries.GetHugoPlaces(ctx)
	if err != nil {
		return nil, err
	}
	return &placeConverter{&place}, nil
}

func (c placeConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, place := range *c.place {

			placeDto := dto.AdminAreaPlaceDto{
				Name:       place.PlaceName.String,
				Slug:       place.PlaceSlug.String,
				ImageCount: int(place.ImageCount),
				StateSlug:  place.StateSlug.String,
			}

			yield(placeDto)
		}
	}
}
