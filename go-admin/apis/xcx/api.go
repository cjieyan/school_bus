package xcx

import (
	"encoding/json"
	"errors"
	"fmt"
	tools2 "go-admin/apis/tools"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/config"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

//车辆列表
func (a Api) Cars(c *gin.Context) {

	userId := c.GetInt(models.UserId)

	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacherModel.PostId = 1
	teacher, err := teacherModel.Get()

	tools.HasError(err, "暂无权限", -1)

	//获取车辆
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carsData, err := carModel.GetAllByAttendantId()
	tools.HasError(err, "车辆信息不存在", -1)

	var linesRetData []models.ScbLines
	for _, car := range carsData {

		lineModel := models.ScbLines{}
		lineModel.CarId = car.Id
		line, err := lineModel.Get()
		if nil != err {
			continue
		}

		line.CarNos = car.CarNo
		line.CarId = car.Id
		siteModel := models.SchSites{}
		siteModel.LineId = line.Id
		startSite, err := siteModel.GetStart()
		if nil == err {
			line.StartSite = startSite
		}
		endSite, err := siteModel.GetEnd()
		if nil == err {
			line.EndSite = endSite
		}
		linesRetData = append(linesRetData, line)
	}
	app.OK(c, linesRetData, "")
}

//多线路信息
//
func (a Api) Lines(c *gin.Context) {

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
	for _, line := range linesData {
		line.CarNos = carData.CarNo
		line.CarId = carData.Id
		siteModel := models.SchSites{}
		siteModel.LineId = line.Id
		startSite, err := siteModel.GetStart()
		if nil == err {
			line.StartSite = startSite
		}
		endSite, err := siteModel.GetEnd()
		if nil == err {
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

	teacher, car, line, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
	tools.HasError(err, "", -1)

	//获取所有站点
	siteModel := models.SchSites{}
	siteModel.LineId = objParams.LineId
	sitesData, err := siteModel.GetAll()

	tools.HasError(err, "尚未个给您分配站点", -1)

	//获取乘车的学生总数
	studentModel := models.ScbStudents{}
	studentModel.CarId = car.Id
	studentCount, err := studentModel.GetCount()

	ymd := tools.Ymd()
	key := tools.Keys{}.GetOn(ymd, teacher.Id)
	students, _ := redis.Values(tools.RdbSMembers(key))

	followRecordModel := models.ScbFollowRecord{}
	followRecordModel.Ymd, err = strconv.Atoi(ymd)
	followRecordModel.LineId = line.Id
	followRecordModel.AttendantId = userId
	followRecordData, err := followRecordModel.Get()
	isFinished := -1
	if nil != err && followRecordData.IsFinished == 1 {
		isFinished = 1
	} else if len(students) > 0 {
		isFinished = 0
	}

	rsp := make(map[string]interface{})
	rsp["teacher"] = teacher                 //跟车员信息
	rsp["car"] = car                         //车辆信息
	rsp["line"] = line                       //线路信息
	rsp["sites"] = sitesData                 //站点信息
	rsp["studentCount"] = studentCount       //所有學生
	rsp["studentGetOnCount"] = len(students) //已上車的學生
	rsp["isFinished"] = isFinished           // -1行程尚未开始 0未结束 1 已结束
	app.OK(c, rsp, "")
}

//
//func (a Api) LineStart(c *gin.Context) {
//
//	objParams := models.LineStartReq{}
//	err := c.ShouldBindJSON(&objParams)
//	if nil != err {
//		tools.HasError(err, "", -1)
//	}
//
//	userId := c.GetInt(models.UserId)
//
//	//获取跟车员信息
//	teacher, car, _, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
//	tools.HasError(err, "", -1)
//
//	followRecordModel := models.ScbFollowRecord{}
//	followRecordModel.CarId = car.Id
//	followRecordModel.LineId = objParams.LineId
//	followRecordModel.AttendantId = teacher.Id
//	followRecordModel.Ymd, _ = strconv.Atoi(tools.Ymd())
//
//	followRecordData, err := followRecordModel.Get()
//	ret := models.LineFinishRsp{}
//	msg := "行程已经开始"
//	if gorm.IsRecordNotFoundError(err) {
//		followRecordModel.IsFinished = 0
//		followRecordData, err := followRecordModel.Create()
//		fmt.Println("followRecordData, err...", followRecordData, err)
//		ret.IsFinished = 1
//		msg = "操作成功"
//	}else if nil != err {
//		tools.HasError(err, "", -1)
//	} else if 1 == followRecordData.IsFinished {
//		msg = "行程已结束"
//	}
//	app.OK(c, nil, msg)
//}

func (a Api) LineCheck(c *gin.Context) {

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

	lineStartAt, err := redis.Int(tools.RdbHGet(swipeNowKey, swipeRecordKey))
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

	//获取跟车员信息
	teacher, car, _, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
	tools.HasError(err, "", -1)

	ymd := tools.Ymd()
	//清除当前行程的刷脸数据
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)

	//打卡总数
	//count, err := tools.RdbHlen(swipeAtKey)

	all, err := tools.RdbHGetAll(swipeAtKey)
	fmt.Println("all err ...", all, err)
	getOff := 0
	getOn := 0
	for _, sSwipe := range all {
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(sSwipe), &swipeAtInfo)
		getOn++
		if 1 == swipeAtInfo.Status {
			getOff++
		}
	}

	if getOn <= 0 && objParams.Force == 0 {
		//app.Error(c, 200, errors.New("线路结束失败"), "线路结束失败, 尚未有人上车")
		tools.HasError(err, "线路结束失败, 尚未有人上车", -1)
		data := models.LineFinishRsp{}
		data.GetOn = getOn
		app.Custum(c, gin.H{
			"code": -1,
			"data": data,
		})
		return
	}

	tools.RdbDel(swipeAtKey)

	studentModel := models.ScbStudents{}
	studentModel.CarId = car.Id
	allCount, err := studentModel.GetCount()

	followRecordModel := models.ScbFollowRecord{}
	followRecordModel.CarId = car.Id
	followRecordModel.LineId = car.LineId
	followRecordModel.AttendantId = teacher.Id
	followRecordModel.Ymd, _ = strconv.Atoi(tools.Ymd())

	followRecord, err := followRecordModel.Get()
	ret := models.LineFinishRsp{}
	msg := "操作失败"

	if gorm.IsRecordNotFoundError(err) {

		followRecordModel.GetOn = getOn
		followRecordModel.GetOff = getOff
		followRecordModel.AllCount = allCount
		followRecordModel.UnGetOn = allCount - getOn
		followRecordModel.IsFinished = 0
		followRecordData, err := followRecordModel.Create()
		fmt.Println("followRecordData, err...", followRecordData, err)
		ret.IsFinished = 1
		msg = "行程结束成功"
	} else if nil != err {
		tools.HasError(err, "", -1)
	} else {
		followRecordModel.IsFinished = 1
		followRecordModel.Update(followRecord.Id)

		ret.IsFinished = 1
		msg = "行程已结束"
	}
	ret.GetOn = getOn
	app.OK(c, nil, msg)
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
	teacher, car, _, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
	tools.HasError(err, "", -1)

	//获取搭乘此辆车的所有学生
	studentModel := models.ScbStudents{}
	studentModel.CarId = car.Id
	studentsData, err := studentModel.GetAllByCarId()

	//年月日
	ymd := tools.Ymd()

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
	unGetOn := 0                  //未上车数量
	getOn := 0                    //已上车数量
	getOff := 0                   //已下车数量
	allCount := len(studentsData) //此车辆内 学生数量
	//整理返回数据
	var studentsDataRet []models.ScbStudents
	for _, student := range studentsData {
		status, sErr := swipesData[student.Id]
		if false == sErr {
			student.SwipeStatus = -1
			unGetOn++
		} else {
			student.SwipeStatus = status
			if 0 == student.SwipeStatus {
				getOn++
			} else {
				getOff++
			}
		}
		student.HeadImgSmall = config.ApplicationConfig.ImageUrl + strings.Replace(student.HeadImg, ".", "_small.", 1)
		student.HeadImg = config.ApplicationConfig.ImageUrl + student.HeadImg
		studentsDataRet = append(studentsDataRet, student)
	}
	ret := make(map[string]interface{})
	ret["unGetOn"] = unGetOn                 //未上车数量
	ret["getOn"] = getOn                     //已上车数量
	ret["getOff"] = getOff                   //已下车数量
	ret["allCount"] = allCount               //此车辆内 学生数量
	ret["studentsDataRet"] = studentsDataRet //所有学生信息
	app.OK(c, ret, "操作成功")
}

func (a Api) FollowRecord(c *gin.Context) {
	objParams := models.FollowRecordReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err {
		tools.HasError(err, "", -1)
	}
	if objParams.PageSize <= 0 {
		objParams.PageSize = 10
	}
	if objParams.PageSize > 100 {
		objParams.PageSize = 100
	}

	userId := c.GetInt(models.UserId)

	ymd := tools.Ymd()
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, 1)
	tmpData, err := tools.RdbHGetAll(swipeAtKey)
	fmt.Println("swipeData...", tmpData)

	followModel := models.ScbFollowRecord{}
	followModel.AttendantId = userId

	result, count, err := followModel.GetPage(objParams.PageSize, objParams.PageIndex)
	var followRecordsData []models.ScbFollowRecord
	for _, follow := range result {
		carModel := models.ScbCars{}
		carModel.Id = follow.CarId
		carData, err := carModel.Get()
		if nil == err {
			follow.Car = carData
		}
		lineModel := models.ScbLines{}
		lineModel.Id = follow.LineId
		lineData, err := lineModel.Get()
		if nil == err {
			follow.Line = lineData
		}
		follow.TimeString = follow.CreatedAt.Format("2006-01-02 15:04:05")
		followRecordsData = append(followRecordsData, follow)
	}

	tools.HasError(err, "暂时没有跟车记录", -1)

	ret := make(map[string]interface{})
	ret["count"] = count                                                       //未上车数量
	ret["result"] = followRecordsData                                          //已上车数量
	ret["page_index"] = objParams.PageIndex                                    //已下车数量
	ret["page_size"] = objParams.PageSize                                      //每页数量
	ret["totalpage"] = math.Ceil(float64(count) / float64(objParams.PageSize)) //总页数
	app.OK(c, ret, "操作成功")
}

