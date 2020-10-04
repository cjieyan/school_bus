package models

import (
	"errors"
	orm "go-admin/global"
	"go-admin/tools"
	"time"
)

type ScbTeachers struct {
	Id        int    `json:"id" gorm:"type:int(10) unsigned;primary_key"` //
	Name      string `json:"name" gorm:"type:varchar(100);"`              // 名称
	Phone     string `json:"phone" gorm:"type:varchar(50);"`              // 手机号
	Password  string  `json:"password" gorm:"type:varchar(50);"`              //密码
	ClassId   int 	 `json:"classId" gorm:"type:int(11);"`                // 班级id
	PostId    int `json:"postId" gorm:"type:int(11);"`                 // 岗位id
	Remark    string `json:"remark" gorm:"type:varchar(200);"`            // 备注
	IsDeleted int `json:"isDeleted" gorm:"type:tinyint(4);"`           // 0未删除 1已删除
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (ScbTeachers) TableName() string {
	return "scb_teachers"
}

// 创建ScbTeachers
func (e *ScbTeachers) Create() (ScbTeachers, error) {
	var doc ScbTeachers
	e.CreatedAt = time.Now()
	pLen := len(e.Password)
	if pLen > 0 && pLen < 6{
		return doc, errors.New("密码大于等于6个字符")
	}
	e.Password, _ = tools.PasswordHash(e.Password)
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbTeachers
func (e *ScbTeachers) Get() (ScbTeachers, error) {
	var doc ScbTeachers
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
	}

	if e.ClassId != 0 {
		table = table.Where("class_id = ?", e.ClassId)
	}

	if e.PostId !=0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if e.IsDeleted != 0 {
		table = table.Where("is_deleted = ?", e.IsDeleted)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbTeachers带分页
func (e *ScbTeachers) GetPage(pageSize int, pageIndex int) ([]ScbTeachers, int, error) {
	var doc []ScbTeachers

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
	}

	if e.ClassId != 0 {
		table = table.Where("class_id = ?", e.ClassId)
	}

	if e.PostId !=0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if e.IsDeleted != 0 {
		table = table.Where("is_deleted = ?", e.IsDeleted)
	}

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
	table.Where("`is_delete` = 0").Count(&count)
	return doc, count, nil
}

// 更新ScbTeachers
func (e *ScbTeachers) Update(id int) (update ScbTeachers, err error) {

	pLen := len(e.Password)
	if pLen > 0 && pLen < 6{
		err = errors.New("密码大于等于6个字符")
		return
	}
	e.Password, _ = tools.PasswordHash(e.Password)
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

// 删除ScbTeachers
func (e *ScbTeachers) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).
		Where("id = ?", id).
		Delete(&ScbTeachers{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbTeachers) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&ScbTeachers{}).Error; err != nil {
		return
	}
	Result = true
	return
}
func (e *ScbTeachers) GetAttendants() ([]ScbTeachers, error){
	var doc []ScbTeachers

	table := orm.Eloquent.Select("*").
		Where("post_id = ?", 1).
		Table(e.TableName())
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
