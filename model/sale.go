package model

type Sale struct {
	ID          int     `gorm:"column:id;primary_key"`
	Quantity    int     `gorm:"column:quantity"`
	Price       float32 `gorm:"column:price"`
	ArticleID   int     `gorm:"column:article_id"`
	BillID      int     `gorm:"column:bill_id"`
	DrugstoreID int     `gorm:"column:drugstore_id"`
}

func (u *Sale) TableName() string {
	return "sale"
}
