package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-admin/models"
	"go-admin/tools"
	"io/ioutil"
	"net/http"
	"strconv"
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

type RequestBody struct {
	Id int `json:"id"`

	Name string `json:"name"`
}

//人脸注册
func (b *BdApi) FacesetAdd(userId, image string) (faceToken string) {
	if "" == image {
		return ""
	}
	imageArr := strings.Split(image, ";base64,")
	if len(imageArr) > 1 {
		image = imageArr[1]
	}

	token := b.getToken()

	urlStr := "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add" +
		"?access_token=" + token

	reqModel := models.BdApiFacesetAddReq{}
	reqModel.Image = image
	reqModel.ImageType = "BASE64"
	reqModel.GroupId = GroupId
	reqModel.UserId = userId
	reqModel.ActionType = "REPLACE"

	requestBody := new(bytes.Buffer)

	json.NewEncoder(requestBody).Encode(reqModel)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	req, err := http.NewRequest("POST", urlStr, requestBody)
	if err != nil {
		fmt.Println("post req err -> ", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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
	fmt.Println("FacesetAdd content: ----> ", string(content), rsp.ErrorCode)

	if 0 == rsp.ErrorCode {
		return rsp.Result.FaceToken
	}
	fmt.Println("FacesetAdd    rsp....", rsp)
	return ""
}

// 多个人脸识别
func (b *BdApi) MutilSearch(image string) (ret []int) {
	token := b.getToken()
	urlStr := "https://aip.baidubce.com/rest/2.0/face/v3/multi-search" +
		"?access_token=" + token

	reqModel := models.BdApiFacesetMutilSearchReq{}
	reqModel.Image = image
	reqModel.ImageType = "BASE64"
	reqModel.GroupIdList = GroupId
	reqModel.MaxFaceNum = 10
	reqModel.QualityControl = "NORMAL"
	reqModel.LivenessControl = "NORMAL"
	reqModel.MaxUserNum = 10

	requestBody := new(bytes.Buffer)

	json.NewEncoder(requestBody).Encode(reqModel)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	fmt.Println("requestBody", requestBody)
	req, err := http.NewRequest("POST", urlStr, requestBody)
	if err != nil {
		fmt.Println("post req err -> ", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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
	rsp := models.BdApiFacesetMutilSearchRsp{}
	err = json.Unmarshal(content, &rsp)
	fmt.Println("MutilSearch content: ----> ", string(content), rsp.ErrorCode)

	faceTokenArray := []int{}
	if 0 == rsp.ErrorCode && rsp.Result.FaceNum > 0 {
		for _, faceList := range rsp.Result.FaceList {

			for _, faceUser := range faceList.UserList {
				if faceUser.Score > 80 {
					userId, _ := strconv.Atoi(faceUser.UserID)
					faceTokenArray = append(faceTokenArray, userId)
				}
			}
			// for _, user := face.UserList.UserList{
			// 	faceTokenArray = append(faceTokenArray, user.U
			// }
		}
		ret = faceTokenArray
	}
	return

}
