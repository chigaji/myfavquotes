package db

import (
	"github.com/chigaji/myfavquotes/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func SetUpDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "./sqlitedb.db")

	checkError(err)
	// defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&model.FinancialQuotes{})
	db.AutoMigrate(&model.LifeQuotes{})
	db.AutoMigrate(&model.LoveQuotes{})

	return db
}
