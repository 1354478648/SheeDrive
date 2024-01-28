package admin

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/model/entity"
	"SheeDrive/internal/service"
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

// Login implements service.IAdmin.
func (*iAdmin) Login(ctx context.Context, username string, password string) (admin *entity.Admin, err error) {
	err = dao.Admin.Ctx(ctx).Where(do.Admin{
		Username: username,
		Password: password,
	}).Scan(&admin)

	if admin == nil {
		err = gerror.New("用户名或密码不正确")
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
	var adminList []model.AdminBase
	if err := md.Scan(&adminList); err != nil {
		return out, err
	}
	// 为CreateUserName字段赋值
	for i, admin := range adminList {
		createUser, err := service.Admin().GetById(ctx, admin.CreateUser)
		if err != nil {
			return out, err
		}
		adminList[i].CreateUserName = createUser.Name
	}
	out.Items = adminList

	return
}

// GetById implements service.IAdmin.
func (*iAdmin) GetById(ctx context.Context, id int64) (admin *entity.Admin, err error) {
	err = dao.Admin.Ctx(ctx).Where(do.Admin{
		Id: id,
	}).Scan(&admin)

	return
}
