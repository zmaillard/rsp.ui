package converter

import (
	"context"
	"encoding/json"
	"highway-sign-portal-builder/pkg/db"
	"highway-sign-portal-builder/pkg/generator"
)

type HighwayNames []highwayName

type highwayName struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	StateName string `json:"stateName"`
	StateSlug string `json:"stateSlug"`
}

func NewHighwayNameConverter(ctx context.Context, queries db.Querier) (generator.Lookup, error) {
	hwyNames, err := queries.GetHugoHighwayNames(ctx)
	if err != nil {
		return nil, err
	}

	var hwyNameDtos []highwayName
	for i := range hwyNames {
		tagDto := highwayName{
			Id:        uint(hwyNames[i].ID),
			Name:      hwyNames[i].Name.String,
			Slug:      hwyNames[i].Slug,
			StateName: hwyNames[i].StateName.String,
			StateSlug: hwyNames[i].StateSlug.String,
		}
		hwyNameDtos = append(hwyNameDtos, tagDto)
	}

	return HighwayNames(hwyNameDtos), nil
}

func (HighwayNames) OutLookupFiles() []string {
	return []string{"data/highwaynames.json"}
}

func (hn HighwayNames) GetLookup() ([]byte, error) {
	return json.Marshal(hn)
}
