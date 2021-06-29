package db

import (
	"fmt"

	"github.com/chigaji/myfavquotes/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CreateBD() {

	db, _ := gorm.Open("sqlite3", "./sqlitedb.db")

	db.AutoMigrate(&model.FinancialQuotes{})

	// fq1 := model.FinancialQuotes{ID: 1, Quote: "Money is never enough"}

	// db.Create(&fq1)

	var q model.FinancialQuotes

	db.First(&q)

	defer db.Close()

	fmt.Println(q.ID, ":", q.Quote)

}
