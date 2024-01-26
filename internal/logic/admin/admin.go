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
	// 初始化out
	out = &model.AdminGetListOutput{}
	var (
		model = dao.Admin.Ctx(ctx)
	)

	// 构造动态SQL语句
	// 判断是否有关键字Username查询
	if in.Username != "" {
		model = model.Where(dao.Admin.Columns().Username, in.Username)
	}
	// 判断是否有关键字Name查询
	if in.Name != "" {
		model = model.WhereLike(dao.Admin.Columns().Name, "%"+in.Name+"%")
	}
	// 判断是否有关键字BeforeDate和AfterDate查询
	if (in.BeforeDate != nil) && (in.AfterDate != nil) {
		model = model.WhereBetween(dao.Admin.Columns().CreateTime, in.BeforeDate, in.AfterDate)
	}

	// 设置排序：更新时间降序
	model = model.OrderDesc(dao.Admin.Columns().UpdateTime)
	// 设置分页
	model = model.Page(in.Page, in.PageSize)
	// 执行查询
	var adminList []*entity.Admin
	if err := model.Scan(&adminList); err != nil {
		return out, err
	}
	// 判断是否查出数据
	if len(adminList) == 0 {
		return out, err
	}
	// 将密码字段置为空字符串
	for _, admin := range adminList {
		admin.Password = ""
	}
	out.Items = adminList

	// 计算当前页的数据条数
	out.Total, err = model.Count()
	if err != nil {
		return out, err
	}

	return
}
