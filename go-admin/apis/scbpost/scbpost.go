package scbpost

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
	"strconv"
)

func GetScbPostList(c *gin.Context) {
	var data models.ScbPost
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.PostName = c.Request.FormValue("postName")
	data.PostCode = c.Request.FormValue("postCode")
	data.Status, _ = strconv.Atoi( c.Request.FormValue("status") )

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetScbPost(c *gin.Context) {
	var data models.ScbPost
	data.PostId, _ = tools.StringToInt(c.Param("postId"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertScbPost(c *gin.Context) {
	var data models.ScbPost
	err := c.ShouldBindJSON(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateScbPost(c *gin.Context) {
	var data models.ScbPost
	err := c.MustBindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.PostId)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbPost(c *gin.Context) {
	var data models.ScbPost
	data.UpdateBy = tools.GetUserIdStr(c)

	IDS := tools.IdsStrToIdsIntGroup("postId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}

func GetScbPostAll(c *gin.Context){

	var scbpost models.ScbPost
	posts, _ := scbpost.GetList()
	app.OK(c, posts, "")
}