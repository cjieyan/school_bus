package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scblines"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbLinesRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scblines").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", scblines.GetScbLines)
		r.POST("", scblines.InsertScbLines)
		r.PUT("", scblines.UpdateScbLines)
		r.DELETE("/:id", scblines.DeleteScbLines)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scblinesList", scblines.GetScbLinesList)
	}

}
