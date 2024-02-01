package dealer

import (
	"SheeDrive/internal/consts"
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

	// 构造动态SQL语句
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

// GetById implements service.IDealer.
func (*iDealer) GetById(ctx context.Context, in model.DealerGetByIdInput) (out *model.DealerGetByIdOutput, err error) {
	// 实例化响应结构体
	out = &model.DealerGetByIdOutput{}

	err = dao.Dealer.Ctx(ctx).WithAll().Where(dao.Dealer.Columns().Id, in.Id).Scan(&out.DealerInfoBase)
	if err != nil {
		return nil, gerror.New("该经销商不存在")
	}

	return
}

// Add implements service.IDealer.
func (*iDealer) Add(ctx context.Context, in model.DealerAddInput) (out *model.DealerAddOutput, err error) {
	// 实例化响应结构体
	out = &model.DealerAddOutput{}

	// 执行添加经销商操作
	id, err := dao.Dealer.Ctx(ctx).Data(do.Dealer{
		Name:         in.Name,
		Username:     in.Username,
		Password:     utility.EncryptPassword(consts.DefaultPassword),
		Avatar:       "",
		Phone:        in.Phone,
		DescribeInfo: in.DescribeInfo,
		Status:       1,
	}).InsertAndGetId()
	if err != nil {
		return out, gerror.New("用户名已被占用")
	}
	out.Id = id

	geocode, err := utility.Geocoding(in.DetailAddress, in.City)
	if err != nil {
		return out, err
	}

	// 执行添加经销商地址操作
	_, err = dao.Address.Ctx(ctx).Data(do.Address{
		Id:             utility.GenSnowFlakeId(),
		BelongId:       id,
		BelongCategory: 1,
		LngLat:         geocode.Location,
		Province:       in.Province,
		City:           in.City,
		District:       in.District,
		Detail:         in.DetailAddress,
	}).Insert()
	if err != nil {
		return out, gerror.New("经销商地址添加失败")
	}

	return
}
