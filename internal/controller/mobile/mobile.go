package mobile

import (
	apiMobile "SheeDrive/api/mobile"
	"SheeDrive/internal/service"
	"context"
)

var MobileController = &cMobile{}

type cMobile struct{}

// 获取轮播图
func (c *cMobile) MobileGetSwiper(ctx context.Context, req *apiMobile.MobileGetSwiperReq) (res *apiMobile.MobileGetSwiperRes, err error) {
	out, err := service.Mobile().GetSwiper(ctx)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	res = &apiMobile.MobileGetSwiperRes{
		SwiperInfoList: out.Items,
	}
	return
}
