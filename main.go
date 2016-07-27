package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eyas/blog-try/api"
)

func main() {
  config := api.ParseConfig("./config.ini")
  gin.SetMode(config.Mode)
	app := gin.Default()
	api.Main(app)
	app.Run(config.Host + ":" + config.Port)
}
