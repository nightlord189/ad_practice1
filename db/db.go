package db

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/nightLord189/ad_practice1/config"
	"github.com/nightLord189/ad_practice1/model"
)

func Launch(cfg *config.Config) *(gorm.DB) {
	var address string
	address = "host=" + cfg.DbAddress + " port=" + strconv.Itoa(cfg.DbPort) + " user=" + cfg.DbUser + " dbname=" + cfg.DbName + " password=" + cfg.DbPassword + " sslmode=disable"
	db, err := gorm.Open("postgres", address)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	//defer db.Close()
	db.AutoMigrate(
		&model.DrugType{},
		&model.Firm{},
		&model.Drugstore{},
		&model.Article{},
		&model.Bill{},
		&model.Sale{},
	)
	return db
}
