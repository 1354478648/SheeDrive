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

// GetList implements service.IDealer.
func (*iDealer) GetList(ctx context.Context, in model.DealerGetListInput) (out *model.DealerGetListOutput, err error) {
	// 实例化响应结构体
	out = &model.DealerGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	// 获取*gdb.Model对象
	var (
		md = dao.Dealer.Ctx(ctx)
	)
	// 关联查询
	md = md.WithAll()

	if in.Name != "" {
		md = md.WhereLike(dao.Dealer.Columns().Name, "%"+in.Name+"%")
	}

	// 设置排序：更新时间降序;设置分页查询
	md = md.OrderDesc(dao.Dealer.Columns().UpdateTime).Page(in.Page, in.PageSize)

	// 判断当前页的数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 将查询结果赋值给响应结构体
	if err := md.Scan(&out.Items); err != nil {
		return out, err
	}

	return
}
