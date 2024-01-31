package dealer

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 经销商登录
type DealerLoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Username string `p:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `p:"password" v:"required#请输入密码" dc:"密码"`
}

type DealerLoginRes struct {
	Token      string               `json:"token" dc:"验证token"`
	DealerInfo model.DealerInfoBase `json:"dealer_info" dc:"经销商信息"`
}

// 经销商分页与关键字查询
type DealerGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	Name string `p:"name" dc:"名称"`
}

type DealerGetListRes struct {
	pagination.CommonPaginationRes
}

// 通过Id查询经销商
type DealerGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type DealerGetByIdRes struct {
	DealerInfo model.DealerInfoBase `json:"dealer_info" dc:"经销商信息"`
}

// 添加经销商
type DealerAddReq struct {
	g.Meta        `path:"/add" method:"post"`
	Name          string `p:"name" v:"required#请输入名称" dc:"名称"`
	Username      string `p:"username" v:"required|passport#请输入用户名|请输入正确的用户名格式（字母开头，只能包含字母、数字和下划线，长度在6~18之间）"`
	Phone         string `P:"phone" v:"required|phone#请输入手机号码|请输入正确的手机号码格式"`
	DescribeInfo  string `p:"describe_info" v:"" dc:"描述信息"`
	Province      string `p:"province" v:"required#请输入省份信息" dc:"省"`
	City          string `p:"city" v:"required#请输入城市信息" dc:"市"`
	District      string `p:"district" v:"required#请输入区县信息" dc:"区县"`
	DetailAddress string `p:"detail_address" v:"required#请输入详细地址" dc:"详细地址"`
}

type DealerAddRes struct {
	Id int64 `json:"id" dc:"主键id"`
}
