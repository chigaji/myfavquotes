package api

import (
	"github.com/chigaji/myfavquotes/db"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/", handleHome)
	api.GET("/createQuote", createQuote)
	r.Run(":4000")

}
func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is an app that gets your favorite quotes depending on topic of your interest",
	})
}

func createQuote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is an app that gets your favorite quotes depending on topic of your interest",
	})
}
func RunDB() {
	db.SetUpDB()
}
