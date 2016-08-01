package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-eyas/blog-try/api/handler"
	"github.com/go-eyas/blog-try/api/config"
)

/**
routes
*/
func routes(router *gin.Engine) {
	router.GET("/", handler.IndexPage)
}

/**
static
 */
func static(app *gin.Engine) {
	conf := config.Get()
	var staticFolder string
	if conf.ClientMode == "dev" {
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
	app.LoadHTMLGlob("api/views/*")
	routes(app)
	static(app)
}
