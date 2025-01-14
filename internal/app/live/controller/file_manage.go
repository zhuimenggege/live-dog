package live

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/service"
)

type fileManageController struct {
}

var FileManage = fileManageController{}

func (f *fileManageController) List(ctx context.Context, req *v1.GetFileInfoListReq) (res *v1.GetFileInfoListRes, err error) {
	res, err = service.FileManage().List(ctx, req)
	return
}

func (f *fileManageController) Delete(ctx context.Context, req *v1.DeleteFileInfoReq) (res *v1.DeleteFileInfoRes, err error) {
	res, err = service.FileManage().Delete(ctx, req)
	return
}
