package scblines

import (
	"fmt"
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

	var linesData []models.ScbLines
	for _, line := range result{
		if "" != line.CarIds{
			carIdsInt := []int{}
			carIdsArr := strings.Split(line.CarIds, ",")
			for _, val := range carIdsArr{
				carId, cErr := strconv.Atoi(val)
				if nil == cErr && carId > 0 {
					carIdsInt = append(carIdsInt, carId)
				}
			}
			if len(carIdsInt ) > 0 {
				carModel := models.ScbCars{}
				carsData, err := carModel.GetbyIds(carIdsInt)
				if nil == err{
					carNos := ""
					for _, car := range carsData{
						carNos = carNos + " " + car.CarNo
					}
					line.CarNos = carNos
				}
			}
		}
		linesData = append(linesData, line)
	}
	app.PageOK(c, linesData, count, pageIndex, pageSize, "")
}

func GetScbLines(c *gin.Context) {
	var data models.ScbLines
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}
//新增线路
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
//获取站点的所有车辆
func GetCars(c *gin.Context){
	var data models.ScbLines
	data.Id, _ = tools.StringToInt(c.Query("id"))

	lineData, err := data.Get()
	tools.HasError(err, "线路找不到", -1)

	var carIdsSelected []int
	carIds := strings.Split(lineData.CarIds, ",")
	for i := 0; i < len(carIds); i++ {
		carId, _ := strconv.Atoi(carIds[i])
		carIdsSelected = append(carIdsSelected, carId)
	}
	carsData, _ := models.ScbCars{}.GetbyIds(carIdsSelected)
	app.OK(c, carsData,"")

}
//获取线路下的所有站点
func GetSites(c *gin.Context){
	var data models.ScbLines
	data.Id, _ = tools.StringToInt(c.Query("id"))
	lineData, err := data.Get()
	tools.HasError(err, "线路找不到", -1)

	var siteModel models.SchSites
	siteModel.LineId = lineData.Id
	fmt.Println(" lineData.Id...",  lineData.Id)
	sitesData, err := siteModel.GetAll()
	app.OK(c, sitesData,"")
}
