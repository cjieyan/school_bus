package models

import (
	"encoding/json"
	"fmt"
	orm "go-admin/global"
	"go-admin/tools"
	"io/ioutil"
	"log"
	"net/http"
)

type SchSites struct {
	Id        int     `json:"id" gorm:"type:int(11);primary_key"`  //
	LineId    int     `json:"lineId" gorm:"type:int(11);"`         // 线路id
	Name      string  `json:"name" gorm:"type:varchar(100);"`      // 名称
	Purpose   string  `json:"purpose" gorm:"type:int(11);"`        // 用途
	Sort      string  `json:"sort" gorm:"type:int(11);"`           // 排序
	Prop      int     `json:"prop" gorm:"type:int(11);"`           // 站点属性
	ArriveAt  string  `json:"arriveAt" gorm:"type:varchar(50);"`   // 到达时间
	Remark    string  `json:"remark" gorm:"type:varchar(200);"`    // 备注
	Longitude float64 `json:"longitude" gorm:"type:varchar(100);"` // 经度
	Latitude  float64 `json:"latitude" gorm:"type:varchar(100);"`  // 维度
	QqLng     float64 `json:"qq_lng" gorm:"type:varchar(100)"`     // qq经度
	QqLat     float64 `json:"qq_lat" gorm:"type:varchar(100)"`     // qq纬度
	Address   string  `json:"address" gorm:"type:varchar(500);"`   //地址
	Picture   string  `json:"picture" gorm:"type:text;"`   // 图片
	IsDelete  int     `json:"isDelete" gorm:"type:tinyint(4);"`    // 0未删除 1已删除
	DataScope string  `json:"dataScope" gorm:"-"`
	Params    string  `json:"params" gorm:"-"`
	LineName		string `json:"lineName" gorm:"-"`
	PropName string    `json:"propName" gorm:"-"`
	BaseModel
}
type SchSitesQqMapData struct {
	Status    int                         `json:"status"`
	Message   string                      `json:"message"`
	Locations []SchSitesQqMapDataLocation `json:"locations"`
}
type SchSitesQqMapDataLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (SchSites) TableName() string {
	return "sch_sites"
}

// 创建SchSites
func (e *SchSites) Create() (SchSites, error) {
	var doc SchSites
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取SchSites
func (e *SchSites) Get() (SchSites, error) {
	var doc SchSites
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Purpose != "" {
		table = table.Where("purpose = ?", e.Purpose)
	}

	if e.Sort != "" {
		table = table.Where("sort = ?", e.Sort)
	}

	if e.Prop != 0 {
		table = table.Where("prop = ?", e.Prop)
	}

	if e.ArriveAt != "" {
		table = table.Where("arrive_at = ?", e.ArriveAt)
	}

	if e.Remark != "" {
		table = table.Where("remark = ?", e.Remark)
	}

	if e.Longitude != 0.0 {
		table = table.Where("longitude = ?", e.Longitude)
	}

	if e.Latitude != 0.0 {
		table = table.Where("latitude = ?", e.Latitude)
	}

	if e.Picture != "" {
		table = table.Where("picture = ?", e.Picture)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SchSites
func (e *SchSites) GetAll() ([]SchSites, error) {
	var doc []SchSites
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Purpose != "" {
		table = table.Where("purpose = ?", e.Purpose)
	}

	if e.Sort != "" {
		table = table.Where("sort = ?", e.Sort)
	}

	if e.Prop != 0 {
		table = table.Where("prop = ?", e.Prop)
	}
	table = table.Order("sort ASC")

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取SchSites带分页
func (e *SchSites) GetPage(pageSize int, pageIndex int) ([]SchSites, int, error) {
	var doc []SchSites

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Id != 0 {
		table = table.Where("id = ?", e.Id)
	}

	if e.LineId != 0 {
		table = table.Where("line_id = ?", e.LineId)
	}

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	if e.Purpose != "" {
		table = table.Where("purpose = ?", e.Purpose)
	}

	if e.Sort != "" {
		table = table.Where("sort = ?", e.Sort)
	}

	if e.Prop != 0 {
		table = table.Where("prop = ?", e.Prop)
	}

	if e.ArriveAt != "" {
		table = table.Where("arrive_at = ?", e.ArriveAt)
	}

	if e.Remark != "" {
		table = table.Where("remark = ?", e.Remark)
	}

	if e.Longitude != 0.0 {
		table = table.Where("longitude = ?", e.Longitude)
	}

	if e.Latitude != 0.0 {
		table = table.Where("latitude = ?", e.Latitude)
	}

	if e.Picture != "" {
		table = table.Where("picture = ?", e.Picture)
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

// 更新SchSites
func (e *SchSites) Update(id int) (update SchSites, err error) {
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

// 删除SchSites
func (e *SchSites) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&SchSites{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *SchSites) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&SchSites{}).Error; err != nil {
		return
	}
	Result = true
	return
}
// qq地图接口转换
func (e SchSites)MapBd2qq(lon, lat float64)(qqLng , qqLat float64, err error){
	locations := fmt.Sprintf("%v,%v", lat, lon)
	key := "URRBZ-64IW3-JYN3S-YNY5B-AYBKS-STBZH"
	url := "https://apis.map.qq.com/ws/coord/v1/translate?locations=" + locations + "&type=3&key=" + key + "&output=json"

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}
	fmt.Println("locations url.....:", url)

	var qqLngLat SchSitesQqMapData

	err = json.Unmarshal(content, &qqLngLat)
	if nil == err && 0 == qqLngLat.Status && len(qqLngLat.Locations) > 0 {
		ret := qqLngLat.Locations[0]
		qqLat = ret.Lat
		qqLng = ret.Lng
	}
	fmt.Println("qqLngLat......", qqLat, qqLng)
	return
}