package converter

import (
	"context"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"iter"
)

type highwayTypeConverter struct {
	highwayTypes *[]db.GetHugoHighwayTypesRow
}

func NewHighwayTypeConverter(ctx context.Context, db db.Querier) (Converter, error) {
	hwyTypes, err := db.GetHugoHighwayTypes(ctx)
	if err != nil {
		return nil, err
	}
	return &highwayTypeConverter{&hwyTypes}, nil
}

func (h highwayTypeConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, ht := range *h.highwayTypes {
			highwayTypeDto := dto.HighwayTypeDto{
				Name:            ht.HighwayTypeName.String,
				Slug:            ht.HighwayTypeSlug.String,
				Sort:            int(ht.Sort.Int32),
				ImageCount:      int(ht.Imagecount),
				HighwayTaxomomy: ht.Highways,
				Country:         ht.Country.String,
			}
			if ht.Imageid != nil {
				highwayTypeDto.Featured = ht.Imageid.String()
			}

			yield(highwayTypeDto)
		}
	}
}
