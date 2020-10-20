package db

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/nightLord189/ad_practice1/config"
	"github.com/nightLord189/ad_practice1/model"
)

func Launch(cfg *config.Config) *gorm.DB {
	var address string
	address = "host=" + cfg.DbAddress + " port=" + strconv.Itoa(cfg.DbPort) + " user=" + cfg.DbUser + " dbname=" + cfg.DbName + " password=" + cfg.DbPassword + " sslmode=disable"
	db, err := gorm.Open("postgres", address)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	//defer db.Close()
	AutoMigrate(db)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Type{},
		&model.Firm{},
		&model.Drugstore{},
		&model.Article{},
		&model.Bill{},
		&model.Sale{},
	)
	db.Model(&model.Article{}).AddForeignKey("type_id", "type(id)", "NO ACTION", "NO ACTION")
	db.Model(&model.Article{}).AddForeignKey("firm_id", "firm(id)", "NO ACTION", "NO ACTION")
	db.Model(&model.Sale{}).AddForeignKey("article_id", "article(id)", "NO ACTION", "NO ACTION")
	db.Model(&model.Sale{}).AddForeignKey("bill_id", "bill(id)", "NO ACTION", "NO ACTION")
	db.Model(&model.Sale{}).AddForeignKey("drugstore_id", "drugstore(id)", "NO ACTION", "NO ACTION")
}
