package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/wx"
)

// 无需认证的路由代码
func registerWxRouter(v1 *gin.RouterGroup) {
	v1.GET("/index", wx.Api{}.Index)
	v1.GET("/authorize", wx.Api{}.Authorize)
}
