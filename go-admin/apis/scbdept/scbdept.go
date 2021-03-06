package scbdept

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"go-admin/tools/app/msg"
)

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/deptList [get]
// @Security Bearer
func GetDeptList(c *gin.Context) {
	var Dept models.ScbDept
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = tools.StringToInt(c.Request.FormValue("deptId"))
	Dept.DataScope = tools.GetUserIdStr(c)
	result, err := Dept.SetDept(true)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

func GetDeptTree(c *gin.Context) {
	var Dept models.ScbDept
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = tools.StringToInt(c.Request.FormValue("deptId"))
	result, err := Dept.SetDept(false)
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

// @Summary 部门列表数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func GetDept(c *gin.Context) {
	var Dept models.ScbDept
	Dept.DeptId, _ = tools.StringToInt(c.Param("deptId"))
	Dept.DataScope = tools.GetUserIdStr(c)
	result, err := Dept.Get()
	tools.HasError(err, msg.NotFound, 404)
	app.OK(c, result, msg.GetSuccess)
}

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body models.ScbDept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func InsertDept(c *gin.Context) {
	var data models.ScbDept
	err := c.MustBindWith(&data, binding.JSON)
	tools.HasError(err, "", 500)
	data.CreateBy = tools.GetUserIdStr(c)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, msg.CreatedSuccess)
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body models.ScbDept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [put]
// @Security Bearer
func UpdateDept(c *gin.Context) {
	var data models.ScbDept
	err := c.BindJSON(&data)
	tools.HasError(err, "", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.DeptId)
	tools.HasError(err, "", -1)
	app.OK(c, result, msg.UpdatedSuccess)
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dept/{id} [delete]
func DeleteDept(c *gin.Context) {
	var data models.ScbDept
	id, err := tools.StringToInt(c.Param("id"))
	_, err = data.Delete(id)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, "", msg.DeletedSuccess)
}

