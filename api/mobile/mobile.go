package mobile

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 获取轮播图
type MobileGetSwiperReq struct {
	g.Meta `path:"/getSwiper" method:"get" summary:"获取轮播图"`
}

type MobileGetSwiperRes struct {
	SwiperInfoList interface{} `json:"swiperInfoList" dc:"轮播图信息列表"`
}
