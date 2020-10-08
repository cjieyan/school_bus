package tools

import (
	"encoding/json"
	"go-admin/models"
	"go-admin/tools"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"net/url"
)

type BdApi struct {
	AppID string "22794004"
	APIKey string "oLPVBkl3gkURkuZPdN13XefG"
	SecretKey string "qOsoDoVAkvotzLn4ismk4dMmDoNaUrim"
}

//获取token
func (b *BdApi) getToken() (token string) {
	var rsp models.BdApiGetTokenRsp
	rspKey := tools.Keys{}.BdApiTokenRsp()
	exist := tools.RdbCheck(rspKey)
	if exist{
		rspStr, err := tools.RdbGet(rspKey)
		if nil == err && len(rspStr) > 0{
			err = json.Unmarshal([]byte(rspStr), &rsp)
			token = rsp.AccessToken
			return
		}
	}

	urlStr := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" +
		b.APIKey +
		"&client_secret=" +
		b.SecretKey

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Println("getToken->Get failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}

	err = json.Unmarshal(content, &rsp)

	tools.RdbSetExp(rspKey, string(content), rsp.ExpiresIn - 5)
	return rsp.AccessToken
}
//人脸注册
func (b *BdApi) FacesetAdd(userId, image, imageType string) (faceToken string){
	token := b.getToken()

	urlStr := "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add" +
		"?access_token=" + token


	data := url.Values{"image":{image}, "image_type":{"BASE64"}, "group_id":{"classa"}, "user_id":{userId}}
	body := strings.NewReader(data.Encode())

	resp, err := http.Post(urlStr, "application/json", body)
	if err != nil {
		log.Println("FacesetAdd->Post failed:", err)
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}
	rsp := models.BdApiFacesetAddRsp{}
	err = json.Unmarshal(content, &rsp)

	faceToken = rsp.FaceToken
	return

}