package model

import "github.com/gogf/gf/v2/os/gtime"

type CommentInfoBase struct {
	Id          int64       `json:"id"          ` // 主键ID
	OrderId     int64       `json:"orderId"     ` // 订单ID
	Content     string      `json:"content"     ` // 评价内容
	TotalScore  int         `json:"totalScore"  ` // 总评分 1~5星
	DealerScore int         `json:"dealerScore" ` // 经销商评分 1~5星
	CarScore    int         `json:"carScore"    ` // 汽车评分 1~5星
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	DeleteTime  *gtime.Time `json:"deleteTime"  ` // 删除时间

	OrderInfo *OrderInfoBase `orm:"with:id=order_id" json:"orderInfo"`
}

type CommentGetListInput struct {
	Page       int
	PageSize   int
	OrderId    int64
	DealerName string
	CarName    string
	BeforeDate *gtime.Time
	AfterDate  *gtime.Time
}

type CommentGetListOutput struct {
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	Total    int               `json:"total"`
	Items    []CommentInfoBase `json:"items"`
}

type CommentGetByIdInput struct {
	Id int64
}

type CommentGetByIdOutput struct {
	CommentInfoBase
}

type CommentAddInput struct {
	OrderId     int64
	Content     string
	TotalScore  int
	DealerScore int
	CarScore    int
}

type CommentAddOutput struct {
	Id int64
}

type CommentDeleteInput struct {
	Id int64
}
