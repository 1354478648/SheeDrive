package dealer

import (
	"SheeDrive/internal/consts"
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"
	"fmt"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	}).WithAll().Scan(&out.DealerInfoBase)
	if err != nil {
		return nil, gerror.New("用户名或密码不正确")
	}

	// 判断经销商状态是否被禁用
	if out.DealerInfoBase.Status == 0 {
		return nil, gerror.New("该经销商账号已被禁用")
	}

	// 生成token
	out.Token = utility.GenToken(in.Username)
	// 将token保存到redis中
	err = g.Redis().SetEX(ctx, out.Token, out.Token, 86400)
	if err != nil {
		return nil, gerror.New("Token保存失败")
	}

	// 将Token持久化
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, out.DealerInfoBase.Id).Data(
		do.Dealer{Token: out.Token}).Update()
	if err != nil {
		return nil, gerror.New("Token保存失败")
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

	geocode, err := utility.Geocoding(fmt.Sprintf("%s%s%s%s", in.Province, in.City, in.District, in.Detail), in.City)
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
		Detail:         in.Detail,
	}).Insert()
	if err != nil {
		return out, gerror.New("经销商地址添加失败")
	}

	return
}

// Update implements service.IDealer.
func (*iDealer) Update(ctx context.Context, in model.DealerUpdateInput) (err error) {
	// 执行修改经销商操作
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Data(do.Dealer{
		Name:         in.Name,
		Username:     in.Username,
		Phone:        in.Phone,
		DescribeInfo: in.DescribeInfo,
	}).Update()
	if err != nil {
		return gerror.New("用户名已被占用")
	}

	//执行修改经销商地址操作
	geocode, err := utility.Geocoding(fmt.Sprintf("%s%s%s%s", in.Province, in.City, in.District, in.Detail), in.City)
	if err != nil {
		return err
	}
	_, err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().BelongId, in.Id).Data(do.Address{
		LngLat:   geocode.Location,
		Province: in.Province,
		City:     in.City,
		District: in.District,
		Detail:   in.Detail,
	}).Update()
	if err != nil {
		return gerror.New("经销商地址修改失败")
	}

	return
}

// Delete implements service.IDealer.
func (*iDealer) Delete(ctx context.Context, in model.DealerDeleteInput) (err error) {
	// 删除对应id的token
	id := in.Id
	dealerInfo, err := service.Dealer().GetById(ctx, model.DealerGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("未找到该经销商")
	}
	_, err = g.Redis().Del(ctx, dealerInfo.Token)
	if err != nil {
		return gerror.New("token删除失败")
	}

	// 执行删除经销商操作
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除经销商失败")
	}

	// 执行删除经销商地址操作
	_, err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().BelongCategory, 1).Where(dao.Address.Columns().BelongId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除经销商地址失败")
	}

	// 执行删除经销商库存操作
	_, err = dao.Stock.Ctx(ctx).Where(dao.Stock.Columns().DealerId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除经销商库存失败")
	}

	return
}

// UpdateStatus implements service.IDealer.
func (*iDealer) UpdateStatus(ctx context.Context, in model.DealerUpdateStatusInput) (err error) {
	// 获取经销商原状态
	id := in.Id
	dealerInfo, err := service.Dealer().GetById(ctx, model.DealerGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("该经销商不存在")
	}
	// 切换经销商状态
	switch dealerInfo.Status {
	case 0:
		_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, id).Data(do.Dealer{Status: 1}).Update()
		if err != nil {
			return gerror.New("经销商状态切换失败")
		}
	case 1:
		_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, id).Data(do.Dealer{Status: 0}).Update()
		if err != nil {
			return gerror.New("经销商状态切换失败")
		}
	}
	return
}

// UpdatePassword implements service.IDealer.
func (*iDealer) UpdatePassword(ctx context.Context, in model.DealerUpdatePasswordInput) (err error) {
	// 获取原密码
	oldPassword, err := dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Value(dao.Dealer.Columns().Password)
	if err != nil {
		return gerror.New("获取密码失败")
	}

	// 验证原密码
	if utility.EncryptPassword(in.OldPassword) != oldPassword.String() {
		return gerror.New("原密码不正确")
	}

	// 更新密码
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Data(do.Dealer{Password: utility.EncryptPassword(in.NewPassword)}).Update()
	if err != nil {
		return gerror.New("更新密码失败")
	}

	// 删除对应id的token
	id := in.Id
	dealerInfo, err := service.Dealer().GetById(ctx, model.DealerGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("未找到该经销商")
	}
	_, err = g.Redis().Del(ctx, dealerInfo.Token)
	if err != nil {
		return gerror.New("token删除失败")
	}

	return
}

// ResetPassword implements service.IDealer.
func (*iDealer) ResetPassword(ctx context.Context, in model.DealerResetPasswordInput) (err error) {
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Data(do.Dealer{Password: utility.EncryptPassword(consts.DefaultPassword)}).Update()
	if err != nil {
		return gerror.New("重置密码失败")
	}

	// 删除对应id的token
	id := in.Id
	dealerInfo, err := service.Dealer().GetById(ctx, model.DealerGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("未找到该经销商")
	}
	_, err = g.Redis().Del(ctx, dealerInfo.Token)
	if err != nil {
		return gerror.New("token删除失败")
	}

	return
}

// UpdateAvatar implements service.IDealer.
func (*iDealer) UpdateAvatar(ctx context.Context, in model.DealerUpdateAvatarInput) (err error) {
	// 修改头像
	_, err = dao.Dealer.Ctx(ctx).Where(dao.Dealer.Columns().Id, in.Id).Data(do.Dealer{Avatar: in.Url}).Update()
	if err != nil {
		return gerror.New("修改头像失败")
	}
	return
}