// 人脸打卡 支持多脸
func (a Api) FaceSwipe(c *gin.Context) {
	objParams := models.FaceSwipeReq{}

	err := c.ShouldBindJSON(&objParams)
	if nil != err {
		tools.HasError(err, "", -1)
	}

	userId := c.GetInt(models.UserId)

	teacher, car, _, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
	tools.HasError(err, "", -1)

	//年月日
	ymd := tools.Ymd()

	// 获取百度接口识别到的人脸
	api := tools2.BdApi{}
	userIds := api.MutilSearch(objParams.Image)
	if len(userIds) <= 0 {
		tools.HasError(err, "扫码失败,未发识别到人脸", 500)
	}

	studentModel := models.ScbStudents{}
	//获取到识别的学生列表
	studentsData, err := studentModel.GetByFaceTokens(objParams.LineId, car.Id, userIds)
	tools.HasError(err, "扫码失败.您上错车了", 500)

	//创建跟车记录
	followRecordModel := models.ScbFollowRecord{}
	followRecordModel.CarId = car.Id
	followRecordModel.LineId = car.LineId
	followRecordModel.AttendantId = teacher.Id
	followRecordModel.Ymd, _ = strconv.Atoi(tools.Ymd())
	_, err = followRecordModel.Get()
	if gorm.IsRecordNotFoundError(err) {
		followRecordModel.IsFinished = 0
		followRecordData, err := followRecordModel.Create()
		fmt.Println("followRecordData, err...", followRecordData, err)
	}

	var studentsStatus []models.FaceSwipeRspStudentStatus
	for _, studentData := range studentsData {
		studentIdStr := strconv.Itoa(studentData.Id)

		//查询学生是否在车上
		now := int(time.Now().Unix())

		//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
		swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)
		swipeData, err := redis.String(tools.RdbHGet(swipeAtKey, studentIdStr))

		var sStatus models.FaceSwipeRspStudentStatus
		exist := true
		if redis.ErrNil == err { //redis无数据
			exist = false
		} else {
			var swipeAtInfo models.SwipeAt
			fmt.Println("swipeData....", swipeData, err)
			err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)
			fmt.Println("swipeAtInfo SwipeAt err....", err)
			if nil == err {
				sStatus.StudentId = studentData.Id
				//距离上次上车刷脸成功大于5分钟
				if 0 == swipeAtInfo.Status && (now-swipeAtInfo.Time) > 60*1 {
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
					sStatus.Status = 1 //标记为已下车

					//记录下车日志
					carRecordModel := models.ScbCarRecord{}
					carRecordModel.CarId = car.Id
					carRecordModel.QqLongitude = objParams.QqLongitude
					carRecordModel.QqLatitude = objParams.QqLatitude
					carRecordModel.StudentId = studentData.Id
					carRecordModel.Prop = 2
					_, err = carRecordModel.Create()
					fmt.Println("carRecordModel.Create err........", err)
				} else if 1 == swipeAtInfo.Status {
					sStatus.Status = 1 //标记为已下车
				} else {
					sStatus.Status = 0 //标记为已上车
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
				fmt.Println("记录学生刷脸时间.....swipeAtStruct...", err)
			}

			//记录上车日志
			carRecordModel := models.ScbCarRecord{}
			carRecordModel.CarId = car.Id
			carRecordModel.QqLongitude = objParams.QqLongitude
			carRecordModel.QqLatitude = objParams.QqLatitude
			carRecordModel.StudentId = studentData.Id
			carRecordModel.Prop = 1
			_, err = carRecordModel.Create()
			fmt.Println("FaceSwipe... err carRecordModel.Create ...... ", err)

			//记录学生刷脸时间
			tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
			tools.RdbSetKeyExp(swipeAtKey, 86400)
			sStatus.Status = 0
		}
		studentsStatus = append(studentsStatus, sStatus)
	}

	//获取此车辆的所有学生

	getOff, getOn := getSwipeCount(objParams.LineId)
	isFinished := false

	studentsModel := models.ScbStudents{}
	studentsModel.CarId = car.Id
	allStudentsCunt, err := studentsModel.GetCount()
	tools.HasError(err, "此车辆未绑定任一学生", -1)
	if len(studentsStatus) > 0 && 0 == studentsStatus[0].Status {
		if getOff > 0 {
			if allStudentsCunt <= getOff {
				isFinished = true
			}
		} else if getOn > 0 {
			if allStudentsCunt <= getOn {
				isFinished = true
			}
		}
	}

	var rsp models.FaceSwipeRsp
	rsp.StudentStatus = studentsStatus
	rsp.Num = len(studentsStatus)
	rsp.IsFinished = isFinished
	if len(studentsStatus) > 0 {
		app.OK(c, rsp, "")
	} else {
		tools.HasError(errors.New("未识别到人脸"), "未识别到人脸", -1)
	}
}

