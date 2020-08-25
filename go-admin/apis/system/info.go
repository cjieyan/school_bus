package system

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
)

func GetInfo(c *gin.Context) {

	var roles = make([]string, 1)
	roles[0] = tools.GetRoleName(c)

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	RoleMenu := models.RoleMenu{}
	RoleMenu.RoleId = tools.GetRoleId(c)

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if tools.GetRoleName(c) == "admin" || tools.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := RoleMenu.GetPermis()
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysuser := models.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	user, err := sysuser.Get()
	tools.HasError(err, "", 500)

	mp["introduction"] = " am a super administrator"

	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["userName"] = user.NickName
	mp["userId"] = user.UserId
	mp["deptId"] = user.DeptId
	mp["name"] = user.NickName


	mp["carsCount"] = 0

	app.OK(c, mp, "")
}
func GetPreView(c *gin.Context) {

	sysuser := models.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	var mp = make(map[string]interface{})

	count := 0
	//车辆数
	mp["carsCount"] = count

	userAccount := models.UserAccount{
		UserId: sysuser.UserId,
	}
	//剩余查询次数
	data, err := userAccount.Get()
	if nil == err && data.Id > 0 {
		mp["times"] = data.CarTimes
	} else {
		mp["times"] = 0
	}

	ordersModel := models.Orders{UserId: sysuser.UserId}
	waitPayCount := ordersModel.WaitPayCount()
	//待支付订单
	mp["waitPayCount"] = waitPayCount
	//违章车辆
	illegalCount := 0
	mp["illegalCount"] = illegalCount

	app.OK(c, mp, "")
}
