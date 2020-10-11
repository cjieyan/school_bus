package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type ScbStudents struct {
	Id            int    `json:"id" gorm:"type:int(11);primary_key"` //
	Name          string `json:"name" gorm:"type:varchar(100);"`     // 名称
	Number        string `json:"number" gorm:"type:varchar(100);"`   // 学号
	ClassId       int    `json:"classId" gorm:"type:int(11);"`       // 班级id
	LineId        int    `json:"lineId" gorm:"type:int(11);"`        // 线路id
	SiteIdUp      int    `json:"siteIdUp" gorm:"type:int(11);"`      // 站点id
	SiteIdDown    int    `json:"siteIdDown" gorm:"type:int(11);"`    // 站点id
	CarId         int    `json:"carId" gorm:"type:int(11);"`         // 车辆id
	IsPickUp      int    `json:"isPickUp" gorm:"type:tinyint(4);"`
	ParentPhone   string `json:"parentPhone" gorm:"type:varchar(50);"` // 家长电话
	Picture       string `json:"picture" gorm:"type:varchar(200);"`    // 图片
	IsDeleted     int    `json:"isDeleted" gorm:"type:tinyint(4);"`    // 0未删除 1已删除
	DataScope     string `json:"dataScope" gorm:"-"`
	Params        string `json:"params"  gorm:"-"`
	FaceToken     string `json:"faceToken" gorm:"type:varchar(255);"`
	LogId         int `json:"logId" gorm:"type:int(10);"`
	BaiduGroupId  string `json:"baidu_group_id" gorm:"type:varchar(255);"`
	BaiduUserId   string `json:"baidu_user_id" gorm:"type:varchar(255);"`
	BaiduUserInfo string `json:"baidu_user_info" gorm:"type:varchar(255);"`
	BaseModel
}

func (ScbStudents) TableName() string {
	return "scb_students"
}

// 创建ScbStudents
func (e *ScbStudents) Create() (ScbStudents, error) {
	var doc ScbStudents
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取ScbStudents
func (e *ScbStudents) Get() (ScbStudents, error) {
	var doc ScbStudents
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Number != "" {
		table = table.Where("number = ?", e.Number)
	}

	if e.ClassId != 0 {
		table = table.Where("class_id = ?", e.ClassId)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.SiteIdUp != 0 {
		table = table.Where("site_id_up = ?", e.SiteIdUp)
	}

	if e.SiteIdDown != 0 {
		table = table.Where("site_id_down = ?", e.SiteIdDown)
	}

	if e.CarId != 0 {
		table = table.Where("car_id = ?", e.CarId)
	}

	if e.ParentPhone != "" {
		table = table.Where("parent_phone = ?", e.ParentPhone)
	}

	if e.Picture != "" {
		table = table.Where("picture = ?", e.Picture)
	}

	if e.IsDeleted != 0 {
		table = table.Where("is_deleted = ?", e.IsDeleted)
	}

	if e.FaceToken != "" {
		table = table.Where("face_token = ?", e.FaceToken)
	}

	if e.LogId != 0 {
		table = table.Where("log_id = ?", e.LogId)
	}

	if e.BaiduGroupId != "" {
		table = table.Where("baidu_group_id = ?", e.BaiduGroupId)
	}

	if e.BaiduUserId != "" {
		table = table.Where("baidu_user_id = ?", e.BaiduUserId)
	}

	if e.BaiduUserInfo != "" {
		table = table.Where("baidu_user_info = ?", e.BaiduUserInfo)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取ScbStudents带分页
func (e *ScbStudents) GetPage(pageSize int, pageIndex int) ([]ScbStudents, int, error) {
	var doc []ScbStudents

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Number != "" {
		table = table.Where("number = ?", e.Number)
	}

	if e.ClassId != 0 {
		table = table.Where("class_id = ?", e.ClassId)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.SiteIdUp != 0 {
		table = table.Where("site_id = ?", e.SiteIdUp)
	}

	if e.SiteIdDown != 0 {
		table = table.Where("site_id = ?", e.SiteIdDown)
	}

	if e.IsPickUp != 0 {
		table = table.Where("is_pick_up = ?", e.IsPickUp)
	}

	if e.CarId != 0 {
		table = table.Where("car_id = ?", e.CarId)
	}

	if e.ParentPhone != "" {
		table = table.Where("parent_phone = ?", e.ParentPhone)
	}

	if e.Picture != "" {
		table = table.Where("picture = ?", e.Picture)
	}

	// BaiduGroupId
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

// 更新ScbStudents
func (e *ScbStudents) Update(id int) (update ScbStudents, err error) {
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

// 删除ScbStudents
func (e *ScbStudents) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&ScbStudents{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *ScbStudents) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&ScbStudents{}).Error; err != nil {
		return
	}
	Result = true
	return
}
