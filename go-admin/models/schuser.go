package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type SchUser struct {
	Id        int    `json:"id" gorm:"type:int(11);primary_key"` //
	Username  string `json:"username" gorm:"type:varchar(100);"` // 用户名
	Password  string `json:"password" gorm:"type:varchar(100);"` // 密码
	Phone     string `json:"phone" gorm:"type:varchar(100);"`    // 手机号
	Nickname  string `json:"nickname" gorm:"type:varchar(100);"` // 昵称
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (SchUser) TableName() string {
	return "sch_user"
}

// 创建SchUser
func (e *SchUser) Create() (SchUser, error) {
	var doc SchUser
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SchUser
func (e *SchUser) Get() (SchUser, error) {
	var doc SchUser
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SchUser带分页
func (e *SchUser) GetPage(pageSize int, pageIndex int) ([]SchUser, int, error) {
	var doc []SchUser

	table := orm.Eloquent.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 更新SchUser
func (e *SchUser) Update(id int) (update SchUser, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除SchUser
func (e *SchUser) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&SchUser{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *SchUser) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&SchUser{}).Error; err != nil {
		return
	}
	Result = true
	return
}
