package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/chigaji/myfavquotes/backend/db"
	_ "github.com/chigaji/myfavquotes/backend/docs"

	// _ "../docs"
	"github.com/chigaji/myfavquotes/backend/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Favorite Quotes
// @version 1.0
// @description This is an API for My favorite quotes
// @termsOfService http://swagger.io/terms/

// @contact.name Ronald
// @contact.url brothermen.com
// @contact.email chigaji@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1
// @schemes http
func RunServer() {
	r := gin.Default()

	//add cors here
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	api := r.Group("/api/v1")

	//set db connection
	_ = db.SetUpDB()
	//api endpoints

	// home
	api.GET("/home", handleHome)

	// get a random quote of a given type
	api.GET("/quote/:type", getRandomQuote)

	// get all quotes of a given type
	api.GET("/quotes/:type", getQuotes)

	// get a quote of given type by id
	api.GET("/quote/:type/:id", getQuote)

	// create a quote of a given type
	api.POST("/quotes/:type", createQuote)

	// edit a given quote of a type using a given id
	api.PUT("/quotes/:type/:id", editQuote)

	// Delete a quote of a given type by Id
	api.DELETE("/quote/:type/:id", deleteQuote)

	// add swagger
	url := ginSwagger.URL("http://localhost:4000/api/v1/api-docs/doc.json")
	api.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	r.Run(":4000")

}

func getQuote(c *gin.Context) {
	fmt.Println("Get quote called....")
	quoteType := c.Param("type")

	id := c.Param("id")

	db := getDbCOnn()
	defer db.Close()

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		var quote model.FinancialQuotes
		err := db.First(&quote, id).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quote)
		}

	case "love":
		var quote model.LoveQuotes
		err := db.First(&quote, id).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quote)
		}
	case "life":
		var quote model.LifeQuotes
		err := db.First(&quote, id).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quote)
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown quote type"})
	}
	// quote := "Money can't solve all your problems, but neither can poverty"

}
func getRandomQuote(c *gin.Context) {
	fmt.Println("Get Random quote called....")
	quoteType := c.Param("type")

	db := getDbCOnn()
	defer db.Close()

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		var quotes []model.FinancialQuotes
		err := db.Find(&quotes).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			rand.Seed(time.Now().UnixNano()) //initial a random generator
			c.JSON(200, quotes[rand.Intn(len((quotes)))])
		}

	case "love":
		var quotes []model.LoveQuotes
		err := db.Find(&quotes).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			rand.Seed(time.Now().UnixNano()) //initial a random generator
			c.JSON(200, quotes[rand.Intn(len((quotes)))])
		}

	case "life":
		var quotes []model.LifeQuotes
		err := db.Find(&quotes).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			rand.Seed(time.Now().UnixNano()) //initial a random generator
			c.JSON(200, quotes[rand.Intn(len((quotes)))])
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown quote type"})
	}
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
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quotes)
		}

	case "love":
		var quotes []model.LoveQuotes
		err := db.Find(&quotes).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quotes)
		}

	case "life":
		var quotes []model.LifeQuotes
		err := db.Find(&quotes).Error
		if err != nil {
			notFoundError(err, c)
		} else {
			c.JSON(200, quotes)
		}
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

		if err != nil {
			checkError(err, c)
		} else {

			c.JSON(201, gin.H{"success": financialQuote})
		}

	case "love":
		loveQuote := model.LoveQuotes{}
		err := json.Unmarshal(quote, &loveQuote)
		checkError(err, c)

		//insert love quote to the db
		err = db.Create(&loveQuote).Error
		if err != nil {
			checkError(err, c)
		} else {
			c.JSON(201, gin.H{"success": loveQuote})
		}
	case "life":
		lifeQuote := model.LifeQuotes{}
		err := json.Unmarshal(quote, &lifeQuote)
		checkError(err, c)

		//insert love quote to the db
		err = db.Create(&lifeQuote).Error
		if err != nil {
			checkError(err, c)
		} else {
			c.JSON(201, gin.H{"success": lifeQuote})
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid quote Type"})
	}
}

/*
* edit quote
 */
