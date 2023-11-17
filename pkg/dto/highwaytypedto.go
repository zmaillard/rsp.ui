package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type HighwayTypeDto struct {
	Name            string   `yaml:"name"`
	Slug            string   `yaml:"slug"`
	Sort            int      `yaml:"sort"`
	ImageCount      int      `yaml:"imageCount"`
	Featured        string   `yaml:"featured,omitempty"`
	HighwayTaxomomy []string `yaml:"highways"`
	Country         string   `yaml:"country"`
}

func (ht HighwayTypeDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(ht)
	if err != nil {
		return nil, err
	}

	formattedYaml := addYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (ht HighwayTypeDto) OutFile() string {
	return fmt.Sprintf("content/highwaytype/%s.md", ht.Slug)
}
