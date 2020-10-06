package scbcars

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
	"strconv"
)

func GetScbCarsList(c *gin.Context) {
	var data models.ScbCars
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	data.CarNumber = c.Request.FormValue("carNumber")
	data.CarNo = c.Request.FormValue("carNo")
	data.Driver = c.Request.FormValue("driver")
	data.Phone = c.Request.FormValue("phone")
	data.AttendantId, _ = strconv.Atoi(c.Request.FormValue("attendantId"))

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetScbCars(c *gin.Context) {
	var data models.ScbCars
	var err error
	data.Id, err = tools.StringToInt(c.Param("id"))
	fmt.Println("err....", err)
	tools.HasError(err, "", 500)
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertScbCars(c *gin.Context) {
	var data models.ScbCars
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateScbCars(c *gin.Context) {
	var data models.ScbCars
	err := c.MustBindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbCars(c *gin.Context) {
	var data models.ScbCars

	IDS := tools.IdsStrToIdsIntGroup("", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
//获取所有车辆
func GetScbCarsAll(c *gin.Context) {
	var data models.ScbCars
	var err error

	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.GetAll()

	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}
