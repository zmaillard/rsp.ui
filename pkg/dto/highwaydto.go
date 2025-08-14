package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type HighwayDto struct {
	Name            string           `yaml:"name"`
	DisplayName     string           `yaml:"displayName,omitempty"`
	Slug            string           `yaml:"slug"`
	Image           string           `yaml:"imageName"`
	HighwayTypeSlug AdminAreaSlimDto `yaml:"highwayType"`
	Sort            int              `yaml:"sort"`
	Features        []int32          `yaml:"features,omitempty"`
	Places          []string         `yaml:"places,omitempty"`
	States          []string         `yaml:"states,omitempty"`
	Counties        []string         `yaml:"counties,omitempty"`
	Aliases         []string         `yaml:"aliases,omitempty"`
}

func (h HighwayDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(h)
	if err != nil {
		return nil, err
	}

	formattedYaml := AddYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (h HighwayDto) OutFile() string {
	return fmt.Sprintf("content/highway/%s/_index.md", h.Slug)
}
