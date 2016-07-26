package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eyas/blog/api"
)

func main() {
	app := gin.Default()
	api.Main(app)
	app.Run()
}
