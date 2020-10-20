package model

type Drugstore struct {
	ID      int    `gorm:"column:id;primary_key"`
	Address string `gorm:"column:address"`
	Phone   string `gorm:"column:phone"`
}

func (u *Drugstore) TableName() string {
	return "drugstore"
}
