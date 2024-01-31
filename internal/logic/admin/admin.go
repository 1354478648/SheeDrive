package admin

import (
	"SheeDrive/internal/consts"
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type iAdmin struct{}

func New() *iAdmin {
	return &iAdmin{}
}

func init() {
	service.RegisterAdmin(New())
}

// 局部公共方法：判断管理员是否具有超级权限
func isRoot(ctx context.Context, adminId int64) (err error) {
	adminInfo, err := service.Admin().GetById(ctx, model.AdminGetByIdInput{Id: adminId})
	if err != nil {
		return gerror.New("该管理员不存在")
	}
	if adminInfo.IsRoot == 1 {
		return gerror.New("无法修改系统管理员")
	}
	if adminInfo.IsRoot == 0 {
		return nil
	}
	return
}

// Login implements service.IAdmin.
func (*iAdmin) Login(ctx context.Context, in model.AdminLoginInput) (out *model.AdminLoginOutput, err error) {
	// 实例化响应结构体
	out = &model.AdminLoginOutput{}

	err = dao.Admin.Ctx(ctx).Where(do.Admin{
		Username: in.Username,
		Password: utility.EncryptPassword(in.Password),
	}).Scan(&out.AdminInfoBase)
	if err != nil {
		return nil, gerror.New("用户名或密码不正确")
	}

	// 判断管理员状态是否被禁用
	if out.AdminInfoBase.Status == 0 {
		return nil, gerror.New("该管理员已被禁用")
	}

	return
}

// GetAdminList implements service.IAdmin.
func (*iAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	// 1. 实例化响应结构体
	out = &model.AdminGetListOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	// 2. 获取*gdb.Model对象
	var (
		md = dao.Admin.Ctx(ctx)
	)

	// 3. 通过md构造动态SQL语句
	// 3.1 判断是否有关键字Username查询
	if in.Username != "" {
		md = md.Where(dao.Admin.Columns().Username, in.Username)
	}
	// 3.2 判断是否有关键字Name查询
	if in.Name != "" {
		md = md.WhereLike(dao.Admin.Columns().Name, "%"+in.Name+"%")
	}
	// 3.3 判断是否有关键字BeforeDate和AfterDate查询
	if (in.BeforeDate != nil) && (in.AfterDate != nil) {
		md = md.WhereBetween(dao.Admin.Columns().CreateTime, in.BeforeDate, in.AfterDate)
	}

	// 4. 执行分页查询
	// 设置排序：更新时间降序
	md = md.OrderDesc(dao.Admin.Columns().UpdateTime)
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

// GetById implements service.IAdmin.
func (*iAdmin) GetById(ctx context.Context, in model.AdminGetByIdInput) (out *model.AdminGetByIdOutput, err error) {
	// 实例化响应结构体
	out = &model.AdminGetByIdOutput{}

	err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Scan(&out.AdminInfoBase)
	if err != nil {
		return out, gerror.New("该管理员不存在")
	}

	return
}

// Add implements service.IAdmin.
func (*iAdmin) Add(ctx context.Context, in model.AdminAddInput) (out *model.AdminAddOutput, err error) {
	// 实例化响应结构体
	out = &model.AdminAddOutput{}

	// 执行添加操作
	id, err := dao.Admin.Ctx(ctx).Data(do.Admin{
		Name:     in.Name,
		Username: in.Username,
		Password: utility.EncryptPassword(consts.DefaultPassword),
		Avatar:   "",
		Phone:    in.Phone,
		Status:   1,
		IsRoot:   0,
	}).InsertAndGetId()
	if err != nil {
		return out, gerror.New("用户名已被占用")
	}

	// 将自增主键id赋值给响应结构体
	out.Id = id

	return
}

// Update implements service.IAdmin.
func (*iAdmin) Update(ctx context.Context, in model.AdminUpdateInput) (err error) {
	// 判断管理员是否具有超级权限
	err = isRoot(ctx, in.Id)
	if err != nil {
		return err
	}

	// 执行修改操作
	_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Data(
		do.Admin{
			Name:     in.Name,
			Username: in.Username,
			Phone:    in.Phone,
		}).Update()
	if err != nil {
		return gerror.New("用户名已被占用")
	}
	return
}

// Delete implements service.IAdmin.
func (*iAdmin) Delete(ctx context.Context, in model.AdminDeleteInput) (err error) {
	// 判断管理员是否具有超级权限
	err = isRoot(ctx, in.Id)
	if err != nil {
		return err
	}

	// 执行删除操作
	_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("删除管理员失败")
	}
	return
}

// UpdateStatus implements service.IAdmin.
func (*iAdmin) UpdateStatus(ctx context.Context, in model.AdminUpdateStatusInput) (err error) {
	// 判断管理员是否具有超级权限
	err = isRoot(ctx, in.Id)
	if err != nil {
		return err
	}

	// 获取管理员原状态
	id := in.Id
	adminInfo, err := service.Admin().GetById(ctx, model.AdminGetByIdInput{Id: id})
	if err != nil {
		return gerror.New("该管理员不存在")
	}
	// 切换管理员状态
	switch adminInfo.Status {
	case 0:
		_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Data(do.Admin{Status: 1}).Update()
		if err != nil {
			return gerror.New("管理员状态切换失败")
		}
	case 1:
		_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Data(do.Admin{Status: 0}).Update()
		if err != nil {
			return gerror.New("管理员状态切换失败")
		}
	}
	return
}

// UpdatePassword implements service.IAdmin.
func (*iAdmin) UpdatePassword(ctx context.Context, in model.AdminUpdatePasswordInput) (err error) {
	// 判断管理员是否具有超级权限
	err = isRoot(ctx, in.Id)
	if err != nil {
		return err
	}

	// 获取原密码
	oldPassword, err := dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Value(dao.Admin.Columns().Password)
	if err != nil {
		return gerror.New("获取密码失败")
	}

	// 验证原密码
	if utility.EncryptPassword(in.OldPassword) != oldPassword.String() {
		return gerror.New("新旧密码不一致")
	}

	// 更新密码
	_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Data(do.Admin{Password: utility.EncryptPassword(in.NewPassword)}).Update()
	if err != nil {
		return gerror.New("更新密码失败")
	}
	return
}

// ResetPassword implements service.IAdmin.
func (*iAdmin) ResetPassword(ctx context.Context, in model.AdminResetPasswordInput) (err error) {
	// 判断管理员是否具有超级权限
	err = isRoot(ctx, in.Id)
	if err != nil {
		return err
	}

	// 重置密码
	_, err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Id, in.Id).Data(do.Admin{Password: utility.EncryptPassword(consts.DefaultPassword)}).Update()
	if err != nil {
		return gerror.New("重置密码失败")
	}
	return
}

// UpdateAvatar implements service.IAdmin.
func (*iAdmin) UpdateAvatar(ctx context.Context, in model.AdminUpdateAvatarInput) (err error) {
	// 修改头像
	_, err = dao.Admin.Ctx(ctx).Data(do.Admin{Avatar: in.Url}).Where(dao.Admin.Columns().Id, in.Id).Update()
	if err != nil {
		return gerror.New("修改头像失败")
	}
	return
}
