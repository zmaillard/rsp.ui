package dto_test

import (
	"bytes"
	"highway-sign-portal-builder/pkg/dto"
	"path"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestHighwayDto_ToMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.HighwayDto
		wantErr bool
	}{
		{
			name: "complete highway data",
			input: dto.HighwayDto{
				Name:        "Interstate 90",
				DisplayName: "I-90",
				Slug:        "i90",
				Image:       "i90-featured.jpg",
				HighwayTypeSlug: dto.AdminAreaSlimDto{
					Slug: "interstate",
				},
				Sort:     10,
				Features: []int32{101, 102, 103},
				Places:   []string{"seattle", "ellensburg", "spokane"},
				States:   []string{"washington", "idaho", "montana"},
				Counties: []string{"king", "kittitas", "spokane"},
				Aliases:  []string{"I-90", "Interstate 90"},
			},
			wantErr: false,
		},
		{
			name: "minimal highway data",
			input: dto.HighwayDto{
				Name:  "US Highway 101",
				Slug:  "us101",
				Image: "us101.jpg",
				HighwayTypeSlug: dto.AdminAreaSlimDto{
					Slug: "us-highway",
				},
				Sort: 20,
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

func TestHighwayDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		slug     string
		expected string
	}{
		{
			name:     "simple slug",
			slug:     "i90",
			expected: "content/highway/i90/_index.md",
		},
		{
			name:     "hyphenated slug",
			slug:     "us-101",
			expected: "content/highway/us-101/_index.md",
		},
		{
			name:     "state highway slug",
			slug:     "wa520",
			expected: "content/highway/wa520/_index.md",
		},
		{
			name:     "complex slug",
			slug:     "california-sr1",
			expected: "content/highway/california-sr1/_index.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			highway := dto.HighwayDto{
				Slug: tt.slug,
			}

			result := highway.OutFile()

			// Check full path
			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			// Additional checks for path construction
			dir, file := path.Split(result)
			expectedDir := "content/highway/" + tt.slug + "/"
			if dir != expectedDir {
				t.Errorf("OutFile() directory = %s, want %s", dir, expectedDir)
			}

			if file != "_index.md" {
				t.Errorf("OutFile() filename = %s, want _index.md", file)
			}
		})
	}
}
