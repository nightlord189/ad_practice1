package model

type Bill struct {
	ID       int    `gorm:"column:id;primary_key"`
	Date     string `gorm:"column:date"`
	Discount *int   `gorm:"column:discount"`
}

func (u *Bill) TableName() string {
	return "bill"
}
