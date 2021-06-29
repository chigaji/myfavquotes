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

	db.LogMode(true)

	// db.AutoMigrate(&model.FinancialQuotes{})

	// fq1 := model.FinancialQuotes{ID: 1, Quote: "Money is never enough"}

	// db.Create(&fq1)

	// var q model.FinancialQuotes

	// db.First(&q)

	// create tables

	if !db.HasTable(&model.FinancialQuotes{}) {
		db.CreateTable(&model.FinancialQuotes{})
	}

	if !db.HasTable(&model.LifeQuotes{}) {
		db.CreateTable(&model.LifeQuotes{})
	}
	if !db.HasTable(&model.LoveQuotes{}) {
		db.CreateTable(&model.LoveQuotes{})
	}
	defer db.Close()

	return db
}
