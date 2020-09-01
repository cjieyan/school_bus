package models

import (
	"fmt"
	orm "go-admin/global"
	"go-admin/tools"
	"strconv"
	"strings"
)

type ScbLines struct {
	Id              int    `json:"id" gorm:"type:int(11);primary_key"`   //
	Name            string `json:"name" gorm:"type:varchar(250);"`       // 线路名称
	DepartedAt      string `json:"departed_at" gorm:"type:varchar(100);"`     // 出发时间
	ArrivedAt       string `json:"arrivedAt" gorm:"type:varchar(11);"`       // 到达时间
	ChangeExpiredAt string `json:"changeExpiredAt" gorm:"type:varchar(100);"` // 换站截止时间
	CarIds          string `json:"carIds" gorm:"type:varchar(200);"`     // 绑定的车辆
	IsDelete        int   `json:"isDelete" gorm:"type:tinyint(4);"`     // 0正常 1已删除
	DataScope       string `json:"dataScope" gorm:"-"`
	Params          string `json:"params"  gorm:"-"`
	CarIdsSelected  []int  `json:"carIdsSelected" gorm:"-"`
	CarIdsSubmit    []int  `json:"carIdsSubmit" gorm:"-"`
	BaseModel
}

func (ScbLines) TableName() string {
	return "scb_lines"
}

// 创建ScbLines
func (e *ScbLines) Create() (ScbLines, error) {
	var doc ScbLines
	carIds := ""
	for _, carId := range e.CarIdsSubmit{
		if len(carIds) > 0{
			carIds = carIds + "," + strconv.Itoa( carId )
		}else{
			carIds = strconv.Itoa( carId )
		}
	}
	fmt.Println("doc.CarIdsSubmit...,  carIds", e.CarIdsSubmit, carIds)
	e.CarIds = carIds

	doc.CarIds = carIds
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbLines
func (e *ScbLines) Get() (ScbLines, error) {
	var doc ScbLines
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.DepartedAt != "" {
		table = table.Where("departed_at = ?", e.DepartedAt)
	}

	if e.ArrivedAt != "" {
		table = table.Where("arrived_at = ?", e.ArrivedAt)
	}

	if e.ChangeExpiredAt != "" {
		table = table.Where("change_expired_at = ?", e.ChangeExpiredAt)
	}

	if e.CarIds != "" {
		table = table.Where("car_ids = ?", e.CarIds)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	carIds := doc.CarIds
	carIdsSelected := strings.Split(carIds, ",")
	for i:=0; i < len(carIdsSelected); i++ {
		carId , _ := strconv.Atoi(carIdsSelected[i])
		doc.CarIdsSelected = append(doc.CarIdsSelected, carId)
	}

	return doc, nil
}

// 获取ScbLines
func (e *ScbLines) GetCarIds() (ScbLines, error) {

	var doc ScbLines
	table := orm.Eloquent.Table(e.TableName())

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	carIds := doc.CarIds
	carIdsSelected := strings.Split(carIds, ",")
	for i:=0; i < len(carIdsSelected); i++ {
		carId , _ := strconv.Atoi(carIdsSelected[i])
		doc.CarIdsSelected = append(doc.CarIdsSelected, carId)
	}

	return doc, nil
}

// 获取ScbLines带分页
func (e *ScbLines) GetPage(pageSize int, pageIndex int) ([]ScbLines, int, error) {
	var doc []ScbLines

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.DepartedAt != "" {
		table = table.Where("departed_at = ?", e.DepartedAt)
	}

	if e.ArrivedAt != "" {
		table = table.Where("arrived_at = ?", e.ArrivedAt)
	}

	if e.ChangeExpiredAt != "" {
		table = table.Where("change_expired_at = ?", e.ChangeExpiredAt)
	}

	if e.CarIds != "" {
		table = table.Where("car_ids = ?", e.CarIds)
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

// 更新ScbLines
func (e *ScbLines) Update(id int) (update ScbLines, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	carIds := ""
	for _, carId := range e.CarIdsSubmit{
		if len(carIds) > 0{
			carIds = carIds + "," + strconv.Itoa( carId )
		}else{
			carIds = strconv.Itoa( carId )
		}
	}
	fmt.Println("doc.CarIdsSubmit..Update..,  carIds", e.CarIdsSubmit, carIds)
	e.CarIds = carIds
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除ScbLines
func (e *ScbLines) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&ScbLines{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbLines) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&ScbLines{}).Error; err != nil {
		return
	}
	Result = true
	return
}
//获取所有线路
func (e *ScbLines) GetAll()([]ScbLines, error){
	var doc []ScbLines
	table := orm.Eloquent.Select("*").Table(e.TableName())

	// 数据权限控制(如果不需要数据权限请将此处去掉)
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, err
	}
	table = table.Where("`is_delete` = 0")

	if err := table.Find(&doc).Error; err != nil {
		return nil, err
	}


	return doc, nil
}