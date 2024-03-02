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

// 修改经销商
type DealerUpdateReq struct {
	g.Meta        `path:"/update" method:"put"`
	Id            int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Name          string `p:"name" v:"required#请输入名称" dc:"名称"`
	Username      string `p:"username" v:"required|passport#请输入用户名|请输入正确的用户名格式（字母开头，只能包含字母、数字和下划线，长度在6~18之间）"`
	Phone         string `P:"phone" v:"required|phone#请输入手机号码|请输入正确的手机号码格式"`
	DescribeInfo  string `p:"describe_info" v:"" dc:"描述信息"`
	Province      string `p:"province" v:"required#请输入省份信息" dc:"省"`
	City          string `p:"city" v:"required#请输入城市信息" dc:"市"`
	District      string `p:"district" v:"required#请输入区县信息" dc:"区县"`
	DetailAddress string `p:"detail_address" v:"required#请输入详细地址" dc:"详细地址"`
}

type DealerUpdateRes struct{}

// 删除经销商
type DealerDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type DealerDeleteRes struct{}

// 修改经销商状态
type DealerUpdateStatusReq struct {
	g.Meta `path:"/updateStatus" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type DealerUpdateStatusRes struct{}

// 修改经销商密码
type DealerUpdatePasswordReq struct {
	g.Meta          `path:"/updatePassword" method:"put"`
	Id              int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Password        string `p:"password" v:"required|different:NewPassword#请输入原密码|原密码不可与新密码一致" dc:"原密码"`
	NewPassword     string `p:"newPassword" v:"required|password|same:ConfirmPassword#请输入新密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"新密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:NewPassword#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type DealerUpdatePasswordRes struct{}

// 重置经销商密码
type DealerResetPasswordReq struct {
	g.Meta `path:"/resetPassword" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type DealerResetPasswordRes struct{}

// 修改经销商头像
type DealerUpdateAvatarReq struct {
	g.Meta `path:"/updateAvatar" method:"put"`
	Id     int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Url    string `p:"url" v:"required|url#请上传头像文件|头像未找到" dc:"头像文件路径"`
}

type DealerUpdateAvatarRes struct{}

// 通过手机号修改密码
type DealerUpdatePasswordByPhoneReq struct {
	g.Meta          `path:"/updatePasswordByPhone" method:"put"`
	Phone           string `p:"phone" v:"required|phone#请输入手机号|手机号格式不正确" dc:"手机号"`
	Code            int    `p:"code" v:"required#请输入验证码" dc:"验证码"`
	NewPassword     string `p:"newPassword" v:"required|password|same:ConfirmPassword#请输入新密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"新密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:NewPassword#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type DealerUpdatePasswordByPhoneRes struct{}
