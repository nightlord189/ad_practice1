package model

type DrugType struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *DrugType) TableName() string {
	return "drug_type"
}
