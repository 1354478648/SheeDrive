package api

import (
	"SheeDrive/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 管理员登录
type AdminLoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Username string `p:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `p:"password" v:"required#请输入密码" dc:"密码"`
}

type AdminLoginRes struct {
	Token string        `json:"token" dc:"验证token"`
	Admin *entity.Admin `json:"admin_info"`
}

// 管理员列表分页与关键字查询
type AdminGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	CommonPaginationReq
	// 关键字查询可选字段
	Username string `p:"username" dc:"用户名"`
	Name     string `p:"name" dc:"姓名"`
	// 校验规则required-with表示当指定字段有值时，该字段也必须有值
	BeforeDate *gtime.Time `p:"before_date" v:"required-with:AfterDate|datetime#请输入完整日期|请输入正确的日期格式" dc:"前时间"`
	AfterDate  *gtime.Time `p:"after_date" v:"required-with:BeforeDate|datetime#请输入完整日期|请输入正确的日期格式" dc:"后时间"`
}

type AdminGetListRes struct {
	CommonPaginationRes
}
