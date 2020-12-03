package router

import (
	"go-admin/pkg/jwtauth"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// 路由示例
func InitExamplesRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	examplesNoCheckRoleRouter(r, authMiddleware)
	// 需要认证的路由
	examplesCheckRoleRouter(r, authMiddleware)
	// 小程序接口
	xcxCheckRouter(r)

	return r
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	// 空接口防止v1定义无使用报错
	v1.GET("/nilcheckrole", nil)

	// {{无需认证路由自动补充在此处请勿删除}}
	registerScbCarRecordRouter(v1)
	//registerOrdersRouter(v1)
	//registerUserAccountRouter(v1)
	// registerCarsRecordRouter(v1, authMiddleware)
	// registerCarsRouter(v1,authMiddleware)
}

// 需要认证的路由示例
func examplesCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	// 空接口防止v1定义无使用报错
	v1.GET("/checkrole", nil)

	// {{认证路由自动补充在此处请勿删除}}
	registerScbFollowRecordRouter(v1, authMiddleware)
	registerScbPostRouter(v1, authMiddleware)
	registerScbDeptRouter(v1, authMiddleware)
	registerSchSitesRouter(v1, authMiddleware)
	registerScbLinesRouter(v1, authMiddleware)
	registerScbStudentsRouter(v1, authMiddleware)
	registerScbTeachersRouter(v1, authMiddleware)
	registerScbCarsRouter(v1, authMiddleware)
}

func xcxCheckRouter(r *gin.Engine) {
	v1 := r.Group("/xcx")
	registerXcxRouter(v1)
}
