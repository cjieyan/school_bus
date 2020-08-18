package models

import (
orm "go-admin/global"
"go-admin/tools"


)

type Cars struct {

        Id int `json:"id" gorm:"type:int(11);primary_key"` // 
        UserId int `json:"userId" gorm:"type:int(11);"` //
        Cartype int `json:"cartype" gorm:"type:varchar(10);"` // 车辆类型
        Carno string `json:"carno" gorm:"type:varchar(50);"` // 车牌号
        Frameno string `json:"frameno" gorm:"type:varchar(100);"` // 车架号
        Engineno string `json:"engineno" gorm:"type:varchar(100);"` // 发动机号
        Purpose string `json:"purpose" gorm:"type:varchar(500);"` // 用途
        Company string `json:"company" gorm:"type:varchar(200);"` // 公司
        Driver string `json:"driver" gorm:"type:varchar(100);"` // 
        DriverPhone string `json:"driverPhone" gorm:"type:varchar(50);"` // 司机号码
        Team string `json:"team" gorm:"type:varchar(200);"` // 
        Remark string `json:"remark" gorm:"type:varchar(500);"` // 
        Status int `json:"status" gorm:"type:tinyint(4);"` // 状态0正常 1禁用
        CarnoMd5 string `json:"carnoMd5" gorm:"type:varchar(32);"` // 车牌号md5
        RenewAt string `json:"renewAt" gorm:"type:int(11);"` // 最后一次更新时间
        SuccMonth string `json:"succMonth" gorm:"type:int(11);"` // 查询成功的年月
        CreatedAt int `json:"created_at" gorm:"created_at:int(11);"` // 最后一次更新时间
        UpdatedAt int `json:"updated_at" gorm:"updated_at:int(11);"` // 最后一次更新时间

        IllegalStatus int `json:"illegalStatus" gorm:"type:tinyint(4);"` // 0 无违章 1有违章
        DataScope   string `json:"dataScope" gorm:"-"`
        Params      string `json:"params"  gorm:"-"`
        BaseModel
}

func (Cars) TableName() string {
        return "ak_cars"
}

// 创建Cars
func (e *Cars) Create() (Cars, error) {
        var doc Cars
        result := orm.Eloquent.Table(e.TableName()).Create(&e)
        if result.Error != nil {
        err := result.Error
        return doc, err
        }
        doc = *e
        return doc, nil
}


// 获取Cars
func (e *Cars) Get() (Cars, error) {
        var doc Cars
        table := orm.Eloquent.Table(e.TableName())

        if e.Id != 0 {
                table = table.Where("id = ?", e.Id)
        }
        if err := table.First(&doc).Error; err != nil {
        return doc, err
        }
        return doc, nil
}

// 获取Cars带分页
func (e *Cars) GetPage(pageSize int, pageIndex int) ([]Cars, int, error) {
        var doc []Cars

        table := orm.Eloquent.Select("*").Table(e.TableName())


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

        // 更新Cars
        func (e *Cars) Update(id int) (update Cars, err error) {
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

// 删除Cars
func (e *Cars) Delete(id int) (success bool, err error) {
        if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&Cars{}).Error; err != nil {
        success = false
        return
        }
        success = true
        return
}

//批量删除
func (e *Cars) BatchDelete(id []int) (Result bool, err error) {
        if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Cars{}).Error; err != nil {
        return
        }
        Result = true
        return
}
//获取用户车辆总数
func (e *Cars)GetCount() int{
        filter := make(map[string]interface{})
        filter["user_id"] = e.UserId
        filter["status"] = 0
        count := 0
        err := orm.Eloquent.
                Table(e.TableName()).
                Where(filter).
                Count(&count).Error
        if nil != err{
                return 0
        }else{
                return count
        }
}
//违章成了统计
func (e *Cars)GetIllegalCount() int{
        filter := make(map[string]interface{})
        filter["user_id"] = e.UserId
        filter["illegal_status"] = 0
        filter["status"] = 0
        count := 0
        err := orm.Eloquent.
                Table(e.TableName()).
                Where(filter).
                Count(&count).Error
        if nil != err{
                return 0
        }else{
                return count
        }
}