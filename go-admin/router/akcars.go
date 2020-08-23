package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/akcars"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerCarsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/akcars").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", akcars.GetAkCars)
		r.POST("", akcars.InsertAkCars)
		r.PUT("", akcars.UpdateAkCars)
		r.DELETE("/:id", akcars.DeleteAkCars)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/akcarsList", akcars.GetAkCarsList)
	}

}
