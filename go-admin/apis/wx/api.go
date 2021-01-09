package wx

import (
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/rand"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/config"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Api struct {

}


const (
	wxAppId           = "wx43a88770e5cfbdb9"                           // 填上自己的参数
	wxAppSecret       = "5f9abfd7f96e8cce52279512cd2b5cff"                       // 填上自己的参数
	oauth2RedirectURI = "https://wx.xcybus.com/wx/authorize" // 填上自己的参数
	oauth2Scope       = "snsapi_userinfo"                 // 填上自己的参数


	wxOriId         = "oriid"
	wxToken         = "xcybus"
	wxEncodedAESKey = "FF5UylipCjKWmSxFVVqQ1SmYDEFoe8Ml2KT1PgR6ELM"
)
var(

	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
)

func (_ Api)Index(c * gin.Context){
	state := string(rand.NewHex())
	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
	c.Redirect(http.StatusMovedPermanently,AuthCodeURL)

}
//生成ticket并且跳转
func (_ Api)Authorize(c *gin.Context){

	code := c.Query("code")
	state := c.Query("state")

	fmt.Printf("code: %+v\r\n state:%+v\n", code, state)
	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("token: %+v\r\n\n", token)

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("userinfo: %+v\r\n\n\n", userinfo)
	var doc models.SchUserAuth
	doc.Appid = wxAppId
	doc.Source = "wechat"
	doc.Openid = userinfo.OpenId

	userAuth, err := doc.Get()
	var userModel, user models.SchUser
	if gorm.IsRecordNotFoundError(err){
		userModel.Nickname = userinfo.Nickname
		user, err = userModel.Create()
		if nil != err{
			return
		}
		doc.UserId = user.Id
		userAuth ,err = doc.Create()
		if nil != err{
			return
		}
	}

	//生成ticket
	ticket := tools.GenRandomString(20) + tools.Uniqid()
	key := tools.Keys{}.Ticket(ticket)
	tools.RdbSetExp(key, strconv.Itoa(userAuth.UserId), 5 * 60)
	IndexURL := "https://h5.xcybus.com/?ticket=" + ticket
	c.Redirect(http.StatusMovedPermanently, IndexURL)
}

//ticket换token
func (_ Api)TicketToken(c *gin.Context){

	objParams := models.TicketTokenReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err{
		tools.HasError(err, "参数错误", -1)
	}
	ticketKey := tools.Keys{}.Ticket(objParams.Ticket)
	userIdStr, err := tools.RdbGet(ticketKey)
	if nil != err{
		tools.HasError(err, "临时票据已失效", -1)
	}
	userId, err := strconv.Atoi(userIdStr)
	if nil != err || userId <= 0{
		tools.HasError(err, "临时票据已失效.", -1)
	}

	token := tools.GenRandomString(36) + userIdStr

	tokenKey := tools.Keys{}.WxToken(token)
	tools.RdbSetExp(tokenKey, userIdStr, 5 * 60)

	rsp := models.TicketTokenRsp{
		Token: token,
	}
	app.OK(c, rsp, "")
}
//绑定
func (_ Api)Bind(c *gin.Context){
	objParams := models.BindReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err{
		tools.HasError(err, "参数错误", -1)
	}

	var userModel models.SchUser
	userModel.Phone = objParams.Mobile
	_, err =userModel.Get()
	if nil != err || ! gorm.IsRecordNotFoundError(err){
		tools.HasError(err, "手机号已被占用", -1)
	}

	userId := c.GetInt(models.UserId)
	_, err = userModel.Update(userId)
	if nil != err{
		tools.HasError(err, "手机号绑定失败,请重试", -1)
	}
	app.OK(c, nil, "手机号绑定成功")
}
// 学生明细
func (_ Api)StudentDetail(c *gin.Context){
	objParams := models.StudentDetailReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err{
		tools.HasError(err, "参数错误", -1)
	}


	userId := c.GetInt(models.UserId)
	userModel := models.SchUser{}
	userModel.Id = userId
	user, err := userModel.Get()
	if err != nil {
		tools.HasError(err, "用户异常", -1)
	}
	if user.Phone == ""{
		tools.HasError(err, "手机号", -1)
	}

	studentModel := models.ScbStudents{}
	studentModel.ParentPhone = userModel.Phone
	studentData, err := studentModel.Get()
	if err != nil{
		tools.HasError(err, "无法查询单相关学生信息", -1)
	}
	studentData.HeadImgSmall = config.ApplicationConfig.ImageUrl + strings.Replace(studentData.HeadImg, ".", "_small.", 1)
	studentData.HeadImg = config.ApplicationConfig.ImageUrl + studentData.HeadImg
	studentData.TimeString = studentData.CreatedAt.Format("2006-01-02 15:04:05")

	studentIdStr := strconv.Itoa(studentData.Id)
	//记录学生刷脸时间  用于标记上/下车状态 以及上/下车时间
	ymd := tools.Ymd()
	swipeAtKey := tools.Keys{}.SwipeAt(ymd, studentData.LineId)
	rStudentInfo, rsErr := tools.RdbHGet(swipeAtKey, studentIdStr)
	swipeData, rErr := redis.String(rStudentInfo, rsErr)

	var swipeAtInfo models.SwipeAt
	err = json.Unmarshal([]byte(swipeData), &swipeAtInfo)
	fmt.Println("swipeAtInfo SwipeAt err =>.", err, "swipeAtInfo SwipeAt rErr. => ", rErr, "swipeAtInfo => ", swipeAtInfo, "swipeData => ", swipeData)
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