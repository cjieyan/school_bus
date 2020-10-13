package xcx

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"strconv"
	"time"

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

	tools.RdbSet(key, idStr)
	tools.RdbSetKeyExp(key, 3600*2)
	rsp := models.XcxLoginRsp{}
	rsp.Token = token

	app.OK(c, rsp, "")
}

func (a Api) Info(c *gin.Context) {

	//userId := c.GetInt(models.UserId)
	//
	//fmt.Println("userId````...", userId)
	//teacherModel := models.ScbTeachers{}
	//teacherModel.Id = userId
	//info, err := teacherModel.Get()
	//tools.HasError(err, "账号异常,请联系管理员", -1)
	//
	//carModel := models.ScbCars{}
	//carModel.AttendantId = teacherModel.Id
	//car, err := carModel.Get()
	//tools.HasError(err, "尚未分配跟随车辆", -1)
	//
	//lineModel := models.ScbLines{}
	//lineModel.CarId = car.Id
	//line, err := lineModel.Get()
	//tools.HasError(err, "该车辆尚未分配线路", -1)
	//
	//siteModel := models.SchSites{}
	//siteModel.LineId = line.Id
	//sites, err := siteModel.GetAll()
	//tools.HasError(err, "该车辆尚未分配站点", -1)
	//
	//rsp := make(map[string]interface{})
	//rsp["info"] = info   //跟车员信息
	//rsp["car"] = car     //车辆信息
	//rsp["line"] = line   //线路信息
	//rsp["sites"] = sites //站点信息
	//app.OK(c, rsp, "")

}

// 上下车日志记录
func (a Api) Record(c *gin.Context) {
	var data models.ScbCarRecord
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	fmt.Println(result)
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

// 人脸token获取用户信息
func (a Api) StudentFaceInfo(c *gin.Context) {
	var data models.ScbStudents
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	info, err := data.Get()
	app.OK(c, info, "")
}
//站点信息
func (a Api) Sites(c *gin.Context) {
	sitemodel := models.SchSites{}
	info, _, err := sitemodel.GetPage(100, 0)
	tools.HasError(err, "账号异常,请联系管理员", -1)
	app.OK(c, info, "")
}
// 线路信息
func (a Api)LineInfo(c *gin.Context){
	userId := c.GetInt(models.UserId)

	lineModel := models.ScbLines{}
	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1{
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carData, err := carModel.Get()

	tools.HasError(err, "车辆信息不存在", -1)

	//获取线路
	lineModel.Id = carData.LineId
	lineData, err := lineModel.GetCarIds()

	tools.HasError(err, "尚未个给您分配路线", -1)

	//获取所有站点
	siteModel := models.SchSites{}
	siteModel.LineId = lineData.Id
	sitesData, err :=siteModel.GetAll()

	tools.HasError(err, "尚未个给您分配站点", -1)
	//获取乘车的学生总数
	studentModel := models.ScbStudents{}
	studentModel.CarId = carData.Id
	studentCount, err := studentModel.GetCount()



	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa( year) + month + strconv.Itoa( day )
	key := tools.Keys{}.GetOn(ymd, teacher.Id)
	students, _ := redis.Values(tools.RdbSMembers(key))

	rsp := make(map[string]interface{})
	rsp["teacher"] = teacher   //跟车员信息
	rsp["car"] = carData     //车辆信息
	rsp["line"] = lineData   //线路信息
	rsp["sites"] = sitesData //站点信息
	rsp["studentCount"] = studentCount //所有學生
	rsp["studentGetOnCount"] = len(students)//已上車的學生
	app.OK(c, rsp, "")
}
// 上车
func (a Api)GetOn(c *gin.Context){

	data := models.XcxGetOn{}

	err := c.ShouldBindJSON(&data)
	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1{
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carData, err := carModel.Get()

	tools.HasError(err, "", 500)

	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa( year) + month + strconv.Itoa( day )

	studentIdStr := strconv.Itoa(data.StudentId)

	student := models.ScbStudents{}
	student.Id = data.StudentId
	studentData, err := student.Get()
	tools.HasError(err, "学生信息不存在", 500)

	if 0 == studentData.CarId{
		tools.HasError(err, "您尚未分配车辆", 500)
	}

	if studentData.CarId != carData.Id{
		tools.HasError(err, "您搭错车了", 500)
	}

	//hour, err := strconv.Atoi(time.Now().Format("15"))

	if nil == err {
		//上车
		key := tools.Keys{}.GetOn(ymd, teacher.Id)
		exist := tools.RdbSISMembers(key, studentIdStr)
		//未上车
		if !exist {
			tools.RdbZAdd(key, studentIdStr)
			tools.RdbSetKeyExp(key, 86400)
			app.OK(c, nil, "上车刷脸成功")
		} else {
			app.OK(c, nil, "您已刷脸上车了")
		}
	}
	tools.HasError(err, "系统时间异常,请稍后重试", 500)
}

func (a Api)GetOff(c *gin.Context){
	data := models.XcxGetOn{}

	err := c.ShouldBindJSON(&data)
	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1{
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carData, err := carModel.Get()

	tools.HasError(err, "", 500)

	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa( year) + month + strconv.Itoa( day )
	//下车
	key := tools.Keys{}.GetOff(ymd, teacher.Id)

	studentIdStr := strconv.Itoa(data.StudentId)

	student := models.ScbStudents{}
	student.Id = data.StudentId
	studentData, err := student.Get()
	tools.HasError(err, "学生信息不存在", 500)

	if 0 == studentData.CarId{
		tools.HasError(err, "您尚未分配车辆", 500)
	}

	if studentData.CarId != carData.Id{
		tools.HasError(err, "您搭错车了", 500)
	}

	exist := tools.RdbSISMembers(key, studentIdStr)
	//未下车
	if ! exist {
		tools.RdbZAdd(key, studentIdStr)
		tools.RdbSetKeyExp(key, 86400)
		app.OK(c, nil, "下车刷脸成功")
	} else {
		app.OK(c, nil, "您已刷脸下车了")
	}

	tools.HasError(err, "系统时间异常,请稍后重试", 500)
	return
}
// 线路结束
func (a Api)LineFinish(c *gin.Context){

	data := models.XcxGetOn{}

	err := c.ShouldBindJSON(&data)
	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1{
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa( year) + month + strconv.Itoa( day )
	//下车
	keyOff := tools.Keys{}.GetOff(ymd, teacher.Id)
	//上车
	keyOn := tools.Keys{}.GetOn(ymd, teacher.Id)

	tools.RdbDel(keyOn)
	tools.RdbDel(keyOff)
}
// 线路开始
func (a Api) LineStart(){

}