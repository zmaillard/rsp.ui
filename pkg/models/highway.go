package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/dto"
	"highway-sign-portal-builder/pkg/generator"
)

type HugoHighway struct {
	ID               uint           `gorm:"column:id,primaryKey"`
	Name             string         `gorm:"column:highway_name"`
	Slug             string         `gorm:"column:slug"`
	Sort             int            `gorm:"column:sort_number"`
	Image            string         `gorm:"column:image_name"`
	HighwayTypeSlug  string         `gorm:"column:highway_type_slug"`
	HighwayTypeName  string         `gorm:"column:highway_type_name"`
	States           pq.StringArray `gorm:"column:states;type:text[]"`
	Counties         pq.StringArray `gorm:"column:counties;type:text[]"`
	Places           pq.StringArray `gorm:"column:places;type:text[]"`
	PreviousFeatures pq.Int32Array  `gorm:"column:previous_features;type:int[]"`
	NextFeatures     pq.Int32Array  `gorm:"column:next_features;type:int[]"`
}

func (HugoHighway) TableName() string {
	return "vwhugohighway"
}

func (h HugoHighway) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayDto{
		Name:  h.Name,
		Slug:  h.Slug,
		Image: h.Image,
		Sort:  h.Sort,
		HighwayTypeSlug: dto.AdminAreaSlimDto{
			Name: h.HighwayTypeName,
			Slug: h.HighwayTypeSlug,
		},
		Features: getFromTo(h.PreviousFeatures, h.NextFeatures),
		Places:   h.Places,
		States:   h.States,
		Counties: h.Counties,
	}

	var aliases []string
	for _, v := range h.States {
		aliases = append(aliases, "/statehighway/"+v+"/"+h.Name)
	}

	highwayDto.Aliases = aliases

	return highwayDto
}

type Highway struct {
	ID            uint        `gorm:"primaryKey" yaml:"-"`
	Name          string      `gorm:"column:highway_name" yaml:"name"`
	Slug          string      `gorm:"column:slug" yaml:"slug"`
	Sort          int         `gorm:"column:sort_number"`
	Image         string      `gorm:"column:image_name"  yaml:"imageName"`
	HighwayTypeId uint        `gorm:"column:highway_type_id" yaml:"-"`
	HighwayType   HighwayType `gorm:"foreignKey:HighwayTypeId" yaml:"highwayType,omitempty"`
	ScopeId       uint        `gorm:"column:scope_id" yaml:"-"`
	Scope         Scope       `gorm:"foreignKey:ScopeId" yaml:"scope,omitempty"`
	CountryId     uint        `gorm:"column:admin_area_country_id" `
	StateId       *uint       `gorm:"column:admin_area_state_id" `
	Features      []uint      `gorm:"-" yaml:"features,omitempty"`
	States        []string    `gorm:"-"`
	Counties      []string    `gorm:"-"`
	Places        []string    `gorm:"-"`
}

func (Highway) TableName() string {
	return "highway"
}

func (h *Highway) AfterFind(tx *gorm.DB) (err error) {
	// Association not loaded - don't bother loading after find
	if h.HighwayType.Name == "" {
		return
	}

	/*
		var places []string
		err = tx.Debug().
			Raw("SELECT DISTINCT st.slug || '_' || p.slug FROM highwaysign_highway hs INNER JOIN highwaysign s on hs.highwaysign_id = s.id INNER JOIN admin_area_place p on s.admin_area_place_id = p.id INNER JOIN admin_area_state st on p.admin_area_stateid = st.id  WHERE hs.highway_id = ? AND hs.is_to = false", h.ID).
			Find(&places).Error
		h.Places = places

		var counties []string
		err = tx.Debug().
			Raw("SELECT DISTINCT st.slug || '_' || c.slug FROM highwaysign_highway hs INNER JOIN highwaysign s on hs.highwaysign_id = s.id INNER JOIN admin_area_county c on s.admin_area_county_id = c.id INNER JOIN admin_area_state st on c.admin_area_stateid = st.id  WHERE hs.highway_id = ?  AND hs.is_to = false", h.ID).
			Find(&counties).Error
		h.Counties = counties

		var states []string
		err = tx.Debug().
			Raw("SELECT DISTINCT st.slug FROM highwaysign_highway hs INNER JOIN highwaysign s on hs.highwaysign_id = s.id INNER JOIN admin_area_state st on s.admin_area_state_id = st.id  WHERE hs.highway_id = ? AND hs.is_to = false", h.ID).
			Find(&states).Error
		h.States = states
	*/
	var featureList []fromTo
	err = tx.Debug().
		Raw("SELECT get_first_highway(?)", h.ID).
		Find(&featureList).Error

	if err != nil {
		return err
	}

	if len(featureList) != 1 {
		return nil
	}

	fl := featureList[0]

	var orderedFeatures []fromTo
	err = tx.Debug().
		Raw("SELECT get_ordered_features(?,?)", h.ID, fl.From).
		Find(&orderedFeatures).Error

	h.Features = getFrom(orderedFeatures)

	return nil
}

type Highways []Highway

func (h Highways) GetLookup() ([]byte, error) {
	slugs := make(map[string]string)
	names := make(map[string]string)

	for _, v := range h {
		slugs[v.Name] = v.Slug
		names[v.Slug] = v.Name
	}
	res := map[string]interface{}{
		"slug": slugs,
		"name": names,
	}

	return json.Marshal(res)
}

func (Highways) OutLookupFile() string {
	return "data/highway.json"
}

func (h Highway) ConvertToDto() generator.Generator {
	highwayDto := dto.HighwayDto{
		Name:  h.Name,
		Slug:  h.Slug,
		Image: h.Image,
		Sort:  h.Sort,
		HighwayTypeSlug: dto.AdminAreaSlimDto{
			Name: h.HighwayType.Name,
			Slug: h.HighwayType.Slug,
		},
		//Features: h.Features,
		Places:   h.Places,
		States:   h.States,
		Counties: h.Counties,
	}

	var aliases []string
	for _, v := range h.States {
		aliases = append(aliases, "/statehighway/"+v+"/"+h.Name)
	}

	highwayDto.Aliases = aliases

	return highwayDto
}
