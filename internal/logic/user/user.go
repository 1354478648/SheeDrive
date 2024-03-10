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
	"github.com/gogf/gf/v2/frame/g"
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

	// 生成token
	out.Token = utility.GenToken(in.Username)
	// 将token保存到redis中
	err = g.Redis().SetEX(ctx, out.Token, out.Token, 86400)
	if err != nil {
		return nil, gerror.New("Token保存失败")
	}

	// 将Token持久化
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, out.UserInfoBase.Id).Data(
		do.Admin{Token: out.Token}).Update()
	if err != nil {
		return nil, gerror.New("Token保存失败")
	}

	return
}

// LoginByPhone implements service.IUser.
func (*iUser) LoginByPhone(ctx context.Context, in model.UserLoginByPhoneInput) (out *model.UserLoginByPhoneOutput, err error) {
	out = &model.UserLoginByPhoneOutput{}

	// 验证码验证
	code, err := g.Redis().Get(ctx, in.Phone)
	if err != nil {
		return nil, gerror.New("验证码获取失败")
	}
	if code.Int() == 0 {
		return nil, gerror.New("验证码已过期")
	}
	if code.Int() != in.Code {
		return nil, gerror.New("验证码错误")
	}

	err = dao.User.Ctx(ctx).Where(do.User{
		Phone: in.Phone,
	}).Scan(&out.UserInfoBase)
	if err != nil {
		return nil, gerror.New("用户不存在")
	}
	// 判断用户状态是否被禁用
	if out.UserInfoBase.Status == 0 {
		return nil, gerror.New("该用户已被禁用")
	}

	// 生成token
	out.Token = utility.GenToken(in.Phone)
	// 将token保存到redis中
	err = g.Redis().SetEX(ctx, out.Token, out.Token, 86400)
	if err != nil {
		return nil, gerror.New("Token保存失败")
	}

	// 将Token持久化
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, out.UserInfoBase.Id).Data(
		do.Admin{Token: out.Token}).Update()
	if err != nil {
		return nil, gerror.New("Token保存失败")
	}

	// 删除验证码
	_, err = g.Redis().Del(ctx, in.Phone)
	if err != nil {
		return nil, gerror.New("验证码删除失败")
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
	// 删除对应id的token
	id := in.Id
	userInfo, err := service.User().GetById(ctx, model.UserGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("未找到该用户")
	}
	_, err = g.Redis().Del(ctx, userInfo.Token)
	if err != nil {
		return gerror.New("token删除失败")
	}

	// 执行删除操作
	_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除用户失败")
	}

	// 执行删除地址操作
	_, err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().BelongCategory, 2).Where(dao.Address.Columns().BelongId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除用户地址失败")
	}

	// 执行删除评价操作
	orderId, err := dao.Order.Ctx(ctx).Fields("id").Where(dao.Order.Columns().UserId, in.Id).Array()
	if err != nil {
		return gerror.New("未找到该用户下的订单")
	}
	_, err = dao.Comment.Ctx(ctx).WhereIn(dao.Comment.Columns().OrderId, orderId).Delete()
	if err != nil {
		return gerror.New("删除评价失败")
	}

	// 执行删除订单操作
	_, err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().UserId, in.Id).Delete()
	if err != nil {
		return gerror.New("删除订单失败")
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

	// 删除对应id的token
	id := in.Id
	userInfo, err := service.User().GetById(ctx, model.UserGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("未找到该用户")
	}
	_, err = g.Redis().Del(ctx, userInfo.Token)
	if err != nil {
		return gerror.New("token删除失败")
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

// UpdatePasswordByPhone implements service.IUser.
func (*iUser) UpdatePasswordByPhone(ctx context.Context, in model.UserUpdatePasswordByPhoneInput) (err error) {
	// 验证码验证
	code, err := g.Redis().Get(ctx, in.Phone)
	if err != nil {
		return gerror.New("验证码获取失败")
	}
	if code.Int() == 0 {
		return gerror.New("验证码已过期")
	}
	if code.Int() != in.Code {
		return gerror.New("验证码错误")
	}
	// 修改密码
	result, err := dao.User.Ctx(ctx).Data(do.User{Password: utility.EncryptPassword(in.Password)}).Where(dao.User.Columns().Phone, in.Phone).Update()
	if err != nil {
		return gerror.New("修改密码失败")
	}
	row, err := result.RowsAffected()
	if row == 0 {
		return gerror.New("手机号不存在")
	}
	if err != nil {
		return gerror.New("修改密码失败")
	}
	_, err = g.Redis().Del(ctx, in.Phone)
	if err != nil {
		return gerror.New("验证码删除失败")
	}
	return
}
