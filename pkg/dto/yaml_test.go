package dto_test

import (
	"bytes"
	"highway-sign-portal-builder/pkg/dto"
	"testing"
)

func TestAddYamlFrontAndEndMatter(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "empty data",
			input:    []byte(""),
			expected: []byte("---\n\n---"),
		},
		{
			name:     "simple data",
			input:    []byte("title: Test"),
			expected: []byte("---\ntitle: Test\n---"),
		},
		{
			name:     "multiline data",
			input:    []byte("title: Test\ndescription: Test description"),
			expected: []byte("---\ntitle: Test\ndescription: Test description\n---"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dto.AddYamlFrontAndEndMatter(tt.input)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("AddYamlFrontAndEndMatter() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestAddYamlFrontAndEndMatterText(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		text     string
		expected []byte
	}{
		{
			name:     "empty data with empty text",
			input:    []byte(""),
			text:     "",
			expected: []byte("---\n\n---\n"),
		},
		{
			name:     "simple data with text",
			input:    []byte("title: Test"),
			text:     "This is content",
			expected: []byte("---\ntitle: Test\n---\nThis is content"),
		},
		{
			name:     "multiline data with multiline text",
			input:    []byte("title: Test\ndescription: Test description"),
			text:     "Line 1\nLine 2\nLine 3",
			expected: []byte("---\ntitle: Test\ndescription: Test description\n---\nLine 1\nLine 2\nLine 3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dto.AddYamlFrontAndEndMatterText(tt.input, tt.text)
			if !bytes.Equal(result, tt.expected) {
				t.Errorf("AddYamlFrontAndEndMatterText() = %q, want %q", result, tt.expected)
			}
		})
	}
}
