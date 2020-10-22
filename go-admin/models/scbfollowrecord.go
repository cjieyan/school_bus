package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbFollowRecord struct {
	Id          int    `json:"id" gorm:"type:int(11);primary_key"` //
	LineId      int `json:"lineId" gorm:"type:int(11);"`        // 线路id
	CarId       int `json:"carId" gorm:"type:int(11);"`         // 车辆id
	AttendantId int `json:"attendantId" gorm:"type:int(11);"`   // 跟车员 关联teachers表的id
	AllCount    int `json:"allCount" gorm:"type:int(11);"`      // 所有人数
	GetOff      int `json:"getOff" gorm:"type:int(11);"`        // 已下车数量
	GetOn       int `json:"getOn" gorm:"type:int(11);"`         // 已上车数量
	UnGetOn     int `json:"unGetOn" gorm:"type:int(11);"`       // 未上车数量
	Leave       int `json:"leave" gorm:"type:int(11);"`         // 请假数量
	IsDelete    int `json:"isDelete" gorm:"type:tinyint(4);"`   // 0未删除 1已删除
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	CreateBy    string `json:"createBy"  gorm:"type:varchar(50)"`
	UpdateBy    string `json:"updateBy"  gorm:"type:varchar(50)"`
	IsFinished  int    `json:"isFinished" gorm:"type:int(4)"` //跟车记录状态 0未结束 1已结束
	Car         ScbCars `json:"car", gorm:"-"`
	Line        ScbLines `json:"line" gorm:"-"`
	BaseModel
}

func (ScbFollowRecord) TableName() string {
	return "scb_follow_record"
}

// 创建ScbFollowRecord
func (e *ScbFollowRecord) Create() (ScbFollowRecord, error) {
	var doc ScbFollowRecord
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbFollowRecord
func (e *ScbFollowRecord) Get() (ScbFollowRecord, error) {
	var doc ScbFollowRecord
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.AttendantId != 0 {
		table = table.Where("attendant_id = ?", e.AttendantId)
	}

	if e.UnGetOn != 0 {
		table = table.Where("un_get_on = ?", e.UnGetOn)
	}

	if e.Leave != 0 {
		table = table.Where("leave = ?", e.Leave)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbFollowRecord带分页
func (e *ScbFollowRecord) GetPage(pageSize int, pageIndex int) ([]ScbFollowRecord, int, error) {
	var doc []ScbFollowRecord

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.AttendantId != 0 {
		table = table.Where("attendant_id = ?", e.AttendantId)
	}

	if e.UnGetOn != 0 {
		table = table.Where("un_get_on = ?", e.UnGetOn)
	}

	if e.Leave != 0 {
		table = table.Where("leave = ?", e.Leave)
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

// 更新ScbFollowRecord
func (e *ScbFollowRecord) Update(id int) (update ScbFollowRecord, err error) {
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

// 删除ScbFollowRecord
func (e *ScbFollowRecord) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&ScbFollowRecord{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbFollowRecord) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&ScbFollowRecord{}).Error; err != nil {
		return
	}
	Result = true
	return
}
