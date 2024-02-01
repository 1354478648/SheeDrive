package car_detail

import (
	"SheeDrive/api/pagination"

	"github.com/gogf/gf/v2/frame/g"
)

// 车辆信息查询
type CarDetailGetListReq struct {
	g.Meta `path:"/list" method:"get"`
	pagination.CommonPaginationReq
	// 关键字查询可选字段
	// 校验规则date-format判断日期是否为指定的日期格式
	Year     string `p:"year" v:"date-format:Y#请输入正确格式的年份" dc:"年份"`
	Brand    string `p:"brand"  dc:"品牌"`
	Model    string `p:"model"  dc:"型号"`
	Category string `p:"category"  dc:"类型"`
	// 校验规则min表示最小值，integer表示正负整数，lt表示比指定字段小
	LowPrice  *int64 `p:"lowPrice" v:"min:0|integer|lt:HighPrice#最低价格不低于0|价格必须是整数|请注意价格区间的前后顺序" dc:"最低价格"`
	HighPrice *int64 `p:"highPrice" v:"min:1000|integer|gt:LowPrice#最低价格不低于1000|价格必须是整数|请注意价格区间的前后顺序" dc:"最高价格"`
}

type CarDetailGetListRes struct {
	pagination.CommonPaginationRes
}
