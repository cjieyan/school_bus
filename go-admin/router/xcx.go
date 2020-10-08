package router

import (
	"go-admin/apis/xcx"
	"go-admin/handler"

	"github.com/gin-gonic/gin"
)

// 无需认证的路由代码
func registerXcxRouter(v1 *gin.RouterGroup) {
	v1.POST("/login", xcx.Api{}.Login)
	v1.POST("/record", xcx.Api{}.Record)
	v1.POST("/faceinfo", xcx.Api{}.StudentFaceInfo)
	v1.POST("/sites", xcx.Api{}.Sites)

	//需要校验header的token
	authGroup := v1.Group("/auth")
	authGroup.Use(handler.XcxCheckToken)
	{
		authGroup.GET("/info", xcx.Api{}.Info)
	}
}