func editQuote(c *gin.Context) {

	db := getDbCOnn()

	defer db.Close()

	//get quote type (financial, love, or life quote)
	quoteType := c.Param("type")

	// get quote id
	id := c.Param("id")

	println("got quote type: ", quoteType)
	// db.Prepa
	// quote := "Money can't solve all your problems, but neither can poverty"
	// quote, err := ioutil.ReadAll(c.Request.Body)
	// notFoundError(err, c)

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		financialQuote := model.FinancialQuotes{}

		// SELECT * FROM financial_quote WHERE id = 1;
		db.First(&financialQuote, id)

		if financialQuote.Quote != "" {
			println(financialQuote.ID)
			if financialQuote.ID != 0 {
				var newQuote model.FinancialQuotes
				c.Bind(&newQuote)
				println(newQuote.Quote, newQuote.ID)
				result := model.FinancialQuotes{
					ID:    financialQuote.ID,
					Quote: newQuote.Quote,
				}
				db.Save(&result)

				c.JSON(200, gin.H{"success": result})
			} else {
				c.JSON(404, gin.H{"error": "quote with Id: " + id + "not Found"})
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}

	case "love":
		loveQuote := model.LoveQuotes{}

		// SELECT * FROM love_quotes WHERE id = 1;
		db.First(&loveQuote, id)

		if loveQuote.Quote != "" {

			if loveQuote.ID != 0 {
				var newQuote model.LoveQuotes
				c.Bind(&newQuote)

				result := model.LoveQuotes{
					ID:    loveQuote.ID,
					Quote: newQuote.Quote,
				}
				db.Save(&result)

				c.JSON(200, gin.H{"success": result})
			} else {
				c.JSON(404, gin.H{"error": "quote of type " + quoteType + " with Id: " + id + "not Found"})
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}

	case "life":
		lifeQuote := model.LifeQuotes{}

		// SELECT * FROM love_quotes WHERE id = 1;
		db.First(&lifeQuote, id)
		if lifeQuote.Quote != "" {

			if lifeQuote.ID != 0 {
				var newQuote model.LifeQuotes
				c.Bind(&newQuote)

				result := model.LifeQuotes{
					ID:    lifeQuote.ID,
					Quote: newQuote.Quote,
				}
				db.Save(&result)

				c.JSON(200, gin.H{"success": result})
			} else {
				c.JSON(404, gin.H{"error": "quote of type " + quoteType + " with Id: " + id + "not Found"})
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid quote Type"})
	}
}

/*
* Delete quote of type using id
 */
func deleteQuote(c *gin.Context) {

	db := getDbCOnn()

	defer db.Close()

	//get quote type (financial, love, or life quote)
	quoteType := c.Param("type")

	// get quote id
	id := c.Param("id")

	println("got quote type: ", quoteType)

	// switch requet quotetype
	switch strings.ToLower(quoteType) {
	case "finance":
		financialQuote := model.FinancialQuotes{}

		// SELECT * FROM financial_quote WHERE id = 1;
		db.First(&financialQuote, id)

		if financialQuote.ID != 0 {
			// DELETE FROM financial_quote WHERE id = finanancial_quote.id
			db.Delete(&financialQuote)
			println("quote =>", financialQuote.Quote)

			c.JSON(200, gin.H{"success": quoteType + " quote with Id: " + id + " deleted"})
		} else {
			c.JSON(404, gin.H{"error": quoteType + " quote with Id: " + id + "not found"})
		}

	case "love":
		loveQuote := model.LoveQuotes{}

		// SELECT * FROM love_quotes WHERE id = 1;
		db.First(&loveQuote, id)

		if loveQuote.ID != 0 {
			db.Delete(&loveQuote)

			println("quote =>", loveQuote.Quote)

			c.JSON(200, gin.H{"success": quoteType + " quote with Id: " + id + " deleted"})
		} else {
			c.JSON(404, gin.H{"error": quoteType + " quote with Id: " + id + "not found"})
		}

	case "life":
		lifeQuote := model.LifeQuotes{}

		// SELECT * FROM love_quotes WHERE id = 1;
		db.First(&lifeQuote, id)

		if lifeQuote.ID != 0 {

			db.Delete(&lifeQuote)

			println("quote =>", lifeQuote.Quote)

			c.JSON(200, gin.H{"success": quoteType + " quote with Id: " + id + " deleted"})
		} else {
			c.JSON(404, gin.H{"error": quoteType + " quote with Id: " + id + "not found"})
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid quote Type"})
	}
}

func getDbCOnn() *gorm.DB {
	db := db.SetUpDB()
	return db
}

func notFoundError(err error, c *gin.Context) {
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error ": err.Error()})
	}
}
func checkError(err error, c *gin.Context) {
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
	}
}
