package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(this *gin.Context) {
	this.HTML(http.StatusOK, "index.tpl", gin.H{
		"Mode": gin.Mode(),
	})
}