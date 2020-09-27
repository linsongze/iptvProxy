package web

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/tvm.m3u8", TvmProxyHandler)
	r.GET("/cetv.m3u8", cetvProxyHandler)
}
