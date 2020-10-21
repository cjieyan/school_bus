package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbfollowrecord"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbFollowRecordRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbfollowrecord").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", scbfollowrecord.GetScbFollowRecord)
		r.POST("", scbfollowrecord.InsertScbFollowRecord)
		r.PUT("", scbfollowrecord.UpdateScbFollowRecord)
		r.DELETE("/:id", scbfollowrecord.DeleteScbFollowRecord)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbfollowrecordList", scbfollowrecord.GetScbFollowRecordList)
	}

}
