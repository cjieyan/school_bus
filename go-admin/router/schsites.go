package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/schsites"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerSchSitesRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/schsites").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", schsites.GetSchSites)
		r.POST("", schsites.InsertSchSites)
		r.PUT("", schsites.UpdateSchSites)
		r.DELETE("/:id", schsites.DeleteSchSites)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/schsitesList", schsites.GetSchSitesList)
	}

}
