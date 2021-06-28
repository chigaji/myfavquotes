package api

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome home rony",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are home",
		})
	})
	r.Run(":4000")

}
