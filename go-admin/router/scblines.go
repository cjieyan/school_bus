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

		//r.GET("/list", scblines.GetScbLinesList)
		//r.GET("/treeselect/:id", scblines.GetScbLinesTreeCarsselect)
		//r.GET("/all-lines", scblines.GetAllLines)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/scblinesList", scblines.GetScbLinesList)
		l.GET("/carsTreeselect/:id", scblines.GetScbLinesTreeCarsselect)
		//获取所有线路
		l.GET("/getAllLines", scblines.GetAllLines)
		//获取线路的所有车辆
		l.GET("/getLines/cars", scblines.GetCars)
		//获取线路的所有站点
		l.GET("/getLines/sites", scblines.GetSites)
	}

}
