package model

import (
	"SheeDrive/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 经销商信息基类
type DealerInfoBase struct {
	// 指定映射关系，用于其他表关联查询的时候找到这个结构体对应的数据表
	g.Meta   `orm:"table:dealer"`
	Id       int64  `json:"id"        ` // 主键ID
	Name     string `json:"name"      ` // 名称
	Username string `json:"username"  ` // 用户名
	// Password   string      `json:"password"` // 密码
	Avatar       string      `json:"avatar"       ` // 头像
	Phone        string      `json:"phone"        ` // 手机号
	DescribeInfo string      `json:"describeInfo" ` //详细信息
	Status       int         `json:"status"       ` // 状态 0:禁用, 1:正常
	Token        string      `json:"token"      `   // token
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
	DeleteTime   *gtime.Time `json:"deleteTime"   ` // 删除时间

	Address *do.Address `orm:"with:belong_id=id" json:"address"`
}

// 经销商添加修改基类
type DealerAddUpdateBase struct {
	Name         string `json:"name"`         //名称
	Username     string `json:"username"`     //用户名
	Phone        string `json:"phone"`        //手机号
	DescribeInfo string `json:"describeInfo"` //详细信息
	Province     string `json:"province"`     //省
	City         string `json:"city"`         //市
	District     string `json:"district"`     //区
	Detail       string `json:"detail"`       //详细地址
}

type DealerLoginInput struct {
	Username string
	Password string
}

type DealerLoginOutput struct {
	Token string `json:token`
	DealerInfoBase
}

type DealerGetListInput struct {
	Page     int
	PageSize int
	Name     string
}

type DealerGetListOutput struct {
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
	Total    int              `json:"total"`
	Items    []DealerInfoBase `json:"items"`
}

type DealerGetByIdInput struct {
	Id int64
}

type DealerGetByIdOutput struct {
	DealerInfoBase
}

type DealerAddInput struct {
	DealerAddUpdateBase
}

type DealerAddOutput struct {
	Id int64
}

type DealerUpdateInput struct {
	Id int64
	DealerAddUpdateBase
}

type DealerDeleteInput struct {
	Id int64
}

type DealerUpdateStatusInput struct {
	Id int64
}

type DealerUpdatePasswordInput struct {
	Id          int64
	OldPassword string
	NewPassword string
}

type DealerResetPasswordInput struct {
	Id int64
}

type DealerUpdateAvatarInput struct {
	Id  int64
	Url string
}
