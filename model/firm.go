package model

type Firm struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *Firm) TableName() string {
	return "firm"
}
