package sms

import (
	apiSms "SheeDrive/api/sms"
	"SheeDrive/utility"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var SmsController = &cSms{}

type cSms struct{}

// 发送手机验证码
func (c *cSms) SmsSendCode(ctx context.Context, req *apiSms.SmsSendCodeReq) (res *apiSms.SmsSendCodeRes, err error) {
	res = &apiSms.SmsSendCodeRes{
		Code: utility.GenRandomCode(),
	}

	err = g.Redis().SetEX(ctx, req.Phone, res.Code, 60)
	if err != nil {
		return nil, gerror.New("验证码保存失败")
	}

	fmt.Printf("--- 这是模拟发送验证码：发送给手机号 %v 的验证码是 %v", req.Phone, res.Code)

	return
}
