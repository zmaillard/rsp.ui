package dto_test

import (
	"bytes"
	"highway-sign-portal-builder/pkg/dto"
	"path"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestHighwayTypeDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name        string
		input       dto.HighwayTypeDto
		wantErr     bool
		checkOutput bool
	}{
		{
			name: "basic highway type",
			input: dto.HighwayTypeDto{
				Name:            "Interstate",
				Slug:            "interstate",
				Sort:            1,
				ImageCount:      250,
				Featured:        "i90/eastbound/exit.jpg",
				HighwayTaxomomy: []string{"i5", "i90", "i15"},
				Country:         "us",
			},
			wantErr:     false,
			checkOutput: true,
		},
		{
			name: "without featured image",
			input: dto.HighwayTypeDto{
				Name:            "State Highway",
				Slug:            "state-highway",
				Sort:            2,
				ImageCount:      150,
				HighwayTaxomomy: []string{"wa-520", "wa-99"},
				Country:         "us",
			},
			wantErr:     false,
			checkOutput: true,
		},
		{
			name: "empty fields",
			input: dto.HighwayTypeDto{
				Name:    "Empty Type",
				Slug:    "empty",
				Country: "ca",
			},
			wantErr:     false,
			checkOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the method being tested
			result, err := tt.input.ToMarkdown()

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// For test cases where we want to verify the output
			if tt.checkOutput {
				// Marshal the input struct directly for comparison
				expected, err := yaml.Marshal(tt.input)
				if err != nil {
					t.Fatalf("Failed to marshal test input: %v", err)
				}

				// Apply the same formatting that the method would
				expected = dto.AddYamlFrontAndEndMatter(expected)

				// Compare
				if !bytes.Equal(result, expected) {
					t.Errorf("ToMarkdown() = %s, want %s", result, expected)
				}
			}
		})
	}
}

func TestHighwayTypeDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.HighwayTypeDto
		expected string
	}{
		{
			name: "simple slug",
			input: dto.HighwayTypeDto{
				Slug: "interstate",
			},
			expected: "content/highwaytype/interstate.md",
		},
		{
			name: "hyphenated slug",
			input: dto.HighwayTypeDto{
				Slug: "state-highway",
			},
			expected: "content/highwaytype/state-highway.md",
		},
		{
			name: "empty slug",
			input: dto.HighwayTypeDto{
				Slug: "",
			},
			expected: "content/highwaytype/.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.OutFile()

			// Check the path structure
			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			// Additional checks for path construction
			dir, file := path.Split(result)
			if dir != "content/highwaytype/" {
				t.Errorf("OutFile() directory = %s, want content/highwaytype/", dir)
			}

			expectedFile := tt.input.Slug + ".md"
			if file != expectedFile {
				t.Errorf("OutFile() filename = %s, want %s", file, expectedFile)
			}
		})
	}
}
