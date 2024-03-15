package comment

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CommentGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	OrderId    int64       `p:"orderId" dc:"订单id"`
	DealerName string      `p:"dealerName" dc:"经销商名称"`
	CarName    string      `p:"carName" dc:"车型名称"`
	BeforeDate *gtime.Time `p:"before_date" v:"required-with:AfterDate|datetime|before-equal:AfterDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"前时间"`
	AfterDate  *gtime.Time `p:"after_date" v:"required-with:BeforeDate|datetime|after-equal:BeforeDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"后时间"`
}

type CommentGetListRes struct {
	pagination.CommonPaginationRes
}

type CommentGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入ID"`
}

type CommentGetByIdRes struct {
	CommentInfo model.CommentInfoBase `json:"commentInfo" dc:"评论信息"`
}

type CommentAddReq struct {
	g.Meta      `path:"/add" method:"post"`
	OrderId     int64  `p:"orderId" v:"required#请输入订单id" dc:"订单ID"`
	Content     string `p:"content" v:"required#请输入评价内容" dc:"评价内容"`
	TotalScore  int    `p:"totalScore" v:"required|integer|min:1|max:5#请选择总评分|总评分格式不正确|总评分不可低于1星|总评分不可高于5星" dc:"总评分"`
	DealerScore int    `p:"dealerScore" v:"required|integer|min:1|max:5#请选择经销商评分|经销商评分格式不正确|经销商评分不可低于1星|经销商评分不可高于5星" dc:"经销商评分"`
	CarScore    int    `p:"carScore" v:"required|integer|min:1|max:5#请选择汽车评分|汽车评分格式不正确|汽车评分不可低于1星|汽车评分不可高于5星" dc:"汽车评分"`
}

type CommentAddRes struct {
	Id int64 `json:"id" dc:"评论ID"`
}

type CommentDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type CommentDeleteRes struct{}

type CommentGetAvgReq struct {
	g.Meta   `path:"/get/avg" method:"get"`
	DealerId int64 `p:"dealerId" v:"required#请输入经销商ID" dc:"经销商ID"`
}

type CommentGetAvgRes struct {
	Avg float64 `json:"avg" dc:"平均评分"`
}

type CommentGetByOrderIdReq struct {
	g.Meta  `path:"/detailByOrderId" method:"get"`
	OrderId int64 `p:"orderId" v:"required#请输入订单ID"`
}

type CommentGetByOrderIdRes struct {
	CommentInfo model.CommentInfoBase `json:"commentInfo" dc:"评论信息"`
}
