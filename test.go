package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./templates/assets")
	//router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("templates/test.html")
	router.GET("/blog", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		c.HTML(http.StatusOK, "test.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
