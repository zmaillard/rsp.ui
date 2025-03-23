package models

import (
	"fmt"
	"strconv"
	"strings"
)

func getFromTo(prev []int32, next []int32) []int32 {
	if len(prev) != len(next) {
		return []int32{}
	}

	if len(prev) == 0 {
		return []int32{}
	}

	allFeatCount := len(prev) + 1
	out := make([]int32, allFeatCount)
	for i, _ := range prev {
		out[i] = prev[i]
	}

	out[allFeatCount-1] = next[len(next)-1]
	return out
}

func getFrom(in []fromTo) []uint {
	allFeatCount := len(in) + 1
	out := make([]uint, allFeatCount)
	for i, _ := range in {
		out[i] = in[i].From
	}

	out[allFeatCount-1] = in[len(in)-1].To
	return out
}

type fromTo struct {
	From uint
	To   uint
}

func (f *fromTo) handleUintArray(rawArray []uint8) error {
	foo := []byte(rawArray)
	baz := string(foo)
	fmt.Println(baz)
	return nil
}

func (f *fromTo) handleString(rawArray string) error {

	//value should look like this
	// (19624,19625)
	if len(rawArray) == 0 {
		return nil
	}

	cleaned := strings.Replace(strings.Replace(rawArray, "(", "", 1), ")", "", 1)
	splitStr := strings.Split(cleaned, ",")

	if len(splitStr) != 2 {
		return nil
	}

	fromInt, err := strconv.Atoi(splitStr[0])
	if err != nil {
		return err
	}

	toInt, err := strconv.Atoi(splitStr[1])
	if err != nil {
		return err
	}

	f.From = uint(fromInt)
	f.To = uint(toInt)

	return nil
}

func (f *fromTo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case string:
		return f.handleString(value.(string))
	case []uint8:
		return f.handleString(string(value.([]byte)))

	}

	return fmt.Errorf("unable to handle fromto")
}
