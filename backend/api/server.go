package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/chigaji/myfavquotes/backend/db"
	"github.com/chigaji/myfavquotes/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RunServer() {
	r := gin.Default()
	api := r.Group("/api")

	//set db connection
	_ = db.SetUpDB()
	//api endpoints

	// home
	api.GET("/", handleHome)

	// get all quotes of a given type
	api.GET("/quotes/:type", getQuotes)

	// get a random quote of a given type
	api.GET("/quote/:type/", getQuote)

	// create a quote of a given type
	api.POST("/quotes/:type", createQuote)

	// edit a given quote of a type using a given id
	api.PUT("/quotes/:type/:id")

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	r.Run(":4000")

}

// func getQuote(c *gin.Context) {
// 	fmt.Println("Get quotes called....")
// 	quoteType := c.Param("type")
// 	id := c.Param("id")

// 	db := getDbCOnn()
// 	defer db.Close()

// 	// switch requet quotetype
// 	switch strings.ToLower(quoteType) {
// 	case "finance":
// 		var quote model.FinancialQuotes
// 		// SELECT * FROM financial_quote WHERE id = 1;
// 		err := db.First(&quote, id).Error
// 		checkError(err, c)
// 		c.JSON(200, quote)

// 	case "love":
// 		var quote model.LoveQuotes
// 		err := db.First(&quote, id).Error
// 		checkError(err, c)
// 		c.JSON(200, quote)

// 	case "life":
// 		var quote model.LifeQuotes
// 		err := db.First(&quote, id).Error
// 		checkError(err, c)
// 		c.JSON(200, quote)
// 	default:
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown quote type"})
// 	}

// }

func getQuote(c *gin.Context) {
	fmt.Println("Get quote called....")
	quoteType := c.Param("type")

	db := getDbCOnn()
	defer db.Close()

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		var quotes []model.FinancialQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)

		rand.Seed(time.Now().UnixNano()) //initial a random generator

		c.JSON(200, quotes[rand.Intn(len((quotes)))])

	case "love":
		var quotes []model.LoveQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)
		rand.Seed(time.Now().UnixNano()) //initial a random generator
		c.JSON(200, quotes[rand.Intn(len((quotes)))])

	case "life":
		var quotes []model.LifeQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)
		rand.Seed(time.Now().UnixNano()) //initial a random generator
		c.JSON(200, quotes[rand.Intn(len((quotes)))])
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown quote type"})
	}
	// quote := "Money can't solve all your problems, but neither can poverty"

}

func getQuotes(c *gin.Context) {
	fmt.Println("Get quotes called....")
	quoteType := c.Param("type")

	db := getDbCOnn()
	defer db.Close()

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		var quotes []model.FinancialQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)
		c.JSON(200, quotes)

	case "love":
		var quotes []model.LoveQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)
		c.JSON(200, quotes)

	case "life":
		var quotes []model.LifeQuotes
		err := db.Find(&quotes).Error
		checkError(err, c)
		c.JSON(200, quotes)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown quote type"})
	}
	// quote := "Money can't solve all your problems, but neither can poverty"

}

func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is an app that gets your favorite quotes depending on topic of your interest",
	})
}

func createQuote(c *gin.Context) {

	db := getDbCOnn()

	defer db.Close()

	//get quote type (financial, love, or life quote)
	quoteType := c.Param("type")

	println("got quote type: ", quoteType)
	// db.Prepa
	// quote := "Money can't solve all your problems, but neither can poverty"
	quote, err := ioutil.ReadAll(c.Request.Body)
	checkError(err, c)

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		financialQuote := model.FinancialQuotes{}
		err := json.Unmarshal(quote, &financialQuote)
		checkError(err, c)

		//insert financial quote to the db
		err = db.Create(&financialQuote).Error
		checkError(err, c)

		c.JSON(201, gin.H{"success": financialQuote})

	case "love":
		loveQuote := model.LoveQuotes{}
		err := json.Unmarshal(quote, &loveQuote)
		checkError(err, c)

		//insert love quote to the db
		err = db.Create(&loveQuote).Error
		checkError(err, c)

		c.JSON(201, gin.H{"success": loveQuote})

	case "life":
		lifeQuote := model.LifeQuotes{}
		err := json.Unmarshal(quote, &lifeQuote)
		checkError(err, c)

		//insert love quote to the db
		err = db.Create(&lifeQuote).Error
		checkError(err, c)

		c.JSON(201, gin.H{"success": lifeQuote})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid quote Type"})
	}
}
func getDbCOnn() *gorm.DB {
	db := db.SetUpDB()
	return db
}

func checkError(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err ": err})
		panic(err)
	}
}
