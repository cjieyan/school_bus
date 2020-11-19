package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbLineCars struct {
	Id        string `json:"id" gorm:"type:int(11);"`     //
	LineId    string `json:"lineId" gorm:"type:int(11);"` //
	CarId     string `json:"carId" gorm:"type:int(11);"`  //
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (ScbLineCars) TableName() string {
	return "scb_line_cars"
}

// 创建ScbLineCars
func (e *ScbLineCars) Create() (ScbLineCars, error) {
	var doc ScbLineCars
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbLineCars
func (e *ScbLineCars) Get() (ScbLineCars, error) {
	var doc ScbLineCars
	table := orm.Eloquent.Table(e.TableName())

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbLineCars带分页
func (e *ScbLineCars) GetPage(pageSize int, pageIndex int) ([]ScbLineCars, int, error) {
	var doc []ScbLineCars

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

// 更新ScbLineCars
func (e *ScbLineCars) Update(id int) (update ScbLineCars, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除ScbLineCars
func (e *ScbLineCars) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" = ?", id).Delete(&ScbLineCars{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbLineCars) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" in (?)", id).Delete(&ScbLineCars{}).Error; err != nil {
		return
	}
	Result = true
	return
}
