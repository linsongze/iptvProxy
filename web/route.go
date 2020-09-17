package web

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/tvm.m3u8", TvmProxyHandler)
}
