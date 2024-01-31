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
