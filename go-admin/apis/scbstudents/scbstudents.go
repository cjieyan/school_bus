package scbstudents

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
)

func GetScbStudentsList(c *gin.Context) {
	var data models.ScbStudents
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
	data.Number = c.Request.FormValue("number")
	data.ClassId = c.Request.FormValue("classId")
	data.LineId, _ = tools.StringToInt(c.Request.FormValue("lineId"))
	data.SiteName = c.Request.FormValue("siteName")
	data.SiteId = c.Request.FormValue("siteId")
	data.CarId = c.Request.FormValue("carId")
	data.ParentPhone = c.Request.FormValue("parentPhone")
	data.Picture = c.Request.FormValue("picture")
	data.IsDeleted = c.Request.FormValue("isDeleted")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetScbStudents(c *gin.Context) {
	var data models.ScbStudents
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertScbStudents(c *gin.Context) {
	var data models.ScbStudents
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateScbStudents(c *gin.Context) {
	var data models.ScbStudents
	err := c.MustBindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbStudents(c *gin.Context) {
	var data models.ScbStudents

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