// 上车or下车打卡 手动打卡
func (a Api) Swipe(c *gin.Context) {

	ret := models.SwipeRsp{}
	objParams := models.SwipeReq{}

	err := c.ShouldBindJSON(&objParams)
	userId := c.GetInt(models.UserId)

	teacher, car, _, err := a.teacherCarLine(userId, objParams.CarId, objParams.LineId)
	tools.HasError(err, "", -1)

	//年月日
	ymd := tools.Ymd()

	studentIdStr := strconv.Itoa(objParams.StudentId)

	//获取学生信息
	student := models.ScbStudents{}
	student.Id = objParams.StudentId
	studentData, err := student.Get()
	tools.HasError(err, "学生信息不存在", 500)

	if 0 == studentData.CarId {
		tools.HasError(err, "您尚未分配车辆", 500)
	}

	if studentData.CarId != car.Id {
		tools.HasError(err, "您搭错车了", 500)
	}
	//查询学生是否在车上
	now := int(time.Now().Unix())

	//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)

	fmt.Println("carData.Id.......", car.Id, car.LineId, teacher.Id)
	followRecordModel := models.ScbFollowRecord{}
	followRecordModel.CarId = car.Id
	followRecordModel.LineId = car.LineId
	followRecordModel.AttendantId = teacher.Id
	followRecordModel.Ymd, _ = strconv.Atoi(tools.Ymd())

	_, err = followRecordModel.Get()
	if gorm.IsRecordNotFoundError(err) {
		followRecordModel.IsFinished = 0
		followRecordData, err := followRecordModel.Create()
		fmt.Println("followRecordData, err...", followRecordData, err)
	}

	swipeData, err := redis.String(tools.RdbHGet(swipeAtKey, studentIdStr))
	fmt.Println("err...", err, redis.ErrNil)
	exist := true
	if redis.ErrNil == err { //redis无数据
		exist = false
	} else {
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)
		fmt.Println("swipeAtInfo SwipeAt err....", err)
		if nil == err {
			//距离上次上车刷脸成功大于5分钟
			if 0 == swipeAtInfo.Status && (now-swipeAtInfo.Time) > 60*1 {
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
				carRecordModel.CarId = car.Id
				carRecordModel.QqLongitude = objParams.QqLongitude
				carRecordModel.QqLatitude = objParams.QqLatitude
				carRecordModel.StudentId = objParams.StudentId
				carRecordModel.Prop = 2
				carRecordModel.Create()

				app.OK(c, ret, "下车打卡成功..")
			} else if 1 == swipeAtInfo.Status {
				ret.Status = 1
				app.OK(c, ret, "您已下车了..")
			} else {
				ret.Status = 0
				app.OK(c, ret, "您已打卡上车了..")
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
		carRecordModel.CarId = car.Id
		carRecordModel.QqLongitude = objParams.QqLongitude
		carRecordModel.QqLatitude = objParams.QqLatitude
		carRecordModel.StudentId = objParams.StudentId
		carRecordModel.Prop = 1
		carRecordModel.Type = 1
		_, err = carRecordModel.Create()
		fmt.Println("swipe carRecordModel.Create...", err)

		//记录学生刷脸时间
		tools.RdbHSet(swipeAtKey, studentIdStr, string(jsonBytes))
		tools.RdbSetKeyExp(swipeAtKey, 86400)
		ret.Status = 0
		app.OK(c, ret, "上车打卡成功")
	}
}

func getSwipeCount(lineId int) (getOff, getOn int) {
	ymd := tools.Ymd()
	//清除当前行程的刷脸数据
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, lineId)

	//打卡总数
	//count, err := tools.RdbHlen(swipeAtKey)

	all, err := tools.RdbHGetAll(swipeAtKey)
	fmt.Println("all err ...", all, err)

	for _, sSwipe := range all {
		var swipeAtInfo models.SwipeAt
		err = json.Unmarshal([]byte(sSwipe), &swipeAtInfo)
		getOn++
		if 1 == swipeAtInfo.Status {
			getOff++
		}
	}
	return
}
func (a Api) StudentInfo(c *gin.Context) {
	// 人脸token获取用户信息

	var objParams models.StudentInfoReq
	err := c.ShouldBindJSON(&objParams)
	tools.HasError(err, "", 500)

	var studentModel models.ScbStudents
	studentModel.Id = objParams.StudentId
	studentData, err := studentModel.Get()
	studentData.HeadImgSmall = config.ApplicationConfig.ImageUrl + strings.Replace(studentData.HeadImg, ".", "_small.", 1)
	studentData.HeadImg = config.ApplicationConfig.ImageUrl + studentData.HeadImg
	studentData.TimeString = studentData.CreatedAt.Format("2006-01-02 15:04:05")

	studentIdStr := strconv.Itoa(studentData.Id)
	//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
	ymd := tools.Ymd()
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, objParams.LineId)
	swipeData, rErr := redis.String(tools.RdbHGet(swipeAtKey, studentIdStr))

	var swipeAtInfo models.SwipeAt
	err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)
	fmt.Println("swipeAtInfo SwipeAt err =>.", err, "swipeAtInfo SwipeAt rErr. => ", rErr, "swipeAtInfo => ", swipeAtInfo)
	status := -1 //未上车
	if redis.ErrNil == rErr {
	} else if rErr == nil {
		if nil == err {
			if 1 == swipeAtInfo.Status {
				status = 1 //已下车
			} else {
				status = 0 //已上车
			}
		}
	}
	studentData.SwipeStatus = status
	app.OK(c, studentData, "")

}

//教师车辆线路
func (a Api) teacherCarLine(userId, carId, lineId int) (teacher models.ScbTeachers, car models.ScbCars, line models.ScbLines, err error) {
	lineModel := models.ScbLines{}
	teacherModel := models.ScbTeachers{}

	//获取跟车员信息
	teacherModel.Id = userId
	teacher, err = teacherModel.Get()
	if err != nil || teacher.PostId != 1 {
		//tools.HasError(err, "您不是跟车员", -1)
		err = errors.New("您不是跟车员")
		return
	}

	//获取车辆信息
	carModel := models.ScbCars{}
	carModel.AttendantId = teacher.Id
	carModel.Id = carId
	car, err = carModel.Get()

	if nil != err {
		err = errors.New("车辆信息不存在")
		return
	}

	//获取线路id 车辆id查找线路
	lineModel.Id = lineId
	lineModel.CarId = carId
	line, err = lineModel.Get()
	if nil != err {
		err = errors.New("尚未个给您分配路线")
		return
	}
	return
}
