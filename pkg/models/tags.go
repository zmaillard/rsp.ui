package models

type Tag struct {
	Id           uint   `gorm:"column:id"`
	Name         string `gorm:"column:name"`
	Slug         string `gorm:"column:slug"`
	IsFlickrOnly bool   `gorm:"column:flickr_only"`
}

func (*Tag) TableName() string {
	return "tag"
}

type HighwaySignTag struct {
	Id     uint `gorm:"column:id"`
	SignId uint `gorm:"column:highwaysign_id"  json:"-"`
	TagId  uint `gorm:"column:tag_id" json:"-"`
	Tag    Tag  `gorm:"foreignKey:TagId"`
}

func (*HighwaySignTag) TableName() string {
	return "tag_highwaysign"
}
