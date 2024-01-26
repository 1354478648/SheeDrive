// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Address is the golang structure of table address for DAO operations like Where/Data.
type Address struct {
	g.Meta         `orm:"table:address, do:true"`
	Id             interface{} // 主键ID
	BelongId       interface{} // 所属ID
	BelongCategory interface{} // 所属分类 1:经销商,2:用户
	Latitude       interface{} // 纬度
	Longitude      interface{} // 经度
	Accuracy       interface{} // 精度
	Country        interface{} // 国家
	Province       interface{} // 省
	City           interface{} // 市
	District       interface{} // 区
	Street         interface{} // 街道
	StreetNumber   interface{} // 门牌号
	PoiName        interface{} // POI信息
	PostalCode     interface{} // 邮政编码
	CityCode       interface{} // 城市代码
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	DeleteTime     *gtime.Time // 删除时间
	CreateUser     interface{} // 创建人
	UpdateUser     interface{} // 修改人
	DeleteUser     interface{} // 删除人
}
