package model

import "github.com/gogf/gf/v2/os/gtime"

// 轮播图信息基类
type SwiperInfoBase struct {
	Id           int64       `json:"id"           ` // 主键ID
	CarId        int64       `json:"carId"        ` // 车辆ID
	ImageUrl     string      `json:"imageUrl"     ` // 图片地址
	DescribeInfo string      `json:"describeInfo" ` // 描述信息
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	DeleteTime   *gtime.Time `json:"deleteTime"   ` // 删除时间

	CarDetailInfo *CarDetailInfoBase `orm:"with:id=car_id" json:"carDetailInfo"`
}

// 获取轮播图
type MobileGetSwiperOutput struct {
	Items []SwiperInfoBase `json:"items"`
}
