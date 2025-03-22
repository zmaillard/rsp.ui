package converter

import (
	"context"
	"github.com/twpayne/go-geom"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/util"
	"iter"
	"math"
)

type featureConverter struct {
	features     *[]db.GetHugoFeaturesRow
	featureLinks *[]db.SignVwhugofeaturelink
}

func NewFeatureConverter(ctx context.Context, db db.Querier) (Converter, error) {
	features, err := db.GetHugoFeatures(ctx)
	if err != nil {
		return nil, err
	}
	featureLinks, err := db.GetHugoFeatureLinks(ctx)
	if err != nil {
		return nil, err
	}
	return &featureConverter{&features, &featureLinks}, nil
}

func (f featureConverter) Convert() iter.Seq[generator.Generator] {
	toLinks := make(map[int][]db.SignVwhugofeaturelink)
	fromLinks := make(map[int][]db.SignVwhugofeaturelink)

	for _, link := range *f.featureLinks {
		fromFeatureId := int(link.FromFeature.Int32)
		toFeatureId := int(link.ToFeature.Int32)
		toLinks[fromFeatureId] = append(toLinks[fromFeatureId], link)
		fromLinks[toFeatureId] = append(fromLinks[toFeatureId], link)
	}

	return func(yield func(generator.Generator) bool) {
		for _, feature := range *f.features {
			featureDto := dto.FeatureDto{
				ID:   uint(feature.ID),
				Name: feature.Name.String,
				Next: util.SliceMap(toLinks[int(feature.ID)], func(link db.SignVwhugofeaturelink) dto.FeatureLinkDto {
					azimuth := azimuth(feature.Point, link.ToPoint)
					return dto.FeatureLinkDto{
						ID:            uint(link.ID),
						RoadName:      link.RoadName.String,
						FromFeatureId: uint(link.FromFeature.Int32),
						ToFeatureId:   uint(link.FromFeature.Int32),
						Highways:      link.Highways,
						Direction:     azimuth,
						Bearing:       bearing(azimuth),
					}
				}),
				Prev: util.SliceMap(fromLinks[int(feature.ID)], func(link db.SignVwhugofeaturelink) dto.FeatureLinkDto {
					azimuth := azimuth(feature.Point, link.FromPoint)
					return dto.FeatureLinkDto{
						ID:            uint(link.ID),
						RoadName:      link.RoadName.String,
						FromFeatureId: uint(link.FromFeature.Int32),
						ToFeatureId:   uint(link.ToFeature.Int32),
						Highways:      link.Highways,
						Direction:     azimuth,
						Bearing:       bearing(azimuth),
					}
				}),
				Signs: feature.Signs,
				State: struct {
					Name string
					Slug string
				}{
					Name: feature.StateName.String,
					Slug: feature.StateSlug.String,
				},
				Country: struct {
					Name string
					Slug string
				}{
					Name: feature.CountryName.String,
					Slug: feature.CountrySlug.String,
				},
			}

			yield(featureDto)
		}
	}
}
func azimuth(a geom.Point, b geom.Point) float64 {
	bearing := math.Atan2(b.Y()-a.Y(), b.X()-a.X()) * (180 / math.Pi)

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
