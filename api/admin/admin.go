package admin

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

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
	Token     string              `json:"token" dc:"验证token"`
	AdminInfo model.AdminInfoBase `json:"admin_info" dc:"管理员信息"`
}

// 管理员列表分页与关键字查询
type AdminGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	Username string `p:"username" dc:"用户名"`
	Name     string `p:"name" dc:"姓名"`
	// 校验规则required-with表示当指定字段有值时，该字段也必须有值
	// 校验规则before-equal（或after-equal）表示判断该日期是否与指定日期相等或在指定日期之前（或之后）
	BeforeDate *gtime.Time `p:"before_date" v:"required-with:AfterDate|datetime|before-equal:AfterDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"前时间"`
	AfterDate  *gtime.Time `p:"after_date" v:"required-with:BeforeDate|datetime|after-equal:BeforeDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"后时间"`
}

type AdminGetListRes struct {
	pagination.CommonPaginationRes
}

// 通过Id查询管理员
type AdminGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type AdminGetByIdRes struct {
	AdminInfo model.AdminInfoBase `json:"admin_info" dc:"管理员信息"`
}

// 添加管理员
type AdminAddReq struct {
	g.Meta `path:"/add" method:"post"`
	Name   string `p:"name" v:"required#请输入姓名" dc:"姓名"`
	// 校验规则passport表示通用帐号规则（字母开头，只能包含字母、数字和下划线，长度在6~18之间）
	Username string `p:"username" v:"required|passport#请输入用户名|请输入正确的用户名格式（字母开头，只能包含字母、数字和下划线，长度在6~18之间）"`
	Phone    string `P:"phone" v:"required|phone#请输入手机号码|请输入正确的手机号码格式"`
}

type AdminAddRes struct {
	Id int64 `json:"id" dc:"主键id"`
}

// 修改管理员
type AdminUpdateReq struct {
	g.Meta `path:"/update" method:"put"`
	Id     int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Name   string `p:"name" v:"required#请输入姓名" dc:"姓名"`
	// 校验规则passport表示通用帐号规则（字母开头，只能包含字母、数字和下划线，长度在6~18之间）
	Username string `p:"username" v:"required|passport#请输入用户名|请输入正确的用户名格式（字母开头，只能包含字母、数字和下划线，长度在6~18之间）"`
	Phone    string `P:"phone" v:"required|phone#请输入手机号码|请输入正确的手机号码格式"`
}

type AdminUpdateRes struct{}

// 删除管理员
type AdminDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type AdminDeleteRes struct{}

// 修改管理员状态
type AdminUpdateStatusReq struct {
	g.Meta `path:"/updateStatus" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type AdminUpdateStatusRes struct{}

// 修改管理员密码
type AdminUpdatePasswordReq struct {
	g.Meta          `path:"/updatePassword" method:"put"`
	Id              int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Password        string `p:"password" v:"required|different:NewPassword#请输入原密码|原密码不可与新密码一致" dc:"原密码"`
	NewPassword     string `p:"newPassword" v:"required|password|same:ConfirmPassword#请输入新密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"新密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:NewPassword#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type AdminUpdatePasswordRes struct{}

// 重置管理员密码
type AdminResetPasswordReq struct {
	g.Meta `path:"/resetPassword" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type AdminResetPasswordRes struct{}

// 修改管理员头像
type AdminUpdateAvatarReq struct {
	g.Meta `path:"/updateAvatar" method:"put"`
	Id     int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Url    string `p:"url" v:"required|url#请上传头像文件|头像未找到" dc:"头像文件路径"`
}

type AdminUpdateAvatarRes struct{}

// 通过手机号修改密码
type AdminUpdatePasswordByPhoneReq struct {
	g.Meta          `path:"/updatePasswordByPhone" method:"put"`
	Phone           string `p:"phone" v:"required|phone#请输入手机号|手机号格式不正确" dc:"手机号"`
	Code            int    `p:"code" v:"required#请输入验证码" dc:"验证码"`
	NewPassword     string `p:"newPassword" v:"required|password|same:ConfirmPassword#请输入新密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"新密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:NewPassword#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type AdminUpdatePasswordByPhoneRes struct{}
