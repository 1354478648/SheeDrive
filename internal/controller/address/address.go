package address

import (
	apiAddress "SheeDrive/api/address"
	"SheeDrive/internal/model"
	"SheeDrive/internal/service"
	"context"
	"fmt"
)

var AddressController = &cAddress{}

type cAddress struct{}

// 用户地址列表查询
func (c *cAddress) UserAddressGetList(ctx context.Context, req *apiAddress.UserAddressGetListReq) (res *apiAddress.UserAddressGetListRes, err error) {
	out, err := service.Address().GetList(ctx, model.UserAddressGetListInput{
		BelongId:       req.Id,
		BelongCategory: 2,
	})
	if err != nil {
		return nil, err
	}

	res = &apiAddress.UserAddressGetListRes{
		AddressInfo: out.AddressInfoBase,
	}
	return
}

// 用户地址添加
func (c *cAddress) UserAddressAdd(ctx context.Context, req *apiAddress.UserAddressAddReq) (res *apiAddress.UserAddressAddRes, err error) {
	out, err := service.Address().Add(ctx, model.UserAddressAddInput{
		BelongId: req.Id,
		UserAddressAddUpdateBase: model.UserAddressAddUpdateBase{
			Province: req.Province,
			City:     req.City,
			District: req.District,
			Detail:   req.DetailAddress,
		},
	})
	if err != nil {
		return nil, err
	}
	res = &apiAddress.UserAddressAddRes{
		Id: fmt.Sprintf("%v", out.Id),
	}
	return
}

// 通过Id查询用户地址
func (c *cAddress) UserAddressGetById(ctx context.Context, req *apiAddress.UserAddressGetByIdReq) (res *apiAddress.UserAddressGetByIdRes, err error) {
	out, err := service.Address().GetById(ctx, model.UserAddressGetByIdInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res = &apiAddress.UserAddressGetByIdRes{
		AddressInfo: out.AddressInfoBase,
	}
	return
}

// 用户地址修改
func (c *cAddress) UserAddressUpdate(ctx context.Context, req *apiAddress.UserAddressUpdateReq) (res *apiAddress.UserAddressUpdateRes, err error) {
	err = service.Address().Update(ctx, model.UserAddressUpdateInput{
		Id: req.Id,
		UserAddressAddUpdateBase: model.UserAddressAddUpdateBase{
			Province: req.Province,
			City:     req.City,
			District: req.District,
			Detail:   req.DetailAddress,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}

// 用户地址删除
func (c *cAddress) UserAddressDelete(ctx context.Context, req *apiAddress.UserAddressDeleteReq) (res *apiAddress.UserAddressDeleteRes, err error) {
	err = service.Address().Delete(ctx, model.UserAddressDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
