package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/xcx"
)

// 无需认证的路由代码
func registerXcxRouter(v1 *gin.RouterGroup) {
	v1.POST("/login", xcx.Api{}.Login)
	v1.GET("/info", xcx.Api{}.Info)
	v1.POST("record", xcx.Api{}.Record)
}