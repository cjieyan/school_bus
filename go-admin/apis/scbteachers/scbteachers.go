package scbteachers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
)

func GetScbTeachersList(c *gin.Context) {
	var data models.ScbTeachers
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Id, _ = tools.StringToInt(c.Request.FormValue("id"))
	data.Name = c.Request.FormValue("name")
	data.Phone = c.Request.FormValue("phone")
	data.ClassId, _ = tools.StringToInt(c.Request.FormValue("classId"))
	data.PostId, _ = tools.StringToInt(c.Request.FormValue("postId"))
	data.IsDeleted = c.Request.FormValue("isDeleted")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}
//获取教师信息
func GetScbTeachers(c *gin.Context) {
	var data models.ScbTeachers
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	scbpost := models.ScbPost{}
	posts, err := scbpost.GetList()

	postIds := make([]int, 0)
	intPostId := result.PostId
	postIds = append(postIds, intPostId)

	app.Custum(c, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"posts":   posts,
	})
}

func InsertScbTeachers(c *gin.Context) {
	var data models.ScbTeachers
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateScbTeachers(c *gin.Context) {
	var data models.ScbTeachers
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbTeachers(c *gin.Context) {
	var data models.ScbTeachers

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
