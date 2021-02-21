package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Static("/res", "./res")
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/button", func(c *gin.Context) {
		c.HTML(http.StatusOK, "button.html", gin.H{"Title": "Button"})
	})

	engine.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{"Title": "Form"})
	})

	engine.POST("/form", func(c *gin.Context) {
		log.Println(c)
	})

	_ = engine.Run(":8888")
}
