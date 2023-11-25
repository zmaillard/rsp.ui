package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"github.com/mmcloughlin/geohash"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
	"highway-sign-portal-builder/pkg/types"
	"time"
)

type HighwaySign struct {
	ID               uint           `gorm:"column:id;primaryKey"`
	Title            string         `gorm:"column:title"`
	Description      string         `gorm:"column:sign_description"`
	FeatureId        uint           `gorm:"column:feature_id"`
	DateTaken        time.Time      `gorm:"column:date_taken"`
	ImageId          types.ImageID  `gorm:"column:imageid"`
	FlickrId         *string        `gorm:"column:flickrid"`
	Point            Point          `gorm:"column:point"`
	Country          string         `gorm:"column:country_slug"`
	State            string         `gorm:"column:state_slug"`
	Place            *string        `gorm:"column:place_slug"`
	StateSubdivision *string        `gorm:"column:county_slug"`
	Tags             pq.StringArray `gorm:"column:tags;type:text[]"`
	Highways         pq.StringArray `gorm:"column:highways;type:text[]"`
	IsTo             pq.StringArray `gorm:"column:is_to;type:text[]"`
	ImageHeight      int            `gorm:"column:image_height"`
	ImageWidth       int            `gorm:"column:image_width"`
}

func (HighwaySign) TableName() string {
	return "vwhugohighwaysign"
}

func (s HighwaySign) X() float64 {
	return s.Point.X
}
func (s HighwaySign) Y() float64 {
	return s.Point.Y
}
func (s HighwaySign) Id() string {
	return s.ImageId.String()
}

func (s HighwaySign) ConvertToDto() generator.Generator {
	gh := geohash.EncodeWithPrecision(s.Y(), s.X(), 5)
	highwaySignDto := dto.HighwaySignDto{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
		FeatureId:   s.FeatureId,
		DateTaken:   s.DateTaken,
		ImageId:     s.ImageId.String(),
		FlickrId:    s.FlickrId,
		Point: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  s.Y(),
			Longitude: s.X(),
		},
		StateSlug:            s.State,
		CountrySlug:          s.Country,
		Recent:               s.DateTaken.Format("2006-01"),
		GeoHash:              gh,
		Tags:                 s.Tags,
		PlaceSlug:            s.Place,
		StateSubdivisionSlug: s.StateSubdivision,
		Highways:             s.Highways,
		ToHighways:           s.IsTo,
		ImageWidth:           s.ImageWidth,
		ImageHeight:          s.ImageHeight,
	}

	return highwaySignDto
}

type HighwaySigns []HighwaySign

func (hs HighwaySigns) GetStateLookup() generator.Lookup {
	return &highwaySignsByState{
		&hs,
	}
}

func (hs HighwaySigns) GetPlaceLookup() generator.Lookup {
	return &highwaySignsByPlace{
		&hs,
	}
}

func (hs HighwaySigns) GetCountyLookup() generator.Lookup {
	return &highwaySignsByCounty{
		&hs,
	}
}

func (hs HighwaySigns) GetGeoJsonLookup() generator.Lookup {
	return &highwaySignsInGeoJson{
		&hs,
	}
}

func (HighwaySigns) OutLookupFiles() []string {
	return []string{"data/images.json"}
}

func (hs HighwaySigns) GetLookup() ([]byte, error) {
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{
			"images":     []string{},
			"mostRecent": "",
			"imageCount": 0,
		})

	}
	var images []string

	dateTaken := hs[0].DateTaken

	for _, v := range hs {
		if dateTaken.Before(v.DateTaken) {
			dateTaken = v.DateTaken
		}
		images = append(images, v.ImageId.String())

	}
	res := map[string]interface{}{
		"images":     images,
		"mostRecent": dateTaken,
		"imageCount": len(images),
	}

	return json.Marshal(res)
}

type highwaySignsByState struct {
	*HighwaySigns
}

func (h highwaySignsByState) GetLookup() ([]byte, error) {
	hs := *h.HighwaySigns
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		_, ok := images[v.State]
		if !ok {
			images[v.State] = make([]string, 0)
		}
		images[v.State] = append(images[v.State], v.ImageId.String())

	}

	return json.Marshal(images)
}

func (h highwaySignsByState) OutLookupFiles() []string {
	return []string{"data/state.json"}
}

type highwaySignsByPlace struct {
	*HighwaySigns
}

func (h highwaySignsByPlace) GetLookup() ([]byte, error) {
	hs := *h.HighwaySigns
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		if v.Place == nil {
			continue
		}

		p := *v.Place
		_, ok := images[p]
		if !ok {
			images[p] = make([]string, 0)
		}
		images[p] = append(images[p], v.ImageId.String())
	}

	return json.Marshal(images)
}

func (h highwaySignsByPlace) OutLookupFiles() []string {
	return []string{"data/place.json"}
}

type highwaySignsByCounty struct {
	*HighwaySigns
}

func (h highwaySignsByCounty) GetLookup() ([]byte, error) {
	hs := *h.HighwaySigns
	// Make sure there are actually records
	if len(hs) == 0 {
		return json.Marshal(map[string]interface{}{})
	}
	images := make(map[string][]string)

	for _, v := range hs {
		if v.StateSubdivision == nil {
			continue
		}

		county := *v.StateSubdivision
		_, ok := images[county]
		if !ok {
			images[county] = make([]string, 0)
		}
		images[county] = append(images[county], v.ImageId.String())
	}

	return json.Marshal(images)
}

func (h highwaySignsByCounty) OutLookupFiles() []string {
	return []string{"data/county.json"}

}

type highwaySignsInGeoJson struct {
	*HighwaySigns
}

func (h highwaySignsInGeoJson) GetLookup() ([]byte, error) {
	hs := *h.HighwaySigns
	var jsonItems []geoJsonFeature

	for _, h := range hs {
		geoJsonFeature := geoJsonFeature{
			Type: "Feature",
			Geometry: struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			}{
				Type:        "Point",
				Coordinates: []float64{h.Point.X, h.Point.Y},
			},
			Properties: struct {
				ImageId string `json:"imageid"`
				Title   string `json:"title"`
			}{
				ImageId: h.ImageId.String(),
				Title:   h.Title,
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
