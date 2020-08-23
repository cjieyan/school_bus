package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/akcarsrecord"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerCarsRecordRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/akcarsrecord").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", akcarsrecord.GetAkCarsRecord)
		r.POST("", akcarsrecord.InsertAkCarsRecord)
		r.PUT("", akcarsrecord.UpdateAkCarsRecord)
		r.DELETE("/:id", akcarsrecord.DeleteAkCarsRecord)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/akcarsrecordList", akcarsrecord.GetAkCarsRecordList)
	}

}
