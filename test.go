package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	engine.Static("/", "./templates")
	//engine.LoadHTMLGlob("templates/*")
	//engine.LoadHTMLFiles("templates/index.html")
	//engine.GET("/", func(c *gin.Context) {
	//	if pusher := c.Writer.Pusher(); pusher != nil {
	//		// use pusher.Push() to do server push
	//		if err := pusher.Push("/assets/app.js", nil); err != nil {
	//			log.Printf("Failed to push: %v", err)
	//		}
	//	}
	//
	//	c.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})
	engine.Run(":8080")
}
