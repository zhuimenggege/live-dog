package admin

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/service"
)

type dictController struct {
}

var DictData = dictController{}

/**
 * 根据获取字典数据
 */
func (d *dictController) DictData(ctx context.Context, req *v1.GetDictDataReq) (res *v1.GetDictDataRes, err error) {
	dictData, err := service.SysDictData().GetDictDataByType(ctx, req.DictType)
	res = &v1.GetDictDataRes{
		DictData: dictData,
	}
	return
}

func (d *dictController) GetDictDataList(ctx context.Context, req *v1.GetDictDataListReq) (res *v1.GetDictDataListRes, err error) {
	res, err = service.SysDictData().GetDictDataList(ctx, req)
	return
}

func (d *dictController) GetDictDataDetail(ctx context.Context, req *v1.GetDictDataDetailReq) (res *v1.GetDictDataDetailRes, err error) {
	res, err = service.SysDictData().GetDictData(ctx, req)
	return
}

func (d *dictController) Add(ctx context.Context, req *v1.PostDictDataReq) (res *v1.PostDictDataRes, err error) {
	res, err = service.SysDictData().Add(ctx, req)
	return
}

func (d *dictController) Update(ctx context.Context, req *v1.PutDictDataReq) (res *v1.PutDictDataRes, err error) {
	res, err = service.SysDictData().Update(ctx, req)
	return
}

func (d *dictController) Delete(ctx context.Context, req *v1.DeleteDictDataReq) (res *v1.DeleteDictDataRes, err error) {
	res, err = service.SysDictData().Delete(ctx, req)
	return
}
