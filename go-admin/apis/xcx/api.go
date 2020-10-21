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
//多线路信息
//
func (a Api)Lines(c *gin.Context){

	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacherModel.PostId = 1
	teacher, err := teacherModel.Get()

	tools.HasError(err, "您尚未绑定线路", -1)

	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carData, err := carModel.Get()
	tools.HasError(err, "车辆信息不存在", -1)

	//获取线路
	lineModel := models.ScbLines{}
	lineModel.CarId = carData.Id
	linesData, err := lineModel.GetAll()

	tools.HasError(err, "线路不存在", -1)
	var linesRetData []models.ScbLines
	for _, line := range linesData{
		line.CarNos = carData.CarNo
		line.CarId = carData.Id
		siteModel := models.SchSites{}
		siteModel.LineId = line.Id
		startSite, err := siteModel.GetStart()
		if nil == err{
			line.StartSite = startSite
		}
		endSite, err := siteModel.GetEnd()
		if nil == err{
			line.EndSite = endSite
		}
		linesRetData = append(linesRetData, line)
	}
	app.OK(c, linesRetData, "")
}

// 单一线路信息
func (a Api) LineInfo(c *gin.Context) {

	objParams := models.LineInfoReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err {
		tools.HasError(err, "", -1)
	}

	userId := c.GetInt(models.UserId)

	lineModel := models.ScbLines{}
	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.LineId = objParams.LineId
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
	sitesData, err := siteModel.GetAll()

	tools.HasError(err, "尚未个给您分配站点", -1)
	//获取乘车的学生总数
	studentModel := models.ScbStudents{}
	studentModel.CarId = carData.Id
	studentCount, err := studentModel.GetCount()

	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa(year) + month + strconv.Itoa(day)
	key := tools.Keys{}.GetOn(ymd, teacher.Id)
	students, _ := redis.Values(tools.RdbSMembers(key))

	rsp := make(map[string]interface{})
	rsp["teacher"] = teacher                 //跟车员信息
	rsp["car"] = carData                     //车辆信息
	rsp["line"] = lineData                   //线路信息
	rsp["sites"] = sitesData                 //站点信息
	rsp["studentCount"] = studentCount       //所有學生
	rsp["studentGetOnCount"] = len(students) //已上車的學生
	app.OK(c, rsp, "")
}

// 上车or下册刷脸打卡
func (a Api) Swipe(c *gin.Context) {

	objParams := models.SwipeReq{}

	err := c.ShouldBindJSON(&objParams)
	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	ret := models.SwipeRsp{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carModel.LineId = objParams.LineId
	carData, err := carModel.Get()

	tools.HasError(err, "无法查询正常线路信息", 500)

	//年月日
	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa(year) + month + strconv.Itoa(day)

	studentIdStr := strconv.Itoa(objParams.StudentId)

	//获取学生信息
	student := models.ScbStudents{}
	student.Id = objParams.StudentId
	studentData, err := student.Get()
	tools.HasError(err, "学生信息不存在", 500)

	if 0 == studentData.CarId {
		tools.HasError(err, "您尚未分配车辆", 500)
	}

	if studentData.CarId != carData.Id {
		tools.HasError(err, "您搭错车了", 500)
	}
	//查询学生是否在车上
	now := int(time.Now().Unix())

	//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)
	swipeData, err := redis.String(tools.RdbHGet(swipeAtKey, studentIdStr))
	fmt.Println(", err ...", err)
	fmt.Println("swipeData,  ...", swipeData)

	exist := true
	if redis.ErrNil == err { //redis无数据
		exist = false
	} else {
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)

		if nil == err {
			//距离上次上车刷脸成功大于5分钟
			if 0 == swipeAtInfo.Status && (now - swipeAtInfo.Time) > 60*1 {
				//将标记为下车状态
				swipeAtStruct := models.SwipeAt{
					Status: 1,
					Time:   now,
				}
				jsonBytes, err := json.Marshal(swipeAtStruct)
				if err != nil {
					fmt.Println(err)
				}
				//记录学生刷脸时间
				tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
				tools.RdbSetKeyExp(swipeAtKey, 86400)
				ret.Status = 1

				//记录下车日志
				carRecordModel := models.ScbCarRecord{}
				carRecordModel.CarId = carData.Id
				carRecordModel.QqLongitude = objParams.QqLongitude
				carRecordModel.QqLatitude = objParams.QqLatitude
				carRecordModel.StudentId = objParams.StudentId
				carRecordModel.Prop = 2
				carRecordModel.Create()



				app.OK(c, ret, "下车刷脸成功..")
			} else if 1 == swipeAtInfo.Status {
				ret.Status = 1
				app.OK(c, ret, "您已下车了..")
			} else {
				ret.Status = 0
				app.OK(c, ret, "您已刷脸上车了..")
			}
		}
	}

	//未上车
	if !exist {
		//记录用户上车信息
		swipeAtStruct := models.SwipeAt{
			Status: 0,
			Time:   now,
		}
		jsonBytes, err := json.Marshal(swipeAtStruct)
		if err != nil {
			fmt.Println("记录学生刷脸时间...", err)
		}

		//记录上车日志
		carRecordModel := models.ScbCarRecord{}
		carRecordModel.CarId = carData.Id
		carRecordModel.QqLongitude = objParams.QqLongitude
		carRecordModel.QqLatitude = objParams.QqLatitude
		carRecordModel.StudentId = objParams.StudentId
		carRecordModel.Prop = 1
		carRecordModel.Create()

		//记录学生刷脸时间
		tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
		tools.RdbSetKeyExp(swipeAtKey, 86400)
		ret.Status = 0
		app.OK(c, ret, "上车刷脸成功")
	}
}

