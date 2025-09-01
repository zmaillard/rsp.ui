package generator_test

import (
	"errors"
	"highway-sign-portal-builder/pkg/generator"
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
)

type mockLookup struct {
	data         []byte
	outputFiles  []string
	shouldFail   bool
	readFileFail bool
}

func (m mockLookup) GetLookup() ([]byte, error) {
	if m.shouldFail {
		return nil, errors.New("mock lookup error")
	}
	return m.data, nil
}

func (m mockLookup) OutLookupFiles() []string {
	return m.outputFiles
}

func TestSaveLookup(t *testing.T) {
	tests := []struct {
		name       string
		basePath   string
		lookup     mockLookup
		wantErr    bool
		checkFiles bool
	}{
		{
			name:     "successful save to single file",
			basePath: "/test",
			lookup: mockLookup{
				data:        []byte(`{"test": "data"}`),
				outputFiles: []string{"test_lookup.json"},
			},
			wantErr:    false,
			checkFiles: true,
		},
		{
			name:     "successful save to multiple files",
			basePath: "/test",
			lookup: mockLookup{
				data:        []byte(`{"test": "multiple"}`),
				outputFiles: []string{"test1.json", "test2.json", "test3.json"},
			},
			wantErr:    false,
			checkFiles: true,
		},
		{
			name:     "empty output files",
			basePath: "/test",
			lookup: mockLookup{
				data:        []byte(`{"test": "empty"}`),
				outputFiles: []string{},
			},
			wantErr:    false,
			checkFiles: false,
		},
		{
			name:     "error getting lookup data",
			basePath: "/test",
			lookup: mockLookup{
				shouldFail:  true,
				outputFiles: []string{"test_lookup.json"},
			},
			wantErr:    true,
			checkFiles: false,
		},
		{
			name:     "nested directory path",
			basePath: "/test/nested/dir",
			lookup: mockLookup{
				data:        []byte(`{"test": "nested"}`),
				outputFiles: []string{"nested_file.json"},
			},
			wantErr:    false,
			checkFiles: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new in-memory filesystem for each test
			appFs := afero.NewMemMapFs()

			// Create directory if it doesn't exist
			if tt.checkFiles {
				err := appFs.MkdirAll(tt.basePath, os.ModePerm)
				if err != nil {
					t.Fatalf("Failed to create directory: %v", err)
				}
			}

			// Call the function being tested
			err := generator.SaveLookup(appFs, tt.basePath, tt.lookup)

			// Check error
			if (err != nil) != tt.wantErr {
				t.Fatalf("SaveLookup() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Skip file checks if we expected an error
			if tt.wantErr {
				return
			}

			// Check if files were created with correct content
			if tt.checkFiles {
				for _, file := range tt.lookup.outputFiles {
					filePath := path.Join(tt.basePath, file)

					// Check if file exists
					exists, err := afero.Exists(appFs, filePath)
					if err != nil {
						t.Fatalf("Error checking file existence: %v", err)
					}
					if !exists {
						t.Errorf("Expected file %s to exist, but it doesn't", filePath)
						continue
					}

					// Read file content
					content, err := afero.ReadFile(appFs, filePath)
					if err != nil {
						t.Fatalf("Failed to read file %s: %v", filePath, err)
					}

					// Compare content
					if string(content) != string(tt.lookup.data) {
						t.Errorf("File %s content = %s, want %s", filePath, content, tt.lookup.data)
					}
				}
			}
		})
	}
}

func TestSaveLookup_FileWriteError(t *testing.T) {
	// Create a read-only filesystem to simulate write errors
	appFs := afero.NewReadOnlyFs(afero.NewMemMapFs())

	lookup := mockLookup{
		data:        []byte(`{"test": "data"}`),
		outputFiles: []string{"test_lookup.json"},
	}

	err := generator.SaveLookup(appFs, "/test", lookup)
	if err == nil {
		t.Errorf("SaveLookup() expected error when writing to read-only filesystem, got nil")
	}
}
