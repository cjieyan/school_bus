package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type CarsRecord struct {
	Id         int    `json:"id" gorm:"type:int(11);primary_key"`   //
	UserId     string `json:"userId" gorm:"type:int(11);"`          //
	CarId      string `json:"carId" gorm:"type:int(11);"`           //
	CarNo      string `json:"carNo" gorm:"type:varchar(100);"`      //
	VioTotal   string `json:"vioTotal" gorm:"type:int(11);"`        //
	FindTotal  string `json:"findTotal" gorm:"type:int(11);"`       //
	ScoreTotal string `json:"scoreTotal" gorm:"type:int(11);"`      //
	OrderNo    string `json:"orderNo" gorm:"type:varchar(100);"`    //
	AkOrdersNo string `json:"akOrdersNo" gorm:"type:varchar(100);"` //
	Status     string `json:"status" gorm:"type:tinyint(4);"`       //
	Code       string `json:"code" gorm:"type:int(11);"`            //
	Msg        string `json:"msg" gorm:"type:varchar(200);"`        //
	FromSite   string `json:"fromSite" gorm:"type:varchar(50);"`    //
	Hide       string `json:"hide" gorm:"type:tinyint(4);"`         //
	DataScope  string `json:"dataScope" gorm:"-"`
	Params     string `json:"params"  gorm:"-"`
	BaseModel
}

func (CarsRecord) TableName() string {
	return "ak_cars_record"
}

// 创建CarsRecord
func (e *CarsRecord) Create() (CarsRecord, error) {
	var doc CarsRecord
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取CarsRecord
func (e *CarsRecord) Get() (CarsRecord, error) {
	var doc CarsRecord
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取CarsRecord带分页
func (e *CarsRecord) GetPage(pageSize int, pageIndex int) ([]CarsRecord, int, error) {
	var doc []CarsRecord

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

// 更新CarsRecord
func (e *CarsRecord) Update(id int) (update CarsRecord, err error) {
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

// 删除CarsRecord
func (e *CarsRecord) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&CarsRecord{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *CarsRecord) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&CarsRecord{}).Error; err != nil {
		return
	}
	Result = true
	return
}