func (a Api) LineStart(c *gin.Context){
	userId := c.GetInt(models.UserId)
	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员", -1)
	}

	swipeNowKey := tools.Keys{}.SwipeNow()
	swipeRecordKey := tools.Keys{}.SwipeNowRecord(teacher.Id)
	now := int(time.Now().Unix())
	nowStr := strconv.Itoa(now)
	tools.RdbHSet(swipeNowKey, swipeRecordKey, nowStr)
	app.OK(c, nil, "操作成功")
}

func(a Api)LineCheck(c *gin.Context){

	userId := c.GetInt(models.UserId)
	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员", -1)
	}

	swipeNowKey := tools.Keys{}.SwipeNow()
	swipeRecordKey := tools.Keys{}.SwipeNowRecord(teacher.Id)

	lineStartAt, err := redis.Int( tools.RdbHGet(swipeNowKey, swipeRecordKey) )
	rsp := models.LineCheckRsp{}
	rsp.StartAt = lineStartAt
	//rsp.Line =
	app.OK(c, rsp, "操作成功")
}

// 线路结束
func (a Api) LineFinish(c *gin.Context) {

	objParams := models.LineFinishReq{}

	err := c.ShouldBindJSON(&objParams)
	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carModel.LineId = objParams.LineId
	_, err = carModel.Get()
	tools.HasError(err, "无法查询正常线路信息", 500)

	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa(year) + month + strconv.Itoa(day)
	//清除当前行程的刷脸数据
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)
	tools.RdbDel(swipeAtKey)

	//结束跟车员当前行程
	swipeNowKey := tools.Keys{}.SwipeNow()
	swipeNowRecordKey := tools.Keys{}.SwipeNowRecord(teacher.Id)

	tools.RdbHDel(swipeNowKey, swipeNowRecordKey)

	app.OK(c, nil, "操作成功")
}

//此线路的所有
func (a Api) LineStudents(c *gin.Context) {
	objParams := models.LineStudentsReq{}
	err := c.ShouldBindJSON(&objParams)

	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err := teacherModel.Get()
	if teacher.PostId != 1 {
		tools.HasError(err, "您不是跟车员, sch_teachers表的postId = 1", -1)
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carModel.LineId = objParams.LineId
	carData, err := carModel.Get()

	tools.HasError(err, "车辆信息不存在", -1)

	//获取搭乘此辆车的所有学生
	studentModel := models.ScbStudents{}
	studentModel.CarId = carData.Id
	studentsData, err := studentModel.GetAllByCarId()

	//年月日
	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	ymd := strconv.Itoa(year) + month + strconv.Itoa(day)

	//获取此线路打卡信息
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)
	studentsSwipe, err := tools.RdbHGetAll(swipeAtKey)
	swipesData := make(map[int]int)
	studentId := 0

	//读取并整理学生是否乘车 上次 下车  status信息 -1 未上车 0 上车 1 下车
	for idx, sSwipe := range studentsSwipe {
		studentId, err = strconv.Atoi(idx)
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(sSwipe), &swipeAtInfo)
		swipesData[studentId] = swipeAtInfo.Status
	}
	unGetOn := 0 //未上车数量
	getOn := 0   //已上车数量
	getOff := 0  //已下车数量
	allCount := len(studentsData) //此车辆内 学生数量
	//整理返回数据
	var studentsDataRet []models.ScbStudents
	for _, student := range studentsData {
		status, sErr := swipesData[student.Id]
		if false == sErr {
			student.SwipeStatus = -1
			unGetOn ++
		} else {
			student.SwipeStatus = status
			if 0 == student.SwipeStatus{
				getOn ++
			}else{
				getOff ++
			}
		}
		studentsDataRet = append(studentsDataRet, student)
	}
	ret := make(map[string]interface{})
	ret["unGetOn"] = unGetOn //未上车数量
	ret["getOn"] = getOn //已上车数量
	ret["getOff"] = getOff //已下车数量
	ret["allCount"] = allCount //此车辆内 学生数量
	ret["studentsDataRet"] = studentsDataRet //所有学生信息
	app.OK(c, ret, "操作成功")
}
