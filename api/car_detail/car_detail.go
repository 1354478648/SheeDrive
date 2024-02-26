package car_detail

import (
	"SheeDrive/api/pagination"
	"SheeDrive/internal/model/entity"

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
	Category string `p:"category"  dc:"分类"`
	Type     string `p:"type" dc:"驱动类型"`
	// 校验规则min表示最小值，integer表示正负整数
	// 有两种情况：价格区间可以只输入一边和如果输入两边还需要判断两边大小是否正确
	// 但如果在这里做两边大小的校验，那么只输入一边的情况也会被做校验，因此需要在业务层做额外校验
	LowPrice  int64 `p:"lowPrice" v:"min:0|integer#最低价格不低于0|价格必须是整数" dc:"最低价格"`
	HighPrice int64 `p:"highPrice" v:"min:1000|integer#最高价格不低于1000|价格必须是整数" dc:"最高价格"`
}

type CarDetailGetListRes struct {
	pagination.CommonPaginationRes
}

// 通过Id查询汽车信息
type CarDetailGetByIdReq struct {
	g.Meta `path:"/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入id" dc:"id"`
}

type CarDetailGetByIdRes struct {
	CarDetail entity.CarDetail `json:"car_detail" dc:"汽车信息"`
}

// 添加汽车信息
type CarDetailAddReq struct {
	g.Meta       `path:"/add" method:"post"`
	Year         string `p:"year" v:"required|date-format:Y#请输入年份|请输入正确格式的年份"  dc:"年份"`
	Brand        string `p:"brand" v:"required#请输入品牌"  dc:"品牌"`
	Model        string `p:"model" v:"required#请输入型号"  dc:"型号"`
	Version      string `p:"version" v:"required#请输入版本"  dc:"版本"`
	Image        string `p:"image" v:"required|url#请上传图片文件|请上传正确的文件路径" dc:"图片"`
	Category     int    `p:"category" v:"required#请输入类型"  dc:"类型"`
	Price        int64  `p:"price" v:"required|min:0|integer#请输入价格|价格必须是正整数|价格必须是正整数" dc:"价格"`
	Type         int    `p:"type" v:"required#请输入驱动类型" dc:"驱动类型"`
	Seats        int    `p:"seats" v:"required#请输入座位数" dc:"座位数"`
	DescribeInfo string `p:"describeInfo" v:"required#请输入描述信息" dc:"描述信息"`
}

type CarDetailAddRes struct {
	Id int64
}

// 修改汽车信息
type CarDetailUpdateReq struct {
	g.Meta       `path:"/update" method:"put"`
	Id           int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Year         string `p:"year" v:"required|date-format:Y#请输入年份|请输入正确格式的年份"  dc:"年份"`
	Brand        string `p:"brand" v:"required#请输入品牌"  dc:"品牌"`
	Model        string `p:"model" v:"required#请输入型号"  dc:"型号"`
	Version      string `p:"version" v:"required#请输入版本"  dc:"版本"`
	Image        string `p:"image" v:"required|url#请上传图片文件|请上传正确的文件路径" dc:"图片"`
	Category     int    `p:"category" v:"required#请输入类型"  dc:"类型"`
	Price        int64  `p:"price" v:"required|min:0|integer#请输入价格|价格必须是正整数|价格必须是正整数" dc:"价格"`
	Type         int    `p:"type" v:"required#请输入驱动类型" dc:"驱动类型"`
	Seats        int    `p:"seats" v:"required#请输入座位数" dc:"座位数"`
	DescribeInfo string `p:"describeInfo" v:"required#请输入描述信息" dc:"描述信息"`
}

type CarDetailUpdateRes struct{}

// 汽车信息删除
type CarDetailDeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type CarDetailDeleteRes struct{}
