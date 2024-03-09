package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderInfoBase struct {
	Id          int64       `json:"id"          ` // 主键ID
	UserId      int64       `json:"userId"      ` // 用户ID
	DealerId    int64       `json:"dealerId"    ` // 经销商ID
	CarId       int64       `json:"carId"       ` // 车辆ID
	AddrId      int64       `json:"addrId"      ` // 用户地址ID
	Status      int         `json:"status"      ` // 订单状态 -1:异常,0:取消,1:未确认,2:已确认,3:签署协议,4:试驾中,5:试驾结束,6:待评价,7:已评价
	OrderTime   *gtime.Time `json:"orderTime"   ` // 预定时间
	ConfirmTime *gtime.Time `json:"confirmTime" ` // 确认时间
	SignTime    *gtime.Time `json:"signTime"    ` // 签署协议时间
	StartTime   *gtime.Time `json:"startTime"   ` // 试驾开始时间
	EndTime     *gtime.Time `json:"endTime"     ` // 试驾结束时间
	CommentTime *gtime.Time `json:"commentTime" ` // 评价时间
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  ` // 更新时间
	DeleteTime  *gtime.Time `json:"deleteTime"  ` // 删除时间

	UserInfo      *UserInfoBase      `orm:"with:id=user_id" json:"userInfo"`
	DealerInfo    *DealerInfoBase    `orm:"with:id=dealer_id" json:"dealerInfo"`
	CarDetailInfo *CarDetailInfoBase `orm:"with:id=car_id" json:"carDetailInfo"`
	Address       *AddressInfoBase   `orm:"with:id=addr_id" json:"address"`
}

type OrderGetListInput struct {
	Page       int
	PageSize   int
	UserName   string
	DealerName string
	CarName    string
	Status     int
	OrderDate  *gtime.Time
}

type OrderGetListOutput struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	Total    int             `json:"total"`
	Items    []OrderInfoBase `json:"items"`
}

type OrderGetByIdInput struct {
	Id int64
}

type OrderGetByIdOutput struct {
	OrderInfoBase
}

type OrderAddInput struct {
	UserId    int64
	DealerId  int64
	CarId     int64
	AddrId    int64
	OrderTime *gtime.Time
}

type OrderAddOutput struct {
	OrderInfoBase
}

type OrderDeleteInput struct {
	Id int64
}

type OrderUpdateInput struct {
	Id int64
}
