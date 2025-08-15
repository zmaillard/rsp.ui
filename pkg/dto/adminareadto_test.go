package dto_test

import (
	"bytes"
	"highway-sign-portal-builder/pkg/dto"
	"path"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestAdminAreaCountryDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.AdminAreaCountryDto
		wantErr bool
	}{
		{
			name: "complete country data",
			input: dto.AdminAreaCountryDto{
				Name:            "United States",
				Slug:            "us",
				SubdivisionName: "State",
				ImageCount:      5000,
				States: []dto.AdminAreaSlimDto{
					{Name: "Washington", Slug: "washington"},
					{Name: "Oregon", Slug: "oregon"},
				},
				HighwayTypes: []dto.AdminAreaSlimDto{
					{Name: "Interstate", Slug: "interstate"},
					{Name: "US Highway", Slug: "us-highway"},
				},
				Featured: "i90/eastbound/sign.jpg",
			},
			wantErr: false,
		},
		{
			name: "minimal country data",
			input: dto.AdminAreaCountryDto{
				Name:       "Canada",
				Slug:       "ca",
				ImageCount: 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToMarkdown()

			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			expected, err := yaml.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal test input: %v", err)
			}

			expected = dto.AddYamlFrontAndEndMatter(expected)

			if !bytes.Equal(result, expected) {
				t.Errorf("ToMarkdown() = %s, want %s", result, expected)
			}
		})
	}
}

func TestAdminAreaCountryDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		slug     string
		expected string
	}{
		{
			name:     "simple slug",
			slug:     "us",
			expected: "content/country/us/_index.md",
		},
		{
			name:     "hyphenated slug",
			slug:     "united-kingdom",
			expected: "content/country/united-kingdom/_index.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := dto.AdminAreaCountryDto{
				Slug: tt.slug,
			}

			result := country.OutFile()

			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			dir, file := path.Split(result)
			expectedDir := "content/country/" + tt.slug + "/"
			if dir != expectedDir {
				t.Errorf("OutFile() directory = %s, want %s", dir, expectedDir)
			}

			if file != "_index.md" {
				t.Errorf("OutFile() filename = %s, want _index.md", file)
			}
		})
	}
}

func TestAdminAreaStateDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.AdminAreaStateDto
		wantErr bool
	}{
		{
			name: "complete state data",
			input: dto.AdminAreaStateDto{
				Name:            "Washington",
				Slug:            "washington",
				SubdivisionName: "State",
				ImageCount:      1200,
				Highways:        []string{"i90", "i5", "us101"},
				CountrySlug:     "us",
				Featured:        "i90/washington/sign.jpg",
				StateCategories: []string{"pacific", "northwest"},
				Counties: []dto.AdminAreaSlimDto{
					{Name: "King", Slug: "king"},
					{Name: "Pierce", Slug: "pierce"},
				},
				Places: []dto.AdminAreaSlimDto{
					{Name: "Seattle", Slug: "seattle"},
					{Name: "Tacoma", Slug: "tacoma"},
				},
				HighwayNames: []string{"Interstate 90", "Interstate 5"},
				Output:       []string{"html", "json"},
			},
			wantErr: false,
		},
		{
			name: "minimal state data",
			input: dto.AdminAreaStateDto{
				Name:        "Oregon",
				Slug:        "oregon",
				ImageCount:  500,
				CountrySlug: "us",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToMarkdown()

			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// Verify layout is set
			if !bytes.Contains(result, []byte("layout: state")) {
				t.Errorf("ToMarkdown() did not set layout to state")
			}

			// We need to set the layout for the expected output as well
			tt.input.Layout = "state"

			expected, err := yaml.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal test input: %v", err)
			}

			expected = dto.AddYamlFrontAndEndMatter(expected)

			if !bytes.Equal(result, expected) {
				t.Errorf("ToMarkdown() = %s, want %s", result, expected)
			}
		})
	}
}

func TestAdminAreaStateDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		slug     string
		expected string
	}{
		{
			name:     "simple slug",
			slug:     "washington",
			expected: "content/state/washington/_index.md",
		},
		{
			name:     "hyphenated slug",
			slug:     "new-york",
			expected: "content/state/new-york/_index.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := dto.AdminAreaStateDto{
				Slug: tt.slug,
			}

			result := state.OutFile()

			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			dir, file := path.Split(result)
			expectedDir := "content/state/" + tt.slug + "/"
			if dir != expectedDir {
				t.Errorf("OutFile() directory = %s, want %s", dir, expectedDir)
			}

			if file != "_index.md" {
				t.Errorf("OutFile() filename = %s, want _index.md", file)
			}
		})
	}
}

func TestAdminAreaCountyDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.AdminAreaCountyDto
		wantErr bool
	}{
		{
			name: "complete county data",
			input: dto.AdminAreaCountyDto{
				Name:       "King County",
				Slug:       "king",
				ImageCount: 300,
				StateSlug:  "washington",
				Aliases:    []string{"King", "King Co"},
			},
			wantErr: false,
		},
		{
			name: "minimal county data",
			input: dto.AdminAreaCountyDto{
				Name:      "Pierce County",
				Slug:      "pierce",
				StateSlug: "washington",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToMarkdown()

			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			expected, err := yaml.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal test input: %v", err)
			}

			expected = dto.AddYamlFrontAndEndMatter(expected)

			if !bytes.Equal(result, expected) {
				t.Errorf("ToMarkdown() = %s, want %s", result, expected)
			}
		})
	}
}

func TestAdminAreaCountyDto_OutFile(t *testing.T) {
	tests := []struct {
		name       string
		stateSlug  string
		countySlug string
		expected   string
	}{
		{
			name:       "simple slugs",
			stateSlug:  "washington",
			countySlug: "king",
			expected:   "content/county/washington_king/_index.md",
		},
		{
			name:       "hyphenated slugs",
			stateSlug:  "new-york",
			countySlug: "new-york",
			expected:   "content/county/new-york_new-york/_index.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			county := dto.AdminAreaCountyDto{
				StateSlug: tt.stateSlug,
				Slug:      tt.countySlug,
			}

			result := county.OutFile()

			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			dir, file := path.Split(result)
			expectedDir := "content/county/" + tt.stateSlug + "_" + tt.countySlug + "/"
			if dir != expectedDir {
				t.Errorf("OutFile() directory = %s, want %s", dir, expectedDir)
			}

			if file != "_index.md" {
				t.Errorf("OutFile() filename = %s, want _index.md", file)
			}
		})
	}
}

func TestAdminAreaPlaceDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.AdminAreaPlaceDto
		wantErr bool
	}{
		{
			name: "complete place data",
			input: dto.AdminAreaPlaceDto{
				Name:       "Seattle",
				Slug:       "seattle",
				ImageCount: 200,
				StateSlug:  "washington",
			},
			wantErr: false,
		},
		{
			name: "minimal place data",
			input: dto.AdminAreaPlaceDto{
				Name:      "Tacoma",
				Slug:      "tacoma",
				StateSlug: "washington",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToMarkdown()

			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			expected, err := yaml.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal test input: %v", err)
			}

			expected = dto.AddYamlFrontAndEndMatter(expected)

			if !bytes.Equal(result, expected) {
				t.Errorf("ToMarkdown() = %s, want %s", result, expected)
			}
		})
	}
}

func TestAdminAreaPlaceDto_OutFile(t *testing.T) {
	tests := []struct {
		name      string
		stateSlug string
		placeSlug string
		expected  string
	}{
		{
			name:      "simple slugs",
			stateSlug: "washington",
			placeSlug: "seattle",
			expected:  "content/place/washington_seattle/_index.md",
		},
		{
			name:      "hyphenated slugs",
			stateSlug: "new-york",
			placeSlug: "new-york-city",
			expected:  "content/place/new-york_new-york-city/_index.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			place := dto.AdminAreaPlaceDto{
				StateSlug: tt.stateSlug,
				Slug:      tt.placeSlug,
			}

			result := place.OutFile()

			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			dir, file := path.Split(result)
			expectedDir := "content/place/" + tt.stateSlug + "_" + tt.placeSlug + "/"
			if dir != expectedDir {
				t.Errorf("OutFile() directory = %s, want %s", dir, expectedDir)
			}

			if file != "_index.md" {
				t.Errorf("OutFile() filename = %s, want _index.md", file)
			}
		})
	}
}
