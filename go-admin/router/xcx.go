package router

import (
	"go-admin/apis/xcx"
	"go-admin/handler"

	"github.com/gin-gonic/gin"
)

// 无需认证的路由代码
func registerXcxRouter(v1 *gin.RouterGroup) {
	v1.POST("/login", xcx.Api{}.Login)//登录信息
	
	//需要校验header的token
	authGroup := v1.Group("/auth")
	authGroup.Use(handler.XcxCheckToken)
	{
		authGroup.POST("/faceinfo", xcx.Api{}.StudentFaceInfo)
		authGroup.POST("/sites", xcx.Api{}.Sites)
		authGroup.GET("/line-info", xcx.Api{}.LineInfo) //线路信息
		authGroup.POST("/swipe", xcx.Api{}.Swipe) //打卡上下车
		authGroup.POST("/line-finish", xcx.Api{}.LineFinish) //结束行程
		authGroup.GET("/line-students", xcx.Api{}.LineStudents) //获取一条线路的所有学生信息
	}
}
