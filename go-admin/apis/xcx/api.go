package xcx

import (
	"fmt"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (a Api) Login(c *gin.Context) {
	objParams := models.XcxLoginReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err {
		tools.HasError(err, "", -1)
	}

	model := models.ScbTeachers{}
	model.Phone = objParams.Phone
	teacher, err := model.Get()
	tools.HasError(err, "账号或密码错误", -1)
	_, err = tools.CompareHashAndPassword(teacher.Password, objParams.Password)
	tools.HasError(err, "账号或密码错误.", -1)
	//idStr := fmt.Sprintf("%v", teacher.Id)
	idStr := strconv.Itoa(teacher.Id)

	token := tools.GenRandomString(36) + idStr
	key := tools.Keys{}.ApiToken(token)

	b, err := tools.RdbSet(key, idStr)
	tools.RdbSetKeyExp(key, 3600*2)
	rsp := models.XcxLoginRsp{}
	rsp.Token = token

	app.OK(c, rsp, "")
}

func (a Api) Info(c *gin.Context) {

	userId := c.GetInt(models.UserId)

	fmt.Println("userId````...", userId)
	teacherModel := models.ScbTeachers{}
	teacherModel.Id = userId
	info, err := teacherModel.Get()
	tools.HasError(err, "账号异常,请联系管理员", -1)

	carModel := models.ScbCars{}
	carModel.AttendantId = teacherModel.Id
	car, err := carModel.Get()
	tools.HasError(err, "尚未分配跟随车辆", -1)

	lineModel := models.ScbLines{}
	lineModel.CarId = car.Id
	line, err := lineModel.Get()
	tools.HasError(err, "该车辆尚未分配线路", -1)

	siteModel := models.SchSites{}
	siteModel.LineId = line.Id
	sites, err := siteModel.GetAll()
	tools.HasError(err, "该车辆尚未分配站点", -1)

	rsp := make(map[string]interface{})
	rsp["info"] = info   //跟车员信息
	rsp["car"] = car     //车辆信息
	rsp["line"] = line   //线路信息
	rsp["sites"] = sites //站点信息
	app.OK(c, rsp, "")

}

//上下车日志记录
func (a Api) Record(c *gin.Context) {
	var data models.ScbCarRecord
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	fmt.Println(result)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

//人脸token获取用户信息
func (a Api) StudentFaceInfo(c *gin.Context) {
	var data models.ScbStudents
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	info, err := data.Get()
	app.OK(c, info, "")
}

func (a Api) Sites(c *gin.Context) {
	sitemodel := models.SchSites{}
	info, _, err := sitemodel.GetPage(100, 0)
	tools.HasError(err, "账号异常,请联系管理员", -1)
	app.OK(c, info, "")
}
