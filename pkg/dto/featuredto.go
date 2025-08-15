package dto

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type FeatureDto struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Next         []FeatureLinkDto `yaml:",omitempty"`
	Prev         []FeatureLinkDto `yaml:",omitempty"`
	Signs        []string         `yaml:",omitempty"`
	HighwayNames []string         `yaml:"highwayNames,omitempty"`
	State        struct {
		Name string
		Slug string
	} `yaml:"state,omitempty"`
	Country struct {
		Name string
		Slug string
	} `yaml:"country,omitempty"`
}

func (f FeatureDto) OutFile() string {
	return fmt.Sprintf("content/feature/%v.md", f.ID)
}

func (f FeatureDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(f)
	if err != nil {
		return nil, err
	}

	formattedYaml := AddYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}
