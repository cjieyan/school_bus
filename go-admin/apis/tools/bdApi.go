package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"go-admin/models"
	"go-admin/tools"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	AppID     string = "22794004"
	APIKey    string = "oLPVBkl3gkURkuZPdN13XefG"
	SecretKey string = "qOsoDoVAkvotzLn4ismk4dMmDoNaUrim"
	GroupId   string = "classa"
)

type BdApi struct {
}

//获取token
func (b *BdApi) getToken() (token string) {
	var rsp models.BdApiGetTokenRsp
	rspKey := tools.Keys{}.BdApiTokenRsp()
	exist := tools.RdbCheck(rspKey)
	if exist {
		rspStr, err := tools.RdbGet(rspKey)
		if nil == err && len(rspStr) > 0 {
			err = json.Unmarshal([]byte(rspStr), &rsp)
			return rsp.AccessToken
		}
	}
	urlStr := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" +
		APIKey +
		"&client_secret=" +
		SecretKey

	resp, err := http.Post(urlStr, "application/json", nil)
	//fmt.Printf("getToken url: -> %v resp:-> %v, err -> %#v\n", urlStr, resp, err)
	//fmt.Println("---")
	if err != nil {
		fmt.Println("getToken->Get failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	err = json.Unmarshal(content, &rsp)
	expiredAt := int(time.Now().Unix()) + rsp.ExpiresIn
	err = tools.RdbSetExp(rspKey, string(content), expiredAt)

	return rsp.AccessToken
}

//人脸注册
func (b *BdApi) FacesetAdd(userId, image string) (faceToken string) {
	token := b.getToken()

	urlStr := "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add" +
		"?access_token=" + token
	data := url.Values{"image": {image},
		"image_type":  {"BASE64"},
		"group_id":    {GroupId},
		"user_id":     {userId},
		"action_type": {"REPLACE"},
	}
	body := strings.NewReader(data.Encode())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	req, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		fmt.Println("post req err -> ", err)
		return
	}
	req.Header.Set("contentType", "application/json")
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))

	if err != nil {
		fmt.Println("post req err -> ", err)
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}
	rsp := models.BdApiFacesetAddRsp{}
	err = json.Unmarshal(content, &rsp)
	fmt.Println("FacesetAdd content: ----> ", image,  string(content), rsp.FaceToken)

	return rsp.FaceToken
}
