package main

import (
	"fmt"
	config "github.com/nightLord189/ad_practice1/config"
	db "github.com/nightLord189/ad_practice1/db"
)

func main() {
	fmt.Println("start")
	conf := config.Load("config.json")
	sqlManager := db.NewSQLManager(conf.SQLPath)
	dbRef := db.Launch(conf)
	dbRef.Exec(sqlManager.Data["clear"])
	dbRef.Exec(sqlManager.Data["init"])
	db.GenerateArticles(dbRef)
}
