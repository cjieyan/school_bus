package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbPost struct {
	PostId    int    `json:"postId" gorm:"type:int(11);primary_key"` //
	PostName  string `json:"postName" gorm:"type:varchar(128);"`     // 岗位名称
	PostCode  string `json:"postCode" gorm:"type:varchar(128);"`     // 岗位代码
	Sort      int `json:"sort" gorm:"type:int(11);"`              // 岗位排序
	Status    int `json:"status" gorm:"type:int(4);"`         // 状态
	Remark    string `json:"remark" gorm:"type:varchar(255);"`       //
	CreateBy  string `json:"createBy" gorm:"type:varchar(128);"`     //
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(128);"`     //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (ScbPost) TableName() string {
	return "scb_post"
}

// 创建ScbPost
func (e *ScbPost) Create() (ScbPost, error) {
	var doc ScbPost
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbPost
func (e *ScbPost) Get() (ScbPost, error) {
	var doc ScbPost
	table := orm.Eloquent.Table(e.TableName())

	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}

	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}

	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbPost带分页
func (e *ScbPost) GetPage(pageSize int, pageIndex int) ([]ScbPost, int, error) {
	var doc []ScbPost

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}

	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}

	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
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
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

// 更新ScbPost
func (e *ScbPost) Update(id int) (update ScbPost, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("post_id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除ScbPost
func (e *ScbPost) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("post_id = ?", id).Delete(&ScbPost{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbPost) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("post_id in (?)", id).Delete(&ScbPost{}).Error; err != nil {
		return
	}
	Result = true
	return
}


func (e *ScbPost) GetList() ([]ScbPost, error) {
	var doc []ScbPost

	table := orm.Eloquent.Table(e.TableName())
	if e.PostId != 0 {
		table = table.Where("post_id = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("post_name = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("post_code = ?", e.PostCode)
	}
	if e.Status != 0 {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}