package models

import "github.com/lib/pq"

type FeatureLink struct {
	ID            uint           `gorm:"column:id;primaryKey"`
	RoadName      string         `gorm:"column:road_name"`
	FromFeatureId uint           `gorm:"column:from_feature"`
	ToFeatureId   uint           `gorm:"column:to_feature"`
	FromFeature   Feature        `gorm:"references:FromFeatureId;foreignKey:ID" `
	ToFeature     Feature        `gorm:"references:ToFeatureId;foreignKey:ID" `
	Highways      pq.StringArray `gorm:"column:highways;type:text[]"`
	FromPoint     Point          `gorm:"column:from_point"`
	ToPoint       Point          `gorm:"column:to_point"`
}

func (FeatureLink) TableName() string {
	return "sign.vwhugofeaturelink"
}
