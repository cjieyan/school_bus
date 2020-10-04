package xcx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"strconv"
)

type Api struct {

}
func (a Api)Login(c *gin.Context){
	objParams := models.XcxLoginReq{}
	err := c.ShouldBindJSON(&objParams)
	if nil != err{
		tools.HasError(err, "", -1)
	}

	model := models.ScbTeachers{}
	model.Phone = objParams.Phone
	teacher, err := model.Get()
	tools.HasError(err, "账号或密码错误", -1)

	err = tools.CompareHashAndPassword(objParams.Password, teacher.Password)
	tools.HasError(err, "账号或密码错误.", -1)

	idStr := fmt.Sprint("%i", model.Id)
	token := tools.GenRandomString(36) + idStr
	key := tools.Keys{}.ApiToken(token)

	tools.RdbSet(key, idStr)
	rsp := models.XcxLoginRsp{}
	rsp.Token = token

	app.OK(c, rsp, "")
}

func (a Api)Info(c *gin.Context){
	token := c.GetHeader("token")
	uidStr, err := tools.RdbGet(token)
	tools.HasError(err, "会话过期,请重新登录", -1)
	uid, _ := strconv.Atoi(uidStr)

	model := models.ScbTeachers{}
	model.Id = uid
	info, err := model.Get()
	tools.HasError(err, "账号异常,请联系管理员", -1)

	app.OK(c, info, "")

}