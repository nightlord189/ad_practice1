package model

type Sale struct {
	ID        int `gorm:"column:id;primary_key"`
	Quantity  int `gorm:"column:quantity"`
	ArticleID int `gorm:"column:article_id"`
	BillID    int `gorm:"column:bill_id"`
	DrugID    int `gorm:"column:drug_id"`
}

func (u *Sale) TableName() string {
	return "sale"
}
