package models

import (
	orm "go-admin/global"
)

type UserAccount struct {
	Id            int    `json:"id" gorm:"type:int(11) unsigned;primary_key"` //
	UserId        int    `json:"userId" gorm:"type:int(11) unsigned;"`        // 用户id
	SumFund       string `json:"sumFund" gorm:"type:int(10) unsigned;"`       // 资金总量
	AvailableFund string `json:"availableFund" gorm:"type:int(10) unsigned;"` // 可用资金
	FrozenFund    string `json:"frozenFund" gorm:"type:int(10) unsigned;"`    // 冻结资金
	PaidFund      string `json:"paidFund" gorm:"type:int(11) unsigned;"`      // 消费总额
	ChargeFund    string `json:"chargeFund" gorm:"type:int(11) unsigned;"`    // 充值总额
	Point         string `json:"point" gorm:"type:int(11) unsigned;"`         // 点数
	CarTimes      string `json:"carTimes" gorm:"type:int(11) unsigned;"`      // 违章查询次数
	DataScope     string `json:"dataScope" gorm:"-"`
	Params        string `json:"params"  gorm:"-"`
	BaseModel
}

func (UserAccount) TableName() string {
	return "ak_user_account"
}

// 创建UserAccount
func (e *UserAccount) Create() (UserAccount, error) {
	var doc UserAccount
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取UserAccount
func (e *UserAccount) Get() (UserAccount, error) {
	var doc UserAccount
	table := orm.Eloquent.Table(e.TableName())

	if e.UserId != 0 {
		table = table.Where("user_id = ?", e.UserId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
