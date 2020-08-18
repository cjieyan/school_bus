package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbcars"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbCarsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbcars").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:", scbcars.GetScbCars)
		r.POST("", scbcars.InsertScbCars)
		r.PUT("", scbcars.UpdateScbCars)
		r.DELETE("/:", scbcars.DeleteScbCars)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbcarsList", scbcars.GetScbCarsList)
	}

}
