package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/scbstudents"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerScbStudentsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/scbstudents").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", scbstudents.GetScbStudents)
		r.POST("", scbstudents.InsertScbStudents)
		r.PUT("", scbstudents.UpdateScbStudents)
		r.DELETE("/:id", scbstudents.DeleteScbStudents)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scbstudentsList", scbstudents.GetScbStudentsList)
	}

}
