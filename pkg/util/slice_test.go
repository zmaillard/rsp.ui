package util_test

import (
	"highway-sign-portal-builder/pkg/util"
	"reflect"
	"strconv"
	"testing"
)

func TestSliceMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapFunc  func(int) string
		expected []string
	}{
		{
			name:     "map integers to strings",
			input:    []int{1, 2, 3, 4, 5},
			mapFunc:  func(i int) string { return strconv.Itoa(i) },
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:     "map integers to doubled integers as strings",
			input:    []int{1, 2, 3},
			mapFunc:  func(i int) string { return strconv.Itoa(i * 2) },
			expected: []string{"2", "4", "6"},
		},
		{
			name:     "empty slice",
			input:    []int{},
			mapFunc:  func(i int) string { return strconv.Itoa(i) },
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SliceMap(tt.input, tt.mapFunc)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SliceMap() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSliceMapWithDifferentTypes(t *testing.T) {
	// Test with struct types
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	extractNames := func(p Person) string {
		return p.Name
	}

	expectedNames := []string{"Alice", "Bob", "Charlie"}
	names := util.SliceMap(people, extractNames)

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf("SliceMap() = %v, want %v", names, expectedNames)
	}

	// Test with pointers
	nums := []*int{new(int), new(int), new(int)}
	*nums[0] = 1
	*nums[1] = 2
	*nums[2] = 3

	deref := func(p *int) int {
		return *p
	}

	expectedInts := []int{1, 2, 3}
	ints := util.SliceMap(nums, deref)

	if !reflect.DeepEqual(ints, expectedInts) {
		t.Errorf("SliceMap() = %v, want %v", ints, expectedInts)
	}
}
