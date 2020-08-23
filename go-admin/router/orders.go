package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/orders"
)

// 无需认证的路由代码
func registerOrdersRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/orders")
	{
		r.GET("/:id", orders.GetOrders)
		r.POST("", orders.InsertOrders)
	}
}
