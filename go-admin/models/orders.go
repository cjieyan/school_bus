package models

import (
        orm "go-admin/global"
)

type Orders struct {
	Id             int    `json:"id" gorm:"type:int(11);primary_key"`           //
	OrderNo        string `json:"orderNo" gorm:"type:varchar(100);"`            // 订单号
	TradeNo        string `json:"tradeNo" gorm:"type:varchar(100);"`            // 第三方流水号
	UserId         int `json:"userId" gorm:"type:int(11);"`                  // 用户id
	GoodsId        string `json:"goodsId" gorm:"type:int(11);"`                 // 商品id
	CarTimes       string `json:"carTimes" gorm:"type:int(11);"`                // 查询次数
	OrderAmount    string `json:"orderAmount" gorm:"type:int(11) unsigned;"`    // 订单金额
	TotalAmount    string `json:"totalAmount" gorm:"type:int(11) unsigned;"`    // 总金额
	DiscountAmount string `json:"discountAmount" gorm:"type:int(11) unsigned;"` // 优惠金额
	ActualAmount   string `json:"actualAmount" gorm:"type:int(11) unsigned;"`   // 实际支付金额
	PayType        string `json:"payType" gorm:"type:tinyint(4);"`              // 0支付宝, 1微信支付
	Type           string `json:"type" gorm:"type:tinyint(4);"`                 // 0充值 1消费
	CouponId       string `json:"couponId" gorm:"type:int(11);"`                // 优惠券id
	Status         string `json:"status" gorm:"type:tinyint(4);"`               // 0:未支付 10:已支付 20:订单作废 30:已发货 40: 已退款
	PaidAt         string `json:"paidAt" gorm:"type:int(11);"`                  // 支付时间
	DataScope      string `json:"dataScope" gorm:"-"`
	Params         string `json:"params"  gorm:"-"`
	BaseModel
}

func (Orders) TableName() string {
	return "ak_orders"
}

// 创建Orders
func (e *Orders) Create() (Orders, error) {
	var doc Orders
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Orders
func (e *Orders) Get() (Orders, error) {
	var doc Orders
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

//待支付订单数量
func (e *Orders)WaitPayCount()int{
	table := orm.Eloquent.Table(e.TableName())
	filter := make(map[string]interface{})
	filter["user_id"] = e.UserId
	filter["type"] = 0
	filter["status"] = 0
	count := 0
	table.Where(filter).First(&count)
	return count
}