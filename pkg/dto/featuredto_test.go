package dto_test

import (
	"bytes"
	"fmt"
	"highway-sign-portal-builder/pkg/dto"
	"path"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestFeatureDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.FeatureDto
		wantErr bool
	}{
		{
			name: "complete feature data",
			input: dto.FeatureDto{
				ID:   123,
				Name: "I-90 Exit 34 Interchange",
				Next: []dto.FeatureLinkDto{
					{ID: 124},
				},
				Prev: []dto.FeatureLinkDto{
					{ID: 122},
				},
				Signs:        []string{"i90-exit34-1.jpg", "i90-exit34-2.jpg"},
				HighwayNames: []string{"Interstate 90", "WA-202"},
			},
			wantErr: false,
		},
		{
			name: "with state and country",
			input: dto.FeatureDto{
				ID:   456,
				Name: "US-101 Junction",
				State: struct {
					Name string
					Slug string
				}{
					Name: "Washington",
					Slug: "washington",
				},
				Country: struct {
					Name string
					Slug string
				}{
					Name: "United States",
					Slug: "us",
				},
			},
			wantErr: false,
		},
		{
			name: "minimal feature data",
			input: dto.FeatureDto{
				ID:   789,
				Name: "Minimal Feature",
			},
			wantErr: false,
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
		})
	}
}

func TestFeatureDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		id       uint
		expected string
	}{
		{
			name:     "feature id 123",
			id:       123,
			expected: "content/feature/123.md",
		},
		{
			name:     "feature id 0",
			id:       0,
			expected: "content/feature/0.md",
		},
		{
			name:     "large feature id",
			id:       999999,
			expected: "content/feature/999999.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feature := dto.FeatureDto{
				ID: tt.id,
			}

			result := feature.OutFile()

			// Check full path
			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			// Additional checks for path construction
			dir, file := path.Split(result)
			if dir != "content/feature/" {
				t.Errorf("OutFile() directory = %s, want content/feature/", dir)
			}

			expectedFile := fmt.Sprintf("%v.md", tt.id)
			if file != expectedFile {
				t.Errorf("OutFile() filename = %s, want %s", file, expectedFile)
			}
		})
	}
}
