package router

import (
	"go-admin/apis/xcx"
	"go-admin/handler"

	"github.com/gin-gonic/gin"
)

// 无需认证的路由代码
func registerXcxRouter(v1 *gin.RouterGroup) {
	v1.POST("/login", xcx.Api{}.Login) //登录信息

	//需要校验header的token
	authGroup := v1.Group("/auth")
	authGroup.Use(handler.XcxCheckToken)
	{
		//authGroup.POST("/faceinfo", xcx.Api{}.StudentFaceInfo)
		authGroup.POST("/sites", xcx.Api{}.Sites)
		authGroup.POST("/lines", xcx.Api{}.Lines)				//线路列表
		authGroup.POST("/cars", xcx.Api{}.Cars)				//线路列表
		authGroup.POST("/line-info", xcx.Api{}.LineInfo)         //线路信息
		authGroup.POST("/swipe", xcx.Api{}.Swipe)                //打卡上下车
		authGroup.POST("/line-start", xcx.Api{}.LineStart)     //结束行程
		authGroup.POST("/line-finish", xcx.Api{}.LineFinish)     //结束行程
		authGroup.POST("/line-students", xcx.Api{}.LineStudents) //获取一条线路的所有学生信息
		authGroup.POST("/follow-record", xcx.Api{}.FollowRecord) //跟车记录
		authGroup.POST("/face-swipe", xcx.Api{}.FaceSwipe)       //人脸打卡
		authGroup.POST("/student-info", xcx.Api{}.StudentInfo)   //学生信息
	}
}
