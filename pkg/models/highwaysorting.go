package models

type HighwaySorting struct {
	ID        uint    `gorm:"primaryKey"  json:"-" `
	SignId    uint    `gorm:"column:highwaysign_id"  json:"-" `
	HighwayId uint    `gorm:"column:highway_id" json:"-" `
	Highway   Highway `gorm:"foreignKey:HighwayId"`
	IsTo      bool    `gorm:"column:is_to"`
}

func (HighwaySorting) TableName() string {
	return "highwaysign_highway"
}
