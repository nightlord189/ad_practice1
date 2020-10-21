package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nightLord189/ad_practice1/model"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Range(min, max int) int {
	return rand.Intn(max+1-min) + min
}

func getPrices() []float32 {
	rand.Seed(time.Now().UnixNano())
	prices := make([]float32, 210)
	for i := 0; i < 210; i++ {
		if i < 105 {
			prices[i] = float32(Range(0, 200))
		} else {
			prices[i] = float32(Range(201, 3000))
		}
	}
	rand.Shuffle(len(prices), func(i, j int) { prices[i], prices[j] = prices[j], prices[i] })
	return prices
}

func GenerateArticles(db *gorm.DB) {
	fmt.Println("generating articles")
	rand.Seed(time.Now().UnixNano())
	prices := getPrices()
	for i := 1; i <= 210; i++ {
		item := model.Article{
			ID:      i,
			Name:    fmt.Sprintf("Товар %d", i),
			Package: "",
			TypeID:  ((i - 1) / 30) + 1,
			FirmID:  Range(1, 20),
			Price:   prices[i-1],
		}
		db.Create(&item)
	}
}

func getBillDate(id int) string {
	year := "2012"
	if id <= 3500 {
		year = "2010"
	} else if id <= 3500+4000 {
		year = "2011"
	}
	return fmt.Sprintf("%s--", year)
}

func getDaysCountInMonth(year, month string) int {
	yearInt, _ := strconv.Atoi(year)
	if string(month[0]) == "0" {
		month = strings.Replace(month, " ", "", -1)
	}
	monthInt, _ := strconv.Atoi(month)
	t := time.Date(yearInt, time.Month(monthInt)+1, 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

func GenerateDatesForBills() []string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 12000)
	monthsWinter := []string{"01", "02", "12"}
	monthsSpring := []string{"03", "04", "05"}
	monthsSummer := []string{"06", "07", "08"}
	monthsAutumn := []string{"09", "10", "11"}
	for i := 0; i < 12000; i++ {
		year := "2012"
		if i < 3500 {
			year = "2010"
		} else if i < 3500+4000 {
			year = "2011"
		}
		month := ""
		randMonth := Range(1, 100)
		if randMonth <= 30 {
			month = monthsWinter[Range(0, 2)]
		} else if randMonth <= 53 {
			month = monthsSpring[Range(0, 2)]
		} else if randMonth <= 76 {
			month = monthsSummer[Range(0, 2)]
		} else if randMonth <= 100 {
			month = monthsAutumn[Range(0, 2)]
		}
		day := fmt.Sprintf("%d", Range(1, getDaysCountInMonth(year, month)))
		if len(day) < 2 {
			day = "0" + day
		}
		result[i] = fmt.Sprintf("%s-%s-%s", year, month, day)
	}
	return result
}

func GenerateBills(db *gorm.DB) {
	fmt.Println("generating bills")
	dates := GenerateDatesForBills()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("rollback")
		}
	}()
	for i := 1; i <= 12000; i++ {
		item := model.Bill{
			ID:   i,
			Date: dates[i-1],
		}
		tx.Create(&item)
	}
	tx.Commit()
}

func getSeasonFromDate(date string) string {
	month := date[5:7]
	switch month {
	case "01", "02", "12":
		return "winter"
	case "03", "04", "05":
		return "spring"
	case "06", "07", "08":
		return "summer"
	case "09", "10", "11":
		return "autumn"
	}
	return ""
}

func getTypeIDWithSeason(season string) int {
	preferredTypeID := 0
	var types []int
	switch season {
	case "winter":
		types = []int{1, 2, 3, 4, 5, 6}
		preferredTypeID = 7
		break
	case "spring":
		types = []int{1, 3, 4, 5, 6, 7}
		preferredTypeID = 2
		break
	case "summer":
		types = []int{1, 2, 3, 4, 6, 7}
		preferredTypeID = 5
		break
	case "autumn":
		types = []int{1, 2, 3, 4, 5, 7}
		preferredTypeID = 6
		break
	}
	randResult := Range(0, 100)
	if randResult <= 50 {
		return preferredTypeID
	}
	return types[Range(0, 5)]
}

func getArticleByDate(articles []model.Article, date string) model.Article {
	season := getSeasonFromDate(date)
	//fmt.Println(season)
	typeID := getTypeIDWithSeason(season)
	articleID := Range((typeID-1)*30, typeID*30)
	if articleID <= 0 {
		articleID = 1
	}
	return articles[articleID-1]
}

func GenerateSales(db *gorm.DB) {
	fmt.Println("generating sales")
	var articles []model.Article
	db.Find(&articles)
	var bills []model.Bill
	findResult := db.Find(&bills)
	if findResult.Error != nil {
		panic("Get bills err: " + findResult.Error.Error())
	}
	id := 1

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("rollback")
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(bills); i++ {
		salesCount := Range(1, 5)
		for j := 0; j < salesCount; j++ {
			article := getArticleByDate(articles, bills[i].Date)
			item := model.Sale{
				ID:          id,
				Quantity:    Range(1, 2),
				BillID:      bills[i].ID,
				DrugstoreID: Range(1, 20),
				Price:       article.Price,
				ArticleID:   article.ID,
			}
			tx.Create(&item)
			id++
		}
	}
	fmt.Println("data processed, preparing to commit")
	tx.Commit()
}
