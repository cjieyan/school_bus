package router

import (
	"go-admin/apis/xcx"
	"go-admin/handler"

	"github.com/gin-gonic/gin"
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
		authGroup.POST("/sites", xcx.Api{}.Sites)
		authGroup.GET("/line-info", xcx.Api{}.LineInfo)
	}
}
