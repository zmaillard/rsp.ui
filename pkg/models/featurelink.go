package models

import "github.com/lib/pq"

type HugoFeatureLink struct {
	ID            uint           `gorm:"column:id;primaryKey"`
	RoadName      string         `gorm:"column:road_name"`
	FromFeatureId uint           `gorm:"column:from_feature"`
	ToFeatureId   uint           `gorm:"column:to_feature"`
	FromFeature   HugoFeature    `gorm:"references:FromFeatureId;foreignKey:ID" `
	ToFeature     HugoFeature    `gorm:"references:ToFeatureId;foreignKey:ID" `
	Highways      pq.StringArray `gorm:"column:highways;type:text[]"`
}

func (HugoFeatureLink) TableName() string {
	return "vwhugofeaturelink"
}

type FeatureLink struct {
	ID            uint      `gorm:"primaryKey"`
	RoadName      string    `gorm:"column:road_name"`
	FromFeatureId uint      `gorm:"column:from_feature"`
	ToFeatureId   uint      `gorm:"column:to_feature" `
	FromFeature   Feature   `gorm:"references:FromFeatureId;foreignKey:ID" `
	ToFeature     Feature   `gorm:"references:ToFeatureId;foreignKey:ID" `
	Highways      []Highway `gorm:"many2many:feature_link_highway;foreignKey:ID;joinForeignKey:feature_link_id;References:ID;"`
}

func (FeatureLink) TableName() string {
	return "feature_link"
}

type FeatureLinkHighway struct {
	ID            uint `gorm:"primaryKey"`
	HighwayId     uint `gorm:"column:highway_id"`
	FeatureLinkId uint `gorm:"column:feature_link_id"`
	IsDescending  bool `gorm:"column:is_descending"`
}

func (FeatureLinkHighway) TableName() string {
	return "feature_link_highway"
}
