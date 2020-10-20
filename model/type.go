package model

type Type struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *Type) TableName() string {
	return "type"
}
