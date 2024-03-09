package address

import (
	"SheeDrive/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 用户地址列表查询
type UserAddressGetListReq struct {
	g.Meta `path:"/user/list" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserAddressGetListRes struct {
	AddressInfoList interface{} `json:"addressInfoList" dc:"用户地址信息列表"`
}

// 用户地址添加
type UserAddressAddReq struct {
	g.Meta        `path:"/user/add" method:"post"`
	Id            int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Province      string `p:"province" v:"required#请输入省份信息" dc:"省"`
	City          string `p:"city" v:"required#请输入城市信息" dc:"市"`
	District      string `p:"district" v:"required#请输入区县信息" dc:"区县"`
	DetailAddress string `p:"detail_address" v:"required#请输入详细地址" dc:"详细地址"`
}

type UserAddressAddRes struct {
	Id string `json:"id" dc:"主键id"`
}

// 通过Id查询用户地址
type UserAddressGetByIdReq struct {
	g.Meta `path:"/user/detail" method:"get"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserAddressGetByIdRes struct {
	AddressInfo model.AddressInfoBase `json:"addressInfo" dc:"用户地址信息"`
}

// 修改用户地址
type UserAddressUpdateReq struct {
	g.Meta        `path:"/user/update" method:"put"`
	Id            int64  `p:"id" v:"required#请输入Id" dc:"id"`
	Province      string `p:"province" v:"required#请输入省份信息" dc:"省"`
	City          string `p:"city" v:"required#请输入城市信息" dc:"市"`
	District      string `p:"district" v:"required#请输入区县信息" dc:"区县"`
	DetailAddress string `p:"detail_address" v:"required#请输入详细地址" dc:"详细地址"`
}

type UserAddressUpdateRes struct{}

// 删除用户地址
type UserAddressDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete"`
	Id     int64 `p:"id" v:"required#请输入Id" dc:"id"`
}

type UserAddressDeleteRes struct{}
