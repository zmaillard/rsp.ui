package dto_test

import (
	"bytes"
	"highway-sign-portal-builder/pkg/dto"
	"path"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestHighwaySignDto_ToMarkdown(t *testing.T) {
	testTime := time.Date(2023, 5, 15, 10, 30, 0, 0, time.UTC)

	tests := []struct {
		name        string
		input       dto.HighwaySignDto
		description string
		wantErr     bool
	}{
		{
			name: "complete sign data",
			input: dto.HighwaySignDto{
				ID:          1,
				Title:       "I-90 Exit 34",
				FeatureId:   101,
				DateTaken:   testTime,
				ImageId:     "i90-exit34-2023",
				FlickrId:    stringPtr("12345678"),
				PlaceSlug:   stringPtr("north-bend"),
				StateSlug:   "washington",
				CountrySlug: "us",
				Highways:    []string{"i-90"},
				ToHighways:  []string{"wa-202"},
				GeoHash:     "c23b62w",
				ImageWidth:  1920,
				ImageHeight: 1080,
				Tags:        []string{"exit", "interstate"},
				Categories:  []string{"guide-sign"},
				Quality:     5,
				PlusCode:    "84VVMM42+",
			},
			description: "This is a sign for Exit 34 on I-90",
			wantErr:     false,
		},
		{
			name: "minimal sign data",
			input: dto.HighwaySignDto{
				ID:          2,
				Title:       "Minimal Sign",
				DateTaken:   testTime,
				ImageId:     "minimal-sign",
				StateSlug:   "oregon",
				CountrySlug: "us",
				Highways:    []string{"us-26"},
				ImageWidth:  800,
				ImageHeight: 600,
				Quality:     3,
			},
			description: "",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set description
			tt.input.Description = tt.description

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
			expected = dto.AddYamlFrontAndEndMatterText(expected, tt.description)

			// Compare
			if !bytes.Equal(result, expected) {
				t.Errorf("ToMarkdown() = %s, want %s", result, expected)
			}

			// Verify description is included correctly
			if tt.description != "" && !bytes.Contains(result, []byte(tt.description)) {
				t.Errorf("ToMarkdown() does not contain description: %s", tt.description)
			}
		})
	}
}

func TestHighwaySignDto_OutFile(t *testing.T) {
	tests := []struct {
		name     string
		imageId  string
		expected string
	}{
		{
			name:     "simple image id",
			imageId:  "i90-exit34",
			expected: "content/sign/i90-exit34.md",
		},
		{
			name:     "complex image id",
			imageId:  "us-101-california-north-exit-417b",
			expected: "content/sign/us-101-california-north-exit-417b.md",
		},
		{
			name:     "image id with special characters",
			imageId:  "ca-99_exit.123",
			expected: "content/sign/ca-99_exit.123.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sign := dto.HighwaySignDto{
				ImageId: tt.imageId,
			}

			result := sign.OutFile()

			// Check full path
			if result != tt.expected {
				t.Errorf("OutFile() = %s, want %s", result, tt.expected)
			}

			// Additional checks for path construction
			dir, file := path.Split(result)
			if dir != "content/sign/" {
				t.Errorf("OutFile() directory = %s, want content/sign/", dir)
			}

			expectedFile := sign.ImageId + ".md"
			if file != expectedFile {
				t.Errorf("OutFile() filename = %s, want %s", file, expectedFile)
			}
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
