package converter

import (
	"context"
	"encoding/json"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/generator"
	"strings"
)

type Tags []tag

type tag struct {
	Id              uint    `json:"id"`
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	IsCategory      bool    `json:"isCategory"`
	CategoryDetails *string `json:"categoryDetails"`
}

func NewTagConverter(ctx context.Context, queries db.Querier) (generator.Lookup, error) {
	tags, err := queries.GetHugoTags(ctx)
	if err != nil {
		return nil, err
	}

	var tagDtos []tag
	for i := range tags {
		tagDto := tag{
			Id:         uint(tags[i].ID),
			Name:       tags[i].Name.String,
			Slug:       tags[i].Slug.String,
			IsCategory: tags[i].IsCategory,
		}
		if tags[i].CategoryDetails.Valid {
			tagDto.CategoryDetails = &tags[i].CategoryDetails.String
		}
		tagDtos = append(tagDtos, tagDto)
	}

	return Tags(tagDtos), nil
}

func (t *tag) MarshalJSON() ([]byte, error) {
	type Alias tag
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

func (t *tag) GetDisplayName() string {
	if t.IsCategory && t.CategoryDetails != nil {
		return *t.CategoryDetails
	}
	return t.Name
}

func (t *tag) GetGrouper() string {
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

func (Tags) OutLookupFiles() []string {
	return []string{"data/tags.json"}
}

func (ts Tags) GetLookup() ([]byte, error) {
	return json.Marshal(ts)
}
