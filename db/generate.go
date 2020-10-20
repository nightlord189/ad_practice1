package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nightLord189/ad_practice1/model"
	"math/rand"
	"time"
)

func getPrices() []float32 {
	prices := make([]float32, 210)
	for i := 0; i < 210; i++ {
		if i < 105 {
			prices[i] = float32(rand.Intn(200 - 0))
		} else {
			prices[i] = float32(rand.Intn(3000-201) + 201)
		}
	}
	rand.Shuffle(len(prices), func(i, j int) { prices[i], prices[j] = prices[j], prices[i] })
	return prices
}

func GenerateArticles(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	prices := getPrices()
	for i := 1; i <= 210; i++ {
		item := model.Article{
			ID:      i,
			Name:    fmt.Sprintf("Товар %d", i),
			Package: "",
			TypeID:  ((i - 1) / 30) + 1,
			FirmID:  rand.Intn(21-1) + 1,
			Price:   prices[i-1],
		}
		fmt.Println(i)
		db.Create(&item)
	}
}
