package user

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 用户登录
type UserLoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	UserName string `p:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `p:"password" v:"required#请输入密码" dc:"密码"`
}

type UserLoginRes struct {
	Token    string             `json:"token" dc:"验证Token"`
	UserInfo model.UserInfoBase `json:"user_info" dc:"用户信息"`
}

// 用户通过手机号登录
type UserLoginByPhoneReq struct {
	g.Meta `path:"/loginByPhone" method:"post"`
	Phone  string `p:"phone" v:"required|phone#请输入您的联系电话|请输入正确格式的联系电话" dc:"联系电话"`
	Code   int    `p:"code" v:"required#请输入验证码" dc:"验证码"`
}

type UserLoginByPhoneRes struct {
	Token    string             `json:"token" dc:"验证Token"`
	UserInfo model.UserInfoBase `json:"user_info" dc:"用户信息"`
}

// 用户注册
type UserRegisterReq struct {
	g.Meta          `path:"/register" method:"post"`
	LastName        string `p:"last_name" v:"required#请输入您的姓" dc:"姓"`
	FirstName       string `p:"first_name" v:"required#请输入您的名" dc:"名"`
	IdNumber        string `p:"id_number" v:"required|resident-id#请输入您的身份证号|请输入正确格式的身份证号" dc:"身份证号"`
	Phone           string `p:"phone" v:"required|phone#请输入您的联系电话|请输入正确格式的联系电话" dc:"联系电话"`
	Password        string `p:"password" v:"required|password|same:ConfirmPassword#请输入密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:Password#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type UserRegisterRes struct {
	UserInfo model.UserInfoBase `json:"user_info" dc:"用户信息"`
}

// 用户列表分页与关键字查询
type UserGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	Username   string      `p:"username" dc:"用户名"`
	Name       string      `p:"name" dc:"姓名"`
	Status     int         `p:"status" d:"-1" dc:"状态"`
	BeforeDate *gtime.Time `p:"before_date" v:"required-with:AfterDate|datetime|before-equal:AfterDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"前时间"`
	AfterDate  *gtime.Time `p:"after_date" v:"required-with:BeforeDate|datetime|after-equal:BeforeDate#请输入完整日期|请输入正确的日期格式|请注意前后日期顺序" dc:"后时间"`
}

type UserGetListRes struct {
	pagination.CommonPaginationRes
}

// 用户信息通过Id获取
type UserGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserGetByIdRes struct {
	UserInfo model.UserInfoBase `json:"user_info" dc:"用户信息"`
}

// 删除用户
type UserDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserDeleteRes struct{}

// 修改用户状态
type UserUpdateStatusReq struct {
	g.Meta `path:"/updateStatus" method:"put"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserUpdateStatusRes struct{}

// 修改用户密码
type UserUpdatePasswordReq struct {
	g.Meta          `path:"/updatePassword" method:"put"`
	Id              int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Password        string `p:"password" v:"required|different:NewPassword#请输入原密码|原密码不可与新密码一致" dc:"原密码"`
	NewPassword     string `p:"newPassword" v:"required|password|same:ConfirmPassword#请输入新密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"新密码"`
	ConfirmPassword string `p:"confirmPassword" v:"required|password|same:NewPassword#请输入确认密码|密码格式不正确（任意可见字符，长度在6~18之间）|两次密码输入不一致" dc:"确认密码"`
}

type UserUpdatePasswordRes struct{}

// 修改用户头像
type UserUpdateAvatarReq struct {
	g.Meta `path:"/updateAvatar" method:"put"`
	Id     int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Url    string `p:"url" v:"required|url#请上传头像文件|头像未找到" dc:"头像文件路径"`
}

type UserUpdateAvatarRes struct{}
