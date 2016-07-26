package api

import "github.com/gin-gonic/gin"

/**
routes
 */
func routes (router *gin.Engine) {
	router.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
}

/**
static folder
 */
func static(app *gin.Engine) {
	app.Static("/static", "./client/prod")
}

/**
init api
 */
func Main(app *gin.Engine) {
	routes(app)
	static(app)
}
