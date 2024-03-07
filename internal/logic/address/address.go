package address

import (
	"SheeDrive/internal/dao"
	"SheeDrive/internal/model"
	"SheeDrive/internal/model/do"
	"SheeDrive/internal/service"
	"SheeDrive/utility"
	"context"
	"fmt"

	"github.com/gogf/gf/errors/gerror"
)

type iAddress struct{}

func New() *iAddress {
	return &iAddress{}
}

func init() {
	service.RegisterAddress(New())
}

// GetList implements service.IAddress.
func (i *iAddress) GetList(ctx context.Context, in model.UserAddressGetListInput) (out *model.UserAddressGetListOutput, err error) {
	out = &model.UserAddressGetListOutput{}

	err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().BelongId, in.BelongId).Where(dao.Address.Columns().BelongCategory, in.BelongCategory).Scan(&out.AddressInfoBase)
	if err != nil {
		return out, gerror.New("地址不存在")
	}

	return
}

// Add implements service.IAddress.
func (i *iAddress) Add(ctx context.Context, in model.UserAddressAddInput) (out *model.UserAddressAddOutput, err error) {
	out = &model.UserAddressAddOutput{}

	id := utility.GenSnowFlakeId()

	geocode, err := utility.Geocoding(fmt.Sprintf("%s%s%s%s", in.Province, in.City, in.District, in.Detail), in.City)
	if err != nil {
		return out, err
	}

	_, err = dao.Address.Ctx(ctx).Data(do.Address{
		Id:             id,
		BelongId:       in.BelongId,
		BelongCategory: 2,
		LngLat:         geocode.Location,
		Province:       in.Province,
		City:           in.City,
		District:       in.District,
		Detail:         in.Detail,
	}).Insert()

	if err != nil {
		return out, err
	}

	out.Id = id

	return
}

// GetById implements service.IAddress.
func (i *iAddress) GetById(ctx context.Context, in model.UserAddressGetByIdInput) (out *model.UserAddressGetByIdOutput, err error) {
	out = &model.UserAddressGetByIdOutput{}

	err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().Id, in.Id).Scan(&out.AddressInfoBase)
	if err != nil {
		return nil, gerror.New("该地址不存在")
	}

	return
}

// Delete implements service.IAddress.
func (i *iAddress) Delete(ctx context.Context, in model.UserAddressDeleteInput) (err error) {
	_, err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().Id, in.Id).Delete()
	if err != nil {
		return gerror.New("地址删除失败")
	}

	// 或许要删除订单

	return
}

// Update implements service.IAddress.
func (i *iAddress) Update(ctx context.Context, in model.UserAddressUpdateInput) (err error) {
	geocode, err := utility.Geocoding(fmt.Sprintf("%s%s%s%s", in.Province, in.City, in.District, in.Detail), in.City)
	if err != nil {
		return err
	}

	_, err = dao.Address.Ctx(ctx).Where(dao.Address.Columns().Id, in.Id).Data(do.Address{
		LngLat:   geocode.Location,
		Province: in.Province,
		City:     in.City,
		District: in.District,
		Detail:   in.Detail,
	}).Update()
	if err != nil {
		return gerror.New("地址更新失败")
	}
	return
}
