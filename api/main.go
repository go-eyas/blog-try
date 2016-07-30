package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eyas/blog-try/api/handler"
)

/**
routes
*/
func routes(router *gin.Engine) {
	router.GET("/", handler.IndexPage)
}

/**
static folder
*/
func static(app *gin.Engine) {
	var staticFolder string
	if gin.Mode() == "debug" {
		staticFolder = "./client/dev"
	} else {
		staticFolder = "./client/prod"
	}
	app.Static("/static", staticFolder)
}

/**
init api
*/
func Main(app *gin.Engine) {
	//gin.SetMode(gin.ReleaseMode)
	app.LoadHTMLGlob("api/views/*")
	routes(app)
	static(app)
}
