package wx

import (
	"fmt"
	"github.com/chanxuehong/rand"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"log"
	"net/http"
	"strconv"
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

func (_ Api)20