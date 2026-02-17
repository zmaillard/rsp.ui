package dto

import (
	"fmt"
	"highway-sign-portal-builder/pkg/generator"
	"time"

	"gopkg.in/yaml.v3"
)

type Converter interface {
	ConvertToDto() generator.Generator
}

type HighwaySignDto struct {
	ID          uint
	Title       string
	Description string `yaml:"-"`
	FeatureId   uint
	DateTaken   time.Time `yaml:"date"`
	ImageId     string
	FlickrId    *string `yaml:",omitempty"`
	PlaceSlug   *string `yaml:"place"`
	Point       struct {
		Latitude  float64
		Longitude float64
	} `yaml:"point,omitempty"`
	StateSlug            string   `yaml:"state"`
	CountrySlug          string   `yaml:"country"`
	Highways             []string `yaml:"highway"`
	ToHighways           []string `yaml:",omitempty"`
	StateSubdivisionSlug *string  `yaml:"county,omitempty"`
	Recent               string
	GeoHash              string
	ImageWidth           int      `yaml:"imageWidth"`
	ImageHeight          int      `yaml:"imageHeight"`
	Tags                 []string `yaml:"tags"`
	Categories           []string `yaml:"categories,omitempty"`
	Quality              int      `yaml:"quality"`
	PlusCode             string   `yaml:"plusCode"`
	LQIP                 *string  `yaml:"lqip,omitempty"`
	HasProcessed         bool     `yaml:"hasProcessed"`
}

func (s HighwaySignDto) OutFile() string {
	return fmt.Sprintf("content/sign/%s.md", s.ImageId)
}

func (s HighwaySignDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return nil, err
	}

	formattedYaml := AddYamlFrontAndEndMatterText(y, s.Description)

	return formattedYaml, nil

}
