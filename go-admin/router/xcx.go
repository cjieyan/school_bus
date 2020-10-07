package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/xcx"
	"go-admin/handler"
)

// 无需认证的路由代码
func registerXcxRouter(v1 *gin.RouterGroup) {
	v1.POST("/login", xcx.Api{}.Login)

	//需要校验header的token
	authGroup := v1.Group("/auth")
	authGroup.Use(handler.XcxCheckToken)
	{
		authGroup.GET("/info", xcx.Api{}.Info)
		authGroup.POST("/record", xcx.Api{}.Record)
		authGroup.POST("/faceinfo", xcx.Api{}.StudentFaceInfo)
	}
}