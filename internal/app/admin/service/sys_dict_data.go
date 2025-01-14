// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
)

type (
	ISysDictData interface {
		GetDictDataByType(ctx context.Context, dictType string) (dictdata []*entity.SysDictData, err error)
		GetDictDataList(ctx context.Context, req *v1.GetDictDataListReq) (res *v1.GetDictDataListRes, err error)
		GetDictData(ctx context.Context, req *v1.GetDictDataDetailReq) (res *v1.GetDictDataDetailRes, err error)
		Add(ctx context.Context, req *v1.PostDictDataReq) (res *v1.PostDictDataRes, err error)
		Update(ctx context.Context, req *v1.PutDictDataReq) (res *v1.PutDictDataRes, err error)
		Delete(ctx context.Context, req *v1.DeleteDictDataReq) (res *v1.DeleteDictDataRes, err error)
	}
)

var (
	localSysDictData ISysDictData
)

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}
