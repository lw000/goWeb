package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	{
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		err := viper.ReadInConfig()   // Find and read the config file
		if err != nil {               // Handle errors reading the config file
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
			} else {
				// Config file was found but another error was produced
			}
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		version := viper.GetInt32("version")
		log.Println(version)
	}

	engine := gin.Default()

	engine.Static("/res", "./res")
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/button", func(c *gin.Context) {
		c.HTML(http.StatusOK, "button.html", gin.H{"Title": "Button风格"})
	})

	engine.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{"Title": "Form风格"})
	})

	engine.POST("/form", func(c *gin.Context) {
		log.Println(c)
	})

	_ = engine.Run(":8888")
}
