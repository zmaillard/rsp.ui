package converter

import (
	"context"
	"encoding/json"
	"fmt"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"iter"

	pluscode "github.com/google/open-location-code/go"
	"github.com/mmcloughlin/geohash"
)

type SignConverter struct {
	signs *[]db.SignVwhugohighwaysign
}

func NewHighwaySignConverter(ctx context.Context, db db.Querier) (Converter, error) {
	signs, err := db.GetHugoHighwaySigns(ctx)
	if err != nil {
		return nil, err
	}
	return &SignConverter{&signs}, nil
}

func (s SignConverter) Convert() iter.Seq[generator.Generator] {
	return func(yield func(generator.Generator) bool) {
		for _, sign := range *s.signs {
			gh := geohash.EncodeWithPrecision(sign.Point.Y(), sign.Point.X(), 5)

			pc := pluscode.Encode(sign.Point.Y(), sign.Point.X(), 10)
			highwaySignDto := dto.HighwaySignDto{
				ID:          uint(sign.ID),
				Title:       sign.Title.String,
				Description: sign.SignDescription.String,
				FeatureId:   uint(sign.FeatureID.Int32),
				DateTaken:   sign.DateTaken.Time,
				ImageId:     sign.Imageid.String(),
				Point: struct {
					Latitude  float64
					Longitude float64
				}{
					Latitude:  sign.Point.Y(),
					Longitude: sign.Point.X(),
				},
				StateSlug:            sign.StateSlug,
				CountrySlug:          sign.CountrySlug,
				PlaceSlug:            sign.PlaceSlug,
				StateSubdivisionSlug: sign.CountySlug,
				Recent:               sign.DateTaken.Time.Format("2006-01"),
				GeoHash:              gh,
				Highways:             sign.Highways,
				ToHighways:           sign.IsTo,
				ImageWidth:           int(sign.ImageWidth.Int32),
				ImageHeight:          int(sign.ImageHeight.Int32),
				Quality:              int(sign.Quality),
				Tags:                 sign.Tags,
				PlusCode:             pc,
				HasProcessed:         sign.HasProcessed.Bool,
			}

			if sign.LqipHash.Valid {
				highwaySignDto.LQIP = &sign.LqipHash.String
			}

			// Add state name to categories
			var cats []string
			for _, v := range sign.Categories {
				stateCat := fmt.Sprintf("%s_%s", sign.StateSlug, v)
				cats = append(cats, stateCat)
			}

			highwaySignDto.Categories = cats
			yield(highwaySignDto)
		}
	}

}

func (s SignConverter) GetHighQualityLookup() generator.Lookup {
	return &highwaySignsHighQuality{
		s.signs,
	}
}

func (s SignConverter) GetStateLookup() generator.Lookup {
	return &highwaySignsByState{
		s.signs,
	}
}

func (s SignConverter) GetPlaceLookup() generator.Lookup {
	return &highwaySignsByPlace{
		s.signs,
	}
}

func (s SignConverter) GetCountyLookup() generator.Lookup {
	return &highwaySignsByCounty{
		s.signs,
	}
}

func (s SignConverter) GetGeoJsonLookup() generator.Lookup {
	return &highwaySignsInGeoJson{
		s.signs,
	}
}

func (SignConverter) OutLookupFiles() []string {
	return []string{"data/images.json"}
}

func (s SignConverter) GetLookup() ([]byte, error) {
	// Make sure there are actually records
	hs := *s.signs
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{
			"images":     []string{},
			"mostRecent": "",
			"imageCount": 0,
		})

	}
	var images []string

	dateTaken := hs[0].DateTaken.Time

	for _, v := range hs {
		if dateTaken.Before(v.DateTaken.Time) {
			dateTaken = v.DateTaken.Time
		}
		images = append(images, v.Imageid.String())

	}
	res := map[string]interface{}{
		"images":     images,
		"mostRecent": dateTaken,
		"imageCount": len(images),
	}

	return json.Marshal(res)
}

