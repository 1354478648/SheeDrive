package mobile

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"

	"github.com/gogf/gf/errors/gerror"
)

type iMobile struct {
}

func New() *iMobile {
	return &iMobile{}
}

func init() {
	service.RegisterMobile(New())
}

// GetSwiper implements service.IMobile.
func (*iMobile) GetSwiper(ctx context.Context) (out *model.MobileGetSwiperOutput, err error) {
	out = &model.MobileGetSwiperOutput{}

	err = dao.Swiper.Ctx(ctx).WithAll().Scan(&out.Items)
	if err != nil {
		return nil, gerror.New("未找到轮播图数据")
	}
	if err != nil {
		return nil, gerror.New("未找到轮播图数据")
	}
	return
}
