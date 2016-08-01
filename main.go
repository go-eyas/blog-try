package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eyas/blog-try/api"
	"github.com/go-eyas/blog-try/api/config"
)

func main() {
	config := config.Get()
	gin.SetMode(config.Mode)
	app := gin.Default()

	/**
	static folder
	*/
	app.Static("/static", func() string {
		var staticFolder string
		if config.ClientMode == "dev" {
			staticFolder = "./client/dev"
		} else {
			staticFolder = "./client/prod"
		}
		return staticFolder;
	}())

	api.Main(app)
	app.Run(config.Host + ":" + config.Port)
}
