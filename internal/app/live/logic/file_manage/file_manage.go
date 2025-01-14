package file_manage

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/model"
	"github.com/shichen437/live-dog/internal/app/live/service"
	"github.com/shichen437/live-dog/internal/pkg/utils"
)

func init() {
	service.RegisterFileManage(New())
}

func New() *sFileManage {
	return &sFileManage{}
}

type sFileManage struct {
}

func (f *sFileManage) List(ctx context.Context, req *v1.GetFileInfoListReq) (res *v1.GetFileInfoListRes, err error) {
	res = &v1.GetFileInfoListRes{}
	base, err := filepath.Abs(utils.Output)
	utils.WriteErrLogT(ctx, err, "无效目录")
	absPath, err := filepath.Abs(filepath.Join(base, req.Path))
	utils.WriteErrLogT(ctx, err, "无效路径")
	if !strings.HasPrefix(absPath, base) {
		return nil, utils.TError(ctx, "无效路径")
	}
	files, err := os.ReadDir(absPath)
	utils.WriteErrLogT(ctx, err, "无效路径")
	if len(files) == 0 || err != nil {
		return
	}
	var list []*model.FileInfo
	for _, file := range files {
		info, err := file.Info()
		if err != nil || isHiddenFile(info) {
			continue
		}
		if req.Filename != "" && !matchPattern(file.Name(), req.Filename) {
			continue
		}
		list = append(list, &model.FileInfo{
			Filename:     file.Name(),
			IsFolder:     file.IsDir(),
			Size:         info.Size(),
			LastModified: info.ModTime().Local().UnixMilli(),
		})
	}
	res.Rows = list
	return
}

func (f *sFileManage) Delete(ctx context.Context, req *v1.DeleteFileInfoReq) (res *v1.DeleteFileInfoRes, err error) {
	res = &v1.DeleteFileInfoRes{}
	if len(req.Filenames) == 0 {
		return
	}
	base, err := filepath.Abs(utils.Output)
	utils.WriteErrLogT(ctx, err, "无效目录")
	absPath, err := filepath.Abs(filepath.Join(base, req.Path))
	utils.WriteErrLogT(ctx, err, "无效路径")
	if !strings.HasPrefix(absPath, base) {
		return nil, utils.TError(ctx, "无效路径")
	}
	for _, filename := range req.Filenames {
		err = os.RemoveAll(filepath.Join(absPath, filename))
		utils.WriteErrLogT(ctx, err, "删除失败")
	}
	return
}

func isHiddenFile(file fs.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	return strings.HasPrefix(file.Name(), ".")
}

func matchPattern(filename, pattern string) bool {
	return strings.Contains(filename, pattern)
}
