package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbteachers"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbTeachersRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbteachers").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", scbteachers.GetScbTeachers)
		r.POST("", scbteachers.InsertScbTeachers)
		r.PUT("", scbteachers.UpdateScbTeachers)
		r.DELETE("/:id", scbteachers.DeleteScbTeachers)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbteachersList", scbteachers.GetScbTeachersList)
	}

}
