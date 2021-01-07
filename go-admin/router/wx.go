package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/wx"
	"go-admin/handler"
)

// 无需认证的路由代码
func registerWxRouter(v1 *gin.RouterGroup) {
	v1.GET("/index", wx.Api{}.Index)
	v1.GET("/authorize", wx.Api{}.Authorize)
	v1.GET("/ticket-token", wx.Api{}.TicketToken)

	authGroup := v1.Group("/auth")
	authGroup.Use(handler.WxCheckToken)
	{
		v1.POST("/bind", wx.Api{}.Bind)
		v1.POST("/student-detail", wx.Api{}.StudentDetail)
	}
}
