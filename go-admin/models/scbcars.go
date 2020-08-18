package models

import (
orm "go-admin/global"
"go-admin/tools"


)

type ScbCars struct {

        Id int `json:"id" gorm:"type:int(11);"` //
        CarNumber string `json:"carNumber" gorm:"type:varchar(100);"` // 车牌编号
        CarNo string `json:"carNo" gorm:"type:varchar(50);"` // 车牌号
        SeatsNum string `json:"seatsNum" gorm:"type:int(11);"` // 座位数
        AttendantId string `json:"attendantId" gorm:"type:int(11);"` // 跟车员
        Driver string `json:"driver" gorm:"type:int(11);"` // 司机
        Phone string `json:"phone" gorm:"type:varchar(50);"` // 手机号
        Dept string `json:"dept" gorm:"type:int(11);"` // 所属部门
        IsDelete string `json:"isDelete" gorm:"type:int(11);"` // 0正常 1已删除
DataScope   string `json:"dataScope" gorm:"-"`
Params      string `json:"params"  gorm:"-"`
BaseModel
}

func (ScbCars) TableName() string {
return "scb_cars"
}

// 创建ScbCars
func (e *ScbCars) Create() (ScbCars, error) {
var doc ScbCars
result := orm.Eloquent.Table(e.TableName()).Create(&e)
if result.Error != nil {
err := result.Error
return doc, err
}
doc = *e
return doc, nil
}


// 获取ScbCars
func (e *ScbCars) Get() (ScbCars, error) {
var doc ScbCars
table := orm.Eloquent.Table(e.TableName())

    
        if e.Id != 0  {
        table = table.Where("id = ?", e.Id)
        }
    
    
        if e.CarNumber != ""  {
        table = table.Where("car_number = ?", e.CarNumber)
        }
    
    
        if e.CarNo != ""  {
        table = table.Where("car_no = ?", e.CarNo)
        }
    
    
    
    
        if e.Driver != ""  {
        table = table.Where("driver = ?", e.Driver)
        }
    
    
        if e.Phone != ""  {
        table = table.Where("phone = ?", e.Phone)
        }
    
    
    
    
    

if err := table.First(&doc).Error; err != nil {
return doc, err
}
return doc, nil
}

// 获取ScbCars带分页
func (e *ScbCars) GetPage(pageSize int, pageIndex int) ([]ScbCars, int, error) {
var doc []ScbCars

table := orm.Eloquent.Select("*").Table(e.TableName())

        if e.Id != 0  {
        table = table.Where("id = ?", e.Id)
        }
    
        if e.CarNumber != ""  {
        table = table.Where("car_number = ?", e.CarNumber)
        }
    
        if e.CarNo != ""  {
        table = table.Where("car_no = ?", e.CarNo)
        }
    
        if e.Driver != ""  {
        table = table.Where("driver = ?", e.Driver)
        }
    
        if e.Phone != ""  {
        table = table.Where("phone = ?", e.Phone)
        }
    

// 数据权限控制(如果不需要数据权限请将此处去掉)
dataPermission := new(DataPermission)
dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
table,err := dataPermission.GetDataScope(e.TableName(), table)
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

// 更新ScbCars
func (e *ScbCars) Update(id int) (update ScbCars, err error) {
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

// 删除ScbCars
func (e *ScbCars) Delete(id int) (success bool, err error) {
if err = orm.Eloquent.Table(e.TableName()).Where(" = ?", id).Delete(&ScbCars{}).Error; err != nil {
success = false
return
}
success = true
return
}

//批量删除
func (e *ScbCars) BatchDelete(id []int) (Result bool, err error) {
if err = orm.Eloquent.Table(e.TableName()).Where(" in (?)", id).Delete(&ScbCars{}).Error; err != nil {
return
}
Result = true
return
}
