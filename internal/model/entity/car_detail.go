// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDetail is the golang structure for table car_detail.
type CarDetail struct {
	Id           int64       `json:"id"           ` // 主键ID
	Year         string      `json:"year"         ` // 年份
	Brand        string      `json:"brand"        ` // 品牌
	Model        string      `json:"model"        ` // 型号
	Version      string      `json:"version"      ` // 版本
	Image        string      `json:"image"        ` // 图片
	Category     int         `json:"category"     ` // 类型 0:其他, 1:轿车, 2:SUV, 3:MPV, 4:卡车, 5:跑车
	Price        int64       `json:"price"        ` // 指导价
	Type         int         `json:"type"         ` // 类型 0:其他, 1:纯电动, 2:插电混动, 3:增程, 4:汽油, 5:汽油+48V轻混系统, 6:油电混动, 7:柴油
	Seats        int         `json:"seats"        ` // 座位数 0:7座以上 1:1座, 2:2座, 4:4座, 5:5座, 6:6座, 7:7座
	DescribeInfo string      `json:"describeInfo" ` // 描述信息
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
	DeleteTime   *gtime.Time `json:"deleteTime"   ` // 删除时间
}
