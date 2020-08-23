package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/useraccount"
)

// 无需认证的路由代码
func registerUserAccountRouter(v1 *gin.RouterGroup) {

	r := v1.Group("/useraccount")
	{
		r.GET("/:id", useraccount.GetUserAccount)
		r.POST("", useraccount.InsertUserAccount)
	}
}
