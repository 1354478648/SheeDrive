package user

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"
	"strings"

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
	// 这里还需要添加身份证的唯一性索引，另外想从err错误信息中提取错误信息，打印不同的错误提示
	// ALTER TABLE `user` ADD CONSTRAINT `uc_idNumber` UNIQUE (`id_number`)
	if err != nil {
		if strings.Contains(err.Error(), "idx_username") {
			return out, gerror.New("该用户名已存在")
		}
		if strings.Contains(err.Error(), "id_number") {
			return out, gerror.New("该身份证号已存在")
		}
		return out, gerror.New("注册失败")
	}

	// 返回新注册用户信息
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

// Delete implements service.IUser.
func (*iUser) Delete(ctx context.Context, in model.UserDeleteInput) (err error) {
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除用户失败")
	}
	return
}

// UpdateAvatar implements service.IUser.
func (*iUser) UpdateAvatar(ctx context.Context, in model.UserUpdateAvatarInput) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{Avatar: in.Url}).Where(dao.User.Columns().Id, in.Id).Update()
	if err != nil {
		return gerror.New("修改头像失败")
	}
	return
}

// UpdatePassword implements service.IUser.
func (*iUser) UpdatePassword(ctx context.Context, in model.UserUpdatePasswordInput) (err error) {
	// 获取原密码
	oldPassword, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Value(dao.User.Columns().Password)
	if err != nil {
		return gerror.New("获取原密码失败")
	}
	if utility.EncryptPassword(in.OldPassword) != oldPassword.String() {
		return gerror.New("新旧密码不一致")
	}
	// 更新密码
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(do.User{Password: utility.EncryptPassword(in.NewPassword)}).Update()
	if err != nil {
		return gerror.New("修改密码失败")
	}
	return
}

// UpdateStatus implements service.IUser.
func (*iUser) UpdateStatus(ctx context.Context, in model.UserUpdateStatusInput) (err error) {
	id := in.Id
	userInfo, err := service.User().GetById(ctx, model.UserGetByIdInput{
		Id: id,
	})
	if err != nil {
		return gerror.New("该用户不存在")
	}
	switch userInfo.Status {
	case 0:
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(do.User{Status: 1}).Update()
		if err != nil {
			return gerror.New("用户状态切换失败")
		}
	case 1:
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Data(do.User{Status: 0}).Update()
		if err != nil {
			return gerror.New("用户状态切换失败")
		}
	}
	return
}
