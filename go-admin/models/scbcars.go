package models

import (
	"fmt"
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbCars struct {
	Id            int    `json:"id" gorm:"type:int(11);"`             //
	CarNumber     string `json:"carNumber" gorm:"type:varchar(100);"` // 车牌编号
	CarNo         string `json:"carNo" gorm:"type:varchar(50);"`      // 车牌号
	SeatsNum      string `json:"seatsNum" gorm:"type:int(11);"`       // 座位数
	AttendantId   int    `json:"attendantId" gorm:"type:int(11);"`    // 跟车员
	LineId        int    `json:"line_id gorm:"type:int(11)"`
	Driver        string `json:"driver" gorm:"type:varchar(50);"` // 司机
	Phone         string `json:"phone" gorm:"type:varchar(50);"`  // 手机号
	Dept          string `json:"dept" gorm:"type:int(11);"`       // 所属部门
	IsDelete      int    `json:"isDelete" gorm:"type:int(11);"`   // 0正常 1已删除
	DataScope     string `json:"dataScope" gorm:"-"`
	Params        string `json:"params"  gorm:"-"`
	AttendantName string `json:"attendantName" gorm:"-"`
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

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.CarNumber != "" {
		table = table.Where("car_number = ?", e.CarNumber)
	}

	if e.CarNo != "" {
		table = table.Where("car_no = ?", e.CarNo)
	}

	if e.Driver != "" {
		table = table.Where("driver = ?", e.Driver)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
	}

	if e.AttendantId != 0 {
		table = table.Where("attendant_id = ? ", e.AttendantId)
	}
	if e.LineId != 0 {
		table = table.Where("line_id = ? ", e.LineId)
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

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.CarNumber != "" {
		table = table.Where("car_number = ?", e.CarNumber)
	}

	if e.CarNo != "" {
		table = table.Where("car_no = ?", e.CarNo)
	}

	if e.Driver != "" {
		table = table.Where("driver = ?", e.Driver)
	}

	if e.Phone != "" {
		table = table.Where("phone = ?", e.Phone)
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
	var carsData []ScbCars
	for _, car := range doc {
		teacherModel := ScbTeachers{}
		teacherModel.Id = car.AttendantId
		teacherData, err := teacherModel.Get()
		if nil == err{
			car.AttendantName = teacherData.Name
		}
		carsData = append(carsData, car)
	}

	return carsData, count, nil
}

// 更新ScbCars 不可以更新"", 0, false值
func (e *ScbCars) Update(id int) (update ScbCars, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	fmt.Println("update.id->....", update.Id, "update.LineId->", update.LineId, "e.Id->", e.Id, "e.LineId->", e.LineId)
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 更新ScbCars  可以更新"", 0, false值
func (e *ScbCars) UpdateMap(id int, updateData map[string]interface{}) (update ScbCars, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	fmt.Println("update.id->....", update.Id, "update.LineId->", update.LineId, "e.Id->", e.Id, "e.LineId->", e.LineId)
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(updateData).Error; err != nil {
		return
	}
	return
}

// 删除ScbCars
func (e *ScbCars) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" id = ?", id).Delete(&ScbCars{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbCars) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where(" id in (?)", id).Delete(&ScbCars{}).Error; err != nil {
		return
	}
	Result = true
	return
}

func (e *ScbCars) GetAll() ([]MenuLable, error) {
	var doc []ScbCars
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

	m := make([]MenuLable, 0)
	for i := 0; i < len(doc); i++ {
		e := MenuLable{}
		e.Id = doc[i].Id
		e.Label = doc[i].CarNo
		m = append(m, e)
	}

	return m, nil
}

//多个id获取车辆列表
func (e ScbCars) GetbyIds(ids []int) ([]ScbCars, error) {
	var doc []ScbCars

	table := orm.Eloquent.Select("*").Table(e.TableName())

	err := table.
		Where("id in ( ? )", ids).
		Where("`is_delete` = 0").
		Find(&doc).Error
	return doc, err

}
