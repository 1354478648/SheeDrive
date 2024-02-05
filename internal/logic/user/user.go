package user

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"

	"github.com/gogf/gf/errors/gerror"
)

type iUser struct{}

func New() *iUser {
	return &iUser{}
}

func init() {
	service.RegisterUser(New())
}

// Login implements service.IUser.
func (*iUser) Login(ctx context.Context, in model.UserLoginInput) (out *model.UserLoginOutput, err error) {
	out = &model.UserLoginOutput{}

	err = dao.User.Ctx(ctx).Where(do.User{
		Username: in.Username,
		Password: utility.EncryptPassword(in.Password),
	}).Scan(&out.UserInfoBase)
	if err != nil {
		return nil, gerror.New("用户名或密码不正确")
	}

	// 判断用户状态是否被禁用
	if out.UserInfoBase.Status == 0 {
		return nil, gerror.New("该用户已被禁用")
	}

	return
}

// Register implements service.IUser.
func (*iUser) Register(ctx context.Context, in model.UserRegisterInput) (out *model.UserRegisterOutput, err error) {
	out = &model.UserRegisterOutput{}

	id := utility.GenSnowFlakeId()
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Id:        id,
		LastName:  in.LastName,
		FirstName: in.FirstName,
		Username:  in.Phone,
		Password:  utility.EncryptPassword(in.Password),
		Phone:     in.Phone,
		IdNumber:  in.IdNumber,
		Sex:       utility.GetGender(in.IdNumber),
		Birthday:  utility.GetBirthDay(in.IdNumber),
		Status:    1,
	}).Insert()
	if err != nil {
		return out, gerror.New("注册失败")
	}

	UserInfo, err := service.User().GetById(ctx, model.UserGetByIdInput{Id: id})
	out.UserInfoBase = UserInfo.UserInfoBase
	return
}

// GetList implements service.IUser.
func (*iUser) GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.UserGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	// 2. 获取*gdb.Model对象
	var (
		md = dao.User.Ctx(ctx)
	)
	// 3. 通过md构造动态SQL语句
	// 3.1 判断是否有关键字Username查询
	if in.Username != "" {
		md = md.Where(dao.User.Columns().Username, in.Username)
	}
	// 3.2 判断是否有关键字Name查询
	if in.Name != "" {
		md = md.WhereLike("CONCAT(last_name, first_name)", "%"+in.Name+"%")
	}
	// 3.3 判断是否有关键字Status查询
	if in.Status != -1 {
		md = md.Where(dao.User.Columns().Status, in.Status)
	}
	// 3.4 判断是否有关键字BeforeDate和AfterDate查询
	if (in.BeforeDate != nil) && (in.AfterDate != nil) {
		md = md.WhereBetween(dao.User.Columns().CreateTime, in.BeforeDate, in.AfterDate)
	}

	// 4. 执行分页查询
	// 设置排序：更新时间降序
	md = md.OrderDesc(dao.User.Columns().UpdateTime)
	// 设置分页
	md = md.Page(in.Page, in.PageSize)

	// 5. 判断当前页的数据条数
	out.Total, err = md.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 6. 将查询结果赋值给响应结构体
	if err := md.Scan(&out.Items); err != nil {
		return out, err
	}

	return
}

// GetById implements service.IUser.
func (*iUser) GetById(ctx context.Context, in model.UserGetByIdInput) (out *model.UserGetByIdOutput, err error) {
	out = &model.UserGetByIdOutput{}

	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Scan(&out.UserInfoBase)
	if err != nil {
		return out, gerror.New("该用户不存在")
	}

	return
}
