package models

import (
	"encoding/json"
	"strings"
)

type Tags []Tag

type Tag struct {
	Id              uint    `gorm:"column:id" json:"id"`
	Name            string  `gorm:"column:name" json:"name"`
	Slug            string  `gorm:"column:slug" json:"slug"`
	IsFlickrOnly    bool    `gorm:"column:flickr_only" json:"-"`
	IsCategory      bool    `gorm:"column:is_category" json:"isCategory"`
	CategoryDetails *string `gorm:"column:category_details" json:"categoryDetails"`
}

func (t *Tag) MarshalJSON() ([]byte, error) {
	type Alias Tag
	return json.Marshal(&struct {
		*Alias
		Grouper string `json:"grouper"`
		Display string `json:"display"`
	}{
		Alias:   (*Alias)(t),
		Grouper: t.GetGrouper(),
		Display: t.GetDisplayName(),
	})
}

func (t *Tag) GetDisplayName() string {
	if t.IsCategory && t.CategoryDetails != nil {
		return *t.CategoryDetails
	}
	return t.Name
}

func (t *Tag) GetGrouper() string {
	if t.Name == "" && (t.IsCategory && t.CategoryDetails == nil) {
		return ""
	}
	var title string
	if t.IsCategory && t.CategoryDetails != nil {
		title = *t.CategoryDetails
	} else {
		title = t.Name
	}

	first := title[0:1]
	if first >= "0" && first <= "9" {
		return "0-9"
	}

	return strings.ToUpper(first)
}

func (*Tag) TableName() string {
	return "sign.tag"
}

func (Tags) OutLookupFiles() []string {
	return []string{"data/tags.json"}
}

func (ts Tags) GetLookup() ([]byte, error) {
	return json.Marshal(ts)
}
