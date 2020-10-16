package schsites

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

func GetSchSitesList(c *gin.Context) {
	var data models.SchSites
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
	data.LineId, _ = tools.StringToInt(c.Request.FormValue("lineId"))
	data.Name = c.Request.FormValue("name")
	data.Purpose = c.Request.FormValue("purpose")
	data.Sort = c.Request.FormValue("sort")
	data.Prop, _ = tools.StringToInt(c.Request.FormValue("prop"))
	data.ArriveAt = c.Request.FormValue("arriveAt")
	data.Remark = c.Request.FormValue("remark")

	data.Longitude, _ = strconv.ParseFloat(c.Request.FormValue("longitude"), 32)
	data.Latitude, _ = strconv.ParseFloat(c.Request.FormValue("latitude"), 32)
	data.Address = c.Request.FormValue("Address")
	data.Picture = c.Request.FormValue("picture")

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	var sitesData []models.SchSites
	for _, site := range result{
		lineModel := models.ScbLines{}
		lineModel.Id = site.LineId
		lineData, err := lineModel.Get()
		if err == nil{
			site.LineName = lineData.Name
		}
		if site.Prop == 2{
			site.PropName = "下车"
		}else{
			site.PropName = "上车"
		}
		sitesData = append(sitesData, site)
	}

	app.PageOK(c, sitesData, count, pageIndex, pageSize, "")
}

func GetSchSites(c *gin.Context) {
	var data models.SchSites
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertSchSites(c *gin.Context) {
	var data models.SchSites
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)

	qqLng, qqLat, qqErr := data.MapBd2qq(data.Longitude, data.Latitude)
	if nil == qqErr{
		data.QqLat = qqLat
		data.QqLng = qqLng
	}

	result, err := data.Create()

	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

func UpdateSchSites(c *gin.Context) {
	var data models.SchSites
	err := c.MustBindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)

	qqLng, qqLat, qqErr := data.MapBd2qq(data.Longitude, data.Latitude)
	if nil == qqErr{
		data.QqLat = qqLat
		data.QqLng = qqLng
	}
	fmt.Println("Latitude.....", data.Latitude, data.Longitude, err)

	result, err := data.Update(data.Id)
	tools.HasError(err, "", -1)

	app.OK(c, result, "")
}

func DeleteSchSites(c *gin.Context) {
	var data models.SchSites

	IDS := tools.IdsStrToIdsIntGroup("id", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, msg.DeletedFail, 500)
	app.OK(c, nil, msg.DeletedSuccess)
}