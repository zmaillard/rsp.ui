package types_test

import (
	"database/sql/driver"
	"encoding/json"
	"highway-sign-portal-builder/pkg/types"
	"reflect"
	"testing"
)

func TestNewImageId(t *testing.T) {
	tests := []struct {
		name     string
		imageId  string
		expected types.ImageID
	}{
		{
			name:     "numeric string",
			imageId:  "12345",
			expected: types.ImageID("12345"),
		},
		{
			name:     "empty string",
			imageId:  "",
			expected: types.ImageID(""),
		},
		{
			name:     "alphanumeric string",
			imageId:  "abc123",
			expected: types.ImageID("abc123"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := types.NewImageId(tt.imageId)
			if result != tt.expected {
				t.Errorf("NewImageId() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestImageID_String(t *testing.T) {
	tests := []struct {
		name     string
		imageId  types.ImageID
		expected string
	}{
		{
			name:     "numeric image id",
			imageId:  types.ImageID("12345"),
			expected: "12345",
		},
		{
			name:     "empty image id",
			imageId:  types.ImageID(""),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.imageId.String()
			if result != tt.expected {
				t.Errorf("String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestImageID_URLMethods(t *testing.T) {
	baseUrl := "https://example.com/"
	imageId := types.ImageID("12345")

	tests := []struct {
		name     string
		method   func(string) string
		expected string
	}{
		{
			name:     "Placeholder",
			method:   imageId.Placeholder,
			expected: "https://example.com/12345/12345_p.jpg",
		},
		{
			name:     "Thumbnail",
			method:   imageId.Thumbnail,
			expected: "https://example.com/12345/12345_t.jpg",
		},
		{
			name:     "Small",
			method:   imageId.Small,
			expected: "https://example.com/12345/12345_s.jpg",
		},
		{
			name:     "Medium",
			method:   imageId.Medium,
			expected: "https://example.com/12345/12345_m.jpg",
		},
		{
			name:     "Large",
			method:   imageId.Large,
			expected: "https://example.com/12345/12345_l.jpg",
		},
		{
			name:     "Original",
			method:   imageId.Original,
			expected: "https://example.com/12345/12345.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method(baseUrl)
			if result != tt.expected {
				t.Errorf("%s() = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestImageID_ImageS3Key(t *testing.T) {
	imageId := types.ImageID("12345")
	expected := "12345"

	result := imageId.ImageS3Key()
	if result != expected {
		t.Errorf("ImageS3Key() = %v, want %v", result, expected)
	}
}

func TestImageID_ToFullImageName(t *testing.T) {
	imageId := types.ImageID("12345")
	expected := "12345.jpg"

	result := imageId.ToFullImageName()
	if result != expected {
		t.Errorf("ToFullImageName() = %v, want %v", result, expected)
	}
}

func TestImageID_Scan(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    types.ImageID
		wantErr bool
	}{
		{
			name:    "valid int64",
			value:   int64(12345),
			want:    types.ImageID("12345"),
			wantErr: false,
		},
		{
			name:    "invalid type",
			value:   "12345",
			want:    types.ImageID(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i types.ImageID
			err := i.Scan(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && i != tt.want {
				t.Errorf("Scan() got = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestImageID_Value(t *testing.T) {
	tests := []struct {
		name    string
		imageId types.ImageID
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "valid numeric image id",
			imageId: types.ImageID("12345"),
			want:    12345,
			wantErr: false,
		},
		{
			name:    "invalid numeric image id",
			imageId: types.ImageID("abc"),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imageId.Value()

			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		imageId types.ImageID
		want    []byte
		wantErr bool
	}{
		{
			name:    "marshal numeric image id",
			imageId: types.ImageID("12345"),
			want:    []byte(`"12345"`),
			wantErr: false,
		},
		{
			name:    "marshal empty image id",
			imageId: types.ImageID(""),
			want:    []byte(`""`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imageId.MarshalJSON()

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestImageID_JSON_Integration(t *testing.T) {
	type TestStruct struct {
		ID types.ImageID `json:"id"`
	}

	// Test marshaling
	ts := TestStruct{ID: types.ImageID("12345")}
	data, err := json.Marshal(ts)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	expected := `{"id":"12345"}`
	if string(data) != expected {
		t.Errorf("JSON marshaling = %v, want %v", string(data), expected)
	}
}
