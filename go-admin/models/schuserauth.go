package models

import (
	orm "go-admin/global"
	"go-admin/tools"
)

type SchUserAuth struct {
	Id        int    `json:"id" gorm:"type:int(11);primary_key"` //
	UserId    int `json:"userId" gorm:"type:int(11);"`        // 用户id
	Appid     string `json:"appid" gorm:"type:varchar(100);"`    //
	Openid    string `json:"openid" gorm:"type:varchar(100);"`   //
	Unionid   string `json:"unionid" gorm:"type:varchar(100);"`  //
	Source    string `json:"source" gorm:"type:varchar(50);"`    // 来源
	Nickname  string `json:"nickname" gorm:"type:varchar(100);"` // 昵称
	Avatar    string `json:"avatar" gorm:"type:varchar(200);"`   // 头像
	Sex       string `json:"sex" gorm:"type:int(11);"`           // 性别
	Province  string `json:"province" gorm:"type:varchar(100);"` // 省
	Country   string `json:"country" gorm:"type:varchar(100);"`  // 国家
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params"  gorm:"-"`
	BaseModel
}

func (SchUserAuth) TableName() string {
	return "sch_user_auth"
}

// 创建SchUserAuth
func (e *SchUserAuth) Create() (SchUserAuth, error) {
	var doc SchUserAuth
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SchUserAuth
func (e *SchUserAuth) Get() (SchUserAuth, error) {
	var doc SchUserAuth
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SchUserAuth带分页
func (e *SchUserAuth) GetPage(pageSize int, pageIndex int) ([]SchUserAuth, int, error) {
	var doc []SchUserAuth

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

// 更新SchUserAuth
func (e *SchUserAuth) Update(id int) (update SchUserAuth, err error) {
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

// 删除SchUserAuth
func (e *SchUserAuth) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&SchUserAuth{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *SchUserAuth) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&SchUserAuth{}).Error; err != nil {
		return
	}
	Result = true
	return
}
