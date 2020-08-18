package useraccount

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
)

func GetUserAccount(c *gin.Context) {
	var data models.UserAccount
	data.Id, _ = tools.StringToInt(c.Param("id"))
	result, err := data.Get()
	tools.HasError(err, "抱歉未找到相关信息", -1)

	app.OK(c, result, "")
}

func InsertUserAccount(c *gin.Context) {
	var data models.UserAccount
	err := c.ShouldBindJSON(&data)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}
