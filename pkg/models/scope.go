package models

type Scope struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"column:scope"`
	Highways []Highway `gorm:"foreignKey:ScopeId" json:",omitempty"`
}

func (ht *Scope) TableName() string {
	return "highway_scope"
}
