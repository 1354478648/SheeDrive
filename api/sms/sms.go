package sms

import "github.com/gogf/gf/v2/frame/g"

// 发送手机验证码
type SmsSendCodeReq struct {
	g.Meta `path:"/send" method:"post"`
	Phone  string `P:"phone" v:"required|phone#请输入手机号码|请输入正确的手机号码格式"`
}

// 模拟发送验证码，现实不应该返回验证码
type SmsSendCodeRes struct {
	Code int `json:"code" dc:"手机验证码"`
}
