package models

import (
	orm "go-admin/global"
	"go-admin/tools"
	_ "time"
)

type ScbDept struct {
	DeptId    int       `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	ParentId  int       `json:"parentId" gorm:""`                         //上级部门
	DeptPath  string    `json:"deptPath" gorm:"size:255;"`                //
	DeptName  string    `json:"deptName"  gorm:"size:128;"`               //部门名称
	Sort      int       `json:"sort" gorm:""`                             //排序
	Leader    string    `json:"leader" gorm:"size:128;"`                  //负责人
	Phone     string    `json:"phone" gorm:"size:11;"`                    //手机
	Email     string    `json:"email" gorm:"size:64;"`                    //邮箱
	Status    string    `json:"status" gorm:"size:4;"`                    //状态
	CreateBy  string    `json:"createBy" gorm:"size:64;"`
	UpdateBy  string    `json:"updateBy" gorm:"size:64;"`
	DataScope string    `json:"dataScope" gorm:"-"`
	Params    string    `json:"params" gorm:"-"`
	Children  []ScbDept `json:"children" gorm:"-"`
	Id			int 	`json:"id" gorm:"-"`
	Value     int 	`json:"value" gorm:"-"`
	Label     string `json:"label" gorm:"-"`
	BaseModel
}

func (ScbDept) TableName() string {
	return "scb_dept"
}

type ScbDeptLable struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []ScbDeptLable `gorm:"-" json:"children"`
}

func (e *ScbDept) Create() (ScbDept, error) {
	var doc ScbDept
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	deptPath := "/" + tools.IntToString(e.DeptId)
	if int(e.ParentId) != 0 {
		var deptP ScbDept
		orm.Eloquent.Table(e.TableName()).Where("dept_id = ?", e.ParentId).First(&deptP)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	var mp = map[string]string{}
	mp["deptPath"] = deptPath
	if err := orm.Eloquent.Table(e.TableName()).Where("dept_id = ?", e.DeptId).Update(mp).Error; err != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	doc.DeptPath = deptPath
	return doc, nil
}

func (e *ScbDept) Get() (ScbDept, error) {
	var doc ScbDept

	table := orm.Eloquent.Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *ScbDept) GetList() ([]ScbDept, error) {
	var doc []ScbDept

	table := orm.Eloquent.Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *ScbDept) GetPage(bl bool) ([]ScbDept, error) {
	var doc []ScbDept

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.DeptId != 0 {
		table = table.Where("dept_id = ?", e.DeptId)
	}
	if e.DeptName != "" {
		table = table.Where("dept_name = ?", e.DeptName)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}
	if e.DeptPath != "" {
		table = table.Where("deptPath like %?%", e.DeptPath)
	}
	if bl {
		// 数据权限控制
		dataPermission := new(DataPermission)
		dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
		tableper, err := dataPermission.GetDataScope("scp_dept", table)
		if err != nil {
			return nil, err
		}
		table = tableper
	}

	if err := table.Order("sort").Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (e *ScbDept) SetDept(bl bool) ([]ScbDept, error) {
	list, err := e.GetPage(bl)

	m := make([]ScbDept, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.Digui(&list, list[i])
		info.Id = info.DeptId
		info.Label = info.DeptName
		info.Value = info.DeptId
		m = append(m, info)
	}
	return m, err
}

func (e ScbDept)Digui(deptlist *[]ScbDept, menu ScbDept) ScbDept {
	list := *deptlist

	min := make([]ScbDept, 0)
	for j := 0; j < len(list); j++ {

		if menu.DeptId != list[j].ParentId {
			continue
		}
		mi := ScbDept{}

		mi.Id = list[j].DeptId
		mi.Label = list[j].DeptName
		mi.Value = list[j].DeptId
		mi.DeptId = list[j].DeptId
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.Children = []ScbDept{}
		ms := e.Digui(deptlist, mi)
		min = append(min, ms)

	}
	menu.Children = min
	return menu
}

func (e *ScbDept) Update(id int) (update ScbDept, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("dept_id = ?", id).First(&update).Error; err != nil {
		return
	}

	deptPath := "/" + tools.IntToString(e.DeptId)
	if e.ParentId != 0 {
		var deptP ScbDept
		orm.Eloquent.Table(e.TableName()).Where("dept_id = ?", e.ParentId).First(&deptP)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	e.DeptPath = deptPath

	//参数1:是要修改的数据
	//参数2:是修改的数据

	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}

	return
}

func (e *ScbDept) Delete(id int) (success bool, err error) {

	user := SysUser{}
	user.DeptId = id
	userlist, err := user.GetList()
	tools.HasError(err, "", 500)
	tools.Assert(len(userlist) <= 0, "当前部门存在用户，不能删除！", 500)

	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		success = false
		return
	}

	if err = tx.Table(e.TableName()).Where("dept_id = ?", id).Delete(&ScbDept{}).Error; err != nil {
		success = false
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		success = false
		return
	}
	success = true

	return
}

func (dept *ScbDept) SetScbDeptLable() (m []ScbDeptLable, err error) {
	deptlist, err := dept.GetList()

	m = make([]ScbDeptLable, 0)
	for i := 0; i < len(deptlist); i++ {
		if deptlist[i].ParentId != 0 {
			continue
		}
		e := ScbDeptLable{}
		e.Id = deptlist[i].DeptId
		e.Label = deptlist[i].DeptName
		deptsInfo := dept.DiguiScbDeptLable(&deptlist, e)

		m = append(m, deptsInfo)
	}
	return
}

func (e ScbDept)DiguiScbDeptLable(deptlist *[]ScbDept, dept ScbDeptLable) ScbDeptLable {
	list := *deptlist

	min := make([]ScbDeptLable, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := ScbDeptLable{list[j].DeptId, list[j].DeptName, []ScbDeptLable{}}
		ms := e.DiguiScbDeptLable(deptlist, mi)
		min = append(min, ms)

	}
	dept.Children = min
	return dept
}
