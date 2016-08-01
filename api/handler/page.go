package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/go-eyas/blog-try/api/config"

)

func IndexPage(this *gin.Context) {
	conf := config.Get();
	this.HTML(http.StatusOK, "index.tpl", gin.H{
		"Mode": conf.Mode,
		"ClientMode": conf.ClientMode,
	})
}