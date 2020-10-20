package model

type Article struct {
	ID      int     `gorm:"column:id;primary_key"`
	Name    string  `gorm:"column:name"`
	Package string  `gorm:"column:package"`
	Price   float32 `gorm:"column:price"`
	TypeID  int     `gorm:"column:type_id"`
	FirmID  int     `gorm:"column:firm_id"`
}

func (u *Article) TableName() string {
	return "article"
}
