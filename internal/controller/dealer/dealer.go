package dealer

import (
	apiDealer "SheeDrive/api/dealer"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"
)

var DealerController = &cDealer{}

type cDealer struct{}

// 经销商登录
func (c *cDealer) DealerLogin(ctx context.Context, req *apiDealer.DealerLoginReq) (res *apiDealer.DealerLoginRes, err error) {
	// 调用Service层接口
	dealer, err := service.Dealer().Login(ctx, model.DealerLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res = &apiDealer.DealerLoginRes{
		Token:      utility.GenToken(dealer.Username),
		DealerInfo: dealer.DealerInfoBase,
	}

	return
}
