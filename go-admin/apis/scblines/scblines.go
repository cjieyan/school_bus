package scblines

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
	"strconv"
	"strings"
)

func GetScbLinesList(c *gin.Context) {
	var data models.ScbLines
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
	data.ArrivedAt = c.Request.FormValue("arrivedAt")
	data.ChangeExpiredAt = c.Request.FormValue("changeExpiredAt")
	data.CarIds = c.Request.FormValue("carIds")
	data.IsDelete, _ = tools.StringToInt(c.Request.FormValue("isDelete"))

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetScbLines(c *gin.Context) {
	var data models.ScbLines
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertScbLines(c *gin.Context) {
	var data models.ScbLines
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateScbLines(c *gin.Context) {
	var data models.ScbLines
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "数据解析失败", -1)
	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteScbLines(c *gin.Context) {
	var data models.ScbLines

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}
func GetScbLinesTreeCarsselect(c *gin.Context) {
	var data models.ScbLines
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)


	carIdsSelected := []int{}
	if data.Id > 0 {
		carIds := strings.Split(result.CarIds, ",")
		for i := 0; i < len(carIds); i++ {
			carId, _ := strconv.Atoi(carIds[i])
			carIdsSelected = append(carIdsSelected, carId)
		}
	}
	var carModel models.ScbCars
	cars, err := carModel.GetAll()
	app.Custum(c, gin.H{
		"code":        200,
		"cars":		   cars,
		"checkedKeys": carIdsSelected,
	})
}

func GetAllLines(c *gin.Context) {
	var data models.ScbLines
	data.DataScope = tools.GetUserIdStr(c)
	result, err := data.GetAll()
	tools.HasError(err, "", -1)
	app.OK(c, result,"")
}
