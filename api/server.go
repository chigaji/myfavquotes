package api

import (
	"github.com/chigaji/myfavquotes/db"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/", handleHome)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are home",
		})
	})
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
