package xcx

import (
	"encoding/json"
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
// 上车or下册刷脸打卡
func (a Api)Swipe(c *gin.Context){

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

	//年月日
	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa( year) + month + strconv.Itoa( day )

	studentIdStr := strconv.Itoa(data.StudentId)

	//获取学生信息
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

	//查询学生是否在车上
	//Todo... 废弃
	//key := tools.Keys{}.GetOn(ymd, teacher.Id)
	//exist := tools.RdbSISMembers(key, studentIdStr)


	now := int(time.Now().Unix())

	//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, teacher.Id)
	swipeData, err := redis.String(tools.RdbHGet(swipeAtKey, studentIdStr))
	fmt.Println(", err ...", err )
	fmt.Println("swipeData,  ...", swipeData,  )

	exist := true
	if redis.ErrNil == err{//redis无数据
		exist = false
	}else{
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)

		fmt.Println("swipeAtInfo....", swipeAtInfo, now - swipeAtInfo.Time)
		if nil == err{
			//距离上次上车刷脸成功大于5分钟
			if 0 == swipeAtInfo.Status  && (now - swipeAtInfo.Time) > 60 * 1{
				//将标记为下车状态
				swipeAtStruct := models.SwipeAt{
					Status: 1,
					Time: now,
				}
				jsonBytes, err := json.Marshal(swipeAtStruct)
				if err != nil {
					fmt.Println(err)
				}
				//记录学生刷脸时间
				tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
				tools.RdbSetKeyExp(swipeAtKey, 86400)
				app.OK(c, nil, "下车刷脸成功..")
			}else if 1 == swipeAtInfo.Status {
				app.OK(c, nil, "您已下车了..")
			}else{
				app.OK(c, nil, "您已刷脸上车了..")
			}
		}
	}

	//未上车
	if !exist {
		//记录用户上车信息
		swipeAtStruct := models.SwipeAt{
			Status: 0,
			Time: now,
		}
		jsonBytes, err := json.Marshal(swipeAtStruct)
		if err != nil {
			fmt.Println("记录学生刷脸时间...", err)
		}
		//记录学生刷脸时间
		tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
		tools.RdbSetKeyExp(swipeAtKey, 86400)

		app.OK(c, nil, "上车刷脸成功")
	}
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
	//清除当前行程的刷脸数据
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, teacher.Id)
	tools.RdbDel(swipeAtKey)
}