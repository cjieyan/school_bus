package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbpost"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbPostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbpost").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:postId", scbpost.GetScbPost)
		r.POST("", scbpost.InsertScbPost)
		r.PUT("", scbpost.UpdateScbPost)
		r.DELETE("/:postId", scbpost.DeleteScbPost)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbpostList", scbpost.GetScbPostList)
		l.GET("/scbpostAll", scbpost.GetScbPostAll)
	}

}
