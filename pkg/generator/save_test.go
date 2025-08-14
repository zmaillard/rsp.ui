package generator_test

import (
	"errors"
	"highway-sign-portal-builder/pkg/generator"
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
)

type mockGenerator struct {
	data       []byte
	outputFile string
	shouldFail bool
}

func (m mockGenerator) ToMarkdown() ([]byte, error) {
	if m.shouldFail {
		return nil, errors.New("mock generator error")
	}
	return m.data, nil
}

func (m mockGenerator) OutFile() string {
	return m.outputFile
}

func TestSaveItem(t *testing.T) {
	tests := []struct {
		name        string
		basePath    string
		generator   mockGenerator
		wantErr     bool
		checkExists bool
	}{
		{
			name:     "successful save",
			basePath: "/test",
			generator: mockGenerator{
				data:       []byte("# Test Markdown"),
				outputFile: "test.md",
			},
			wantErr:     false,
			checkExists: true,
		},
		{
			name:     "nested directory path",
			basePath: "/test",
			generator: mockGenerator{
				data:       []byte("# Nested Test"),
				outputFile: "nested/dir/test.md",
			},
			wantErr:     false,
			checkExists: true,
		},
		{
			name:     "error generating markdown",
			basePath: "/test",
			generator: mockGenerator{
				shouldFail: true,
				outputFile: "error.md",
			},
			wantErr:     true,
			checkExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new in-memory filesystem for each test
			fs := afero.NewMemMapFs()

			// Create base directory
			err := fs.MkdirAll(tt.basePath, os.ModePerm)
			if err != nil {
				t.Fatalf("Failed to create base directory: %v", err)
			}

			// Call the function being tested
			err = generator.SaveItem(fs, tt.basePath, tt.generator)

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Skip file checks if we expected an error
			if tt.wantErr {
				return
			}

			// Check if file was created with correct content
			if tt.checkExists {
				filePath := path.Join(tt.basePath, tt.generator.outputFile)

				// Check if file exists
				exists, err := afero.Exists(fs, filePath)
				if err != nil {
					t.Fatalf("Error checking file existence: %v", err)
				}
				if !exists {
					t.Errorf("Expected file %s to exist, but it doesn't", filePath)
					return
				}

				// Read file content
				content, err := afero.ReadFile(fs, filePath)
				if err != nil {
					t.Fatalf("Failed to read file %s: %v", filePath, err)
				}

				// Compare content
				if string(content) != string(tt.generator.data) {
					t.Errorf("File %s content = %s, want %s", filePath, content, tt.generator.data)
				}
			}
		})
	}
}

func TestSaveItem_FileWriteError(t *testing.T) {
	// Create a read-only filesystem to simulate write errors
	fs := afero.NewReadOnlyFs(afero.NewMemMapFs())

	mockGenerator := mockGenerator{
		data:       []byte("# Test Markdown"),
		outputFile: "test.md",
	}

	err := generator.SaveItem(fs, "/test", mockGenerator)
	if err == nil {
		t.Errorf("SaveItem() expected error when writing to read-only filesystem, got nil")
	}
}
