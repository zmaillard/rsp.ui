package models

import (
	"encoding/json"
)

type Tags []Tag

type Tag struct {
	Id           uint   `gorm:"column:id"`
	Name         string `gorm:"column:name"`
	Slug         string `gorm:"column:slug"`
	IsFlickrOnly bool   `gorm:"column:flickr_only"`
}

func (*Tag) TableName() string {
	return "sign.tag"
}

func (Tags) OutLookupFiles() []string {
	return []string{"data/tags.json"}
}

func (ts Tags) GetLookup() ([]byte, error) {
	tags := make(map[string]string)

	for _, v := range ts {
		tags[v.Slug] = v.Name

	}

	return json.Marshal(tags)
}
