package main

import (
	"fmt"
	config "github.com/nightLord189/ad_practice1/config"
	db "github.com/nightLord189/ad_practice1/db"
)

func main() {
	fmt.Println("start")
	conf := config.Load("config.json")
	db.Launch(conf)
}
