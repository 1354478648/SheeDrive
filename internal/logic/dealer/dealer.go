package dealer

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"

	"github.com/gogf/gf/errors/gerror"
)

type iDealer struct{}

func New() *iDealer {
	return &iDealer{}
}

func init() {
	service.RegisterDealer(New())
}

// Login implements service.IDealer.
func (*iDealer) Login(ctx context.Context, in model.DealerLoginInput) (out *model.DealerLoginOutput, err error) {
	// 实例化响应结构体
	out = &model.DealerLoginOutput{}

	err = dao.Dealer.Ctx(ctx).Where(do.Dealer{
		Username: in.Username,
		Password: utility.EncryptPassword(in.Password),
	}).Scan(&out.DealerInfoBase)
	if err != nil {
		return nil, gerror.New("用户名或密码不正确")
	}

	// 判断经销商状态是否被禁用
	if out.DealerInfoBase.Status == 0 {
		return nil, gerror.New("该经销商账号已被禁用")
	}
	return
}
