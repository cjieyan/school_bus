package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbCarRecord struct {
	Id        int    `json:"id" gorm:"type:int(11);primary_key"` //
	StudentId int `json:"studentId" gorm:"type:int(11);"`     // 学生id
	CarId     int `json:"carId" gorm:"type:int(11);"`         // 车辆id
	SiteId    int `json:"siteId" gorm:"type:int(11);"`        // 线路id
	DataScope string `json:"dataScope" gorm:"-"`
	CreateBy  string `json:"createBy" gorm:"size:64;"`
	UpdateBy  string `json:"updateBy" gorm:"size:64;"`
	QqLongitude  string `json:"qqLongitude" gorm:"size:100;"`
	QqLatitude  string `json:"qqLatitude" gorm:"size:100;"`
	Longitude  string `json:"longitude" gorm:"size:100;"`
	Latitude  string `json:"latitude" gorm:"size:100;"`
	Prop	int `json:"prop" gorm:"prop"`
	Type   int `json:"type" gorm:"type"` //0刷脸打卡,1手动打卡
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (ScbCarRecord) TableName() string {
	return "scb_car_record"
}

// 创建ScbCarRecord
func (e *ScbCarRecord) Create() (ScbCarRecord, error) {
	var doc ScbCarRecord
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbCarRecord
func (e *ScbCarRecord) Get() (ScbCarRecord, error) {
	var doc ScbCarRecord
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbCarRecord带分页
func (e *ScbCarRecord) GetPage(pageSize int, pageIndex int) ([]ScbCarRecord, int, error) {
	var doc []ScbCarRecord

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

// 更新ScbCarRecord
func (e *ScbCarRecord) Update(id int) (update ScbCarRecord, err error) {
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

// 删除ScbCarRecord
func (e *ScbCarRecord) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&ScbCarRecord{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbCarRecord) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&ScbCarRecord{}).Error; err != nil {
		return
	}
	Result = true
	return
}