type highwaySignsHighQuality struct {
	signs *[]db.SignVwhugohighwaysign
}

func (h highwaySignsHighQuality) GetLookup() ([]byte, error) {
	// Make sure there are actually records
	hs := *h.signs
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{
			"images":     []string{},
			"mostRecent": "",
			"imageCount": 0,
		})

	}
	var images []string

	dateTaken := hs[0].DateTaken.Time

	for _, v := range hs {
		if v.Quality > 3 {

			if dateTaken.Before(v.DateTaken.Time) {
				dateTaken = v.DateTaken.Time
			}
			images = append(images, v.Imageid.String())
		}
	}
	res := map[string]interface{}{
		"images":     images,
		"mostRecent": dateTaken,
		"imageCount": len(images),
	}

	return json.Marshal(res)
}

func (h highwaySignsHighQuality) OutLookupFiles() []string {
	return []string{"data/signsquality.json"}
}

type highwaySignsByState struct {
	signs *[]db.SignVwhugohighwaysign
}

func (h highwaySignsByState) GetLookup() ([]byte, error) {
	hs := *h.signs
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		_, ok := images[v.StateSlug]
		if !ok {
			images[v.StateSlug] = make([]string, 0)
		}
		images[v.StateSlug] = append(images[v.StateSlug], v.Imageid.String())

	}

	return json.Marshal(images)
}

func (h highwaySignsByState) OutLookupFiles() []string {
	return []string{"data/state.json"}
}

type highwaySignsByPlace struct {
	signs *[]db.SignVwhugohighwaysign
}

func (h highwaySignsByPlace) GetLookup() ([]byte, error) {
	hs := *h.signs
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		if v.PlaceSlug == nil {
			continue
		}

		p := *v.PlaceSlug
		_, ok := images[p]
		if !ok {
			images[p] = make([]string, 0)
		}
		images[p] = append(images[p], v.Imageid.String())
	}

	return json.Marshal(images)
}

func (h highwaySignsByPlace) OutLookupFiles() []string {
	return []string{"data/place.json"}
}

type highwaySignsByCounty struct {
	signs *[]db.SignVwhugohighwaysign
}

func (h highwaySignsByCounty) GetLookup() ([]byte, error) {
	hs := *h.signs
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		if v.CountySlug == nil {
			continue
		}

		county := *v.CountySlug
		_, ok := images[county]
		if !ok {
			images[county] = make([]string, 0)
		}
		images[county] = append(images[county], v.Imageid.String())
	}

	return json.Marshal(images)
}

func (h highwaySignsByCounty) OutLookupFiles() []string {
	return []string{"data/county.json"}

}

type highwaySignsInGeoJson struct {
	signs *[]db.SignVwhugohighwaysign
}

func (h highwaySignsInGeoJson) GetLookup() ([]byte, error) {
	hs := *h.signs
	var jsonItems []geoJsonFeature

	for _, h := range hs {
		geoJsonFeature := geoJsonFeature{
			Type: "Feature",
			Geometry: struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			}{
				Type:        "Point",
				Coordinates: []float64{h.Point.X(), h.Point.Y()},
			},
			Properties: struct {
				ImageId string `json:"imageid"`
				Title   string `json:"title"`
			}{
				ImageId: h.Imageid.String(),
				Title:   h.Title.String,
			},
		}

		jsonItems = append(jsonItems, geoJsonFeature)
	}

	geoFeatures := struct {
		Type     string           `json:"type"`
		Features []geoJsonFeature `json:"features"`
	}{
		Type:     "FeatureCollection",
		Features: jsonItems,
	}

	return json.Marshal(geoFeatures)

}

func (h highwaySignsInGeoJson) OutLookupFiles() []string {
	return []string{"tiles/signs.json"}
}

type geoJsonFeature struct {
	Type     string `json:"type"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		ImageId string `json:"imageid"`
		Title   string `json:"title"`
	} `json:"properties"`
}
