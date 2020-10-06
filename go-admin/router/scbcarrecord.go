package router

import (
"github.com/gin-gonic/gin"
"go-admin/apis/scbcarrecord"
)

// 无需认证的路由代码
func registerScbCarRecordRouter(v1 *gin.RouterGroup) {

v1.GET("/scbcarrecordList",scbcarrecord.GetScbCarRecordList)

r := v1.Group("/scbcarrecord")
{
r.GET("/:id", scbcarrecord.GetScbCarRecord)
r.POST("", scbcarrecord.InsertScbCarRecord)
r.PUT("", scbcarrecord.UpdateScbCarRecord)
r.DELETE("/:id", scbcarrecord.DeleteScbCarRecord)
}
}
