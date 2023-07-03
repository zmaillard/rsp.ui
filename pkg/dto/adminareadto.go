package dto

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type AdminAreaSlimDto struct {
	Name string
	Slug string
}

type AdminAreaCountryDto struct {
	Name            string
	Slug            string
	SubdivisionName string `yaml:",omitempty"`
	ImageCount      int
	States          []AdminAreaSlimDto
	HighwayTypes    []AdminAreaSlimDto
	Featured        string `yaml:"featured,omitempty"`
}

func (c AdminAreaCountryDto) ToMarkdown() ([]byte, error) {

	y, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}

	formattedYaml := addYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (c AdminAreaCountryDto) OutFile() string {
	return fmt.Sprintf("content/country/%s/_index.md", c.Slug)
}

type AdminAreaStateDto struct {
	Name            string
	Slug            string
	SubdivisionName string
	ImageCount      int
	Highways        []string
	Layout          string
	CountrySlug     string
	Featured        string `yaml:"featured,omitempty"`
	Counties        []AdminAreaSlimDto
	Places          []AdminAreaSlimDto
}

func (s AdminAreaStateDto) ToMarkdown() ([]byte, error) {
	s.Layout = "state"
	y, err := yaml.Marshal(s)
	if err != nil {
		return nil, err
	}

	formattedYaml := addYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (s AdminAreaStateDto) OutFile() string {
	return fmt.Sprintf("content/state/%s/_index.md", s.Slug)
}

type AdminAreaCountyDto struct {
	Name       string
	Slug       string
	ImageCount int
	StateSlug  string
	Aliases    []string
}

func (s AdminAreaCountyDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return nil, err
	}

	formattedYaml := addYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (s AdminAreaCountyDto) OutFile() string {
	return fmt.Sprintf("content/county/%s_%s/_index.md", s.StateSlug, s.Slug)
}

type AdminAreaPlaceDto struct {
	Name       string
	Slug       string
	ImageCount int
	StateSlug  string
}

func (s AdminAreaPlaceDto) ToMarkdown() ([]byte, error) {
	y, err := yaml.Marshal(s)
	if err != nil {
		return nil, err
	}

	formattedYaml := addYamlFrontAndEndMatter(y)

	return formattedYaml, nil
}

func (s AdminAreaPlaceDto) OutFile() string {
	return fmt.Sprintf("content/place/%s_%s/_index.md", s.StateSlug, s.Slug)
}
