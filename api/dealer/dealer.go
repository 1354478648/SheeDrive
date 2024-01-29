package dealer

import (
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
