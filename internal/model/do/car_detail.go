// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDetail is the golang structure of table car_detail for DAO operations like Where/Data.
type CarDetail struct {
	g.Meta       `orm:"table:car_detail, do:true"`
	Id           interface{} // 主键ID
	Year         interface{} // 年份
	Brand        interface{} // 品牌
	Model        interface{} // 型号
	Version      interface{} // 版本
	Image        interface{} // 图片
	Category     interface{} // 类型 0:其他, 1:轿车, 2:SUV, 3:MPV, 4:卡车, 5:跑车
	Color        interface{} // 颜色
	Price        interface{} // 指导价
	Type         interface{} // 类型 0:其他, 1:纯电动, 2:插电混动, 3:增程, 4:汽油, 5:汽油+48V轻混系统, 6:油电混动, 7:柴油
	Seats        interface{} // 座位数 0:7座以上 1:1座, 2:2座, 4:4座, 5:5座, 6:6座, 7:7座
	DescribeInfo interface{} // 描述信息
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	DeleteTime   *gtime.Time // 删除时间
}
