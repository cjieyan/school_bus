package wx

import (
	"fmt"
	"github.com/chanxuehong/rand"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	return
}


