package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbdept"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbdept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:deptId", scbdept.GetDept)
		r.POST("", scbdept.InsertDept)
		r.PUT("", scbdept.UpdateDept)
		r.DELETE("/:deptId", scbdept.DeleteDept)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbdeptList", scbdept.GetDeptList)
		l.GET("/scbdeptTree", scbdept.GetDeptTree)
	}

}
