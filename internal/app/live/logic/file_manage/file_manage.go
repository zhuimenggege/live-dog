package file_manage

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
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

func (c *sFileManage) Play(ctx context.Context, req *v1.GetFilePlayReq) (res *v1.GetFilePlayRes, err error) {
	filePath := req.Path
	if filePath == "" {
		return nil, gerror.New("文件路径不能为空")
	}

	// 获取文件的绝对路径
	base, err := filepath.Abs(utils.Output)
	if err != nil {
		return nil, gerror.New("无效目录")
	}
	absPath, err := filepath.Abs(filepath.Join(base, filePath))
	if err != nil {
		return nil, gerror.New("无效路径")
	}
	if !strings.HasPrefix(absPath, base) {
		return nil, gerror.New("无效路径")
	}

	// 检查文件是否存在
	if !gfile.Exists(absPath) {
		return nil, gerror.Newf("文件不存在: %s", filePath)
	}

	// 获取文件信息
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 获取文件扩展名并设置正确的 Content-Type
	ext := strings.ToLower(filepath.Ext(absPath))
	var contentType string
	switch ext {
	case ".mp4":
		contentType = "video/mp4"
	case ".flv":
		contentType = "video/x-flv"
	case ".aac":
		contentType = "audio/aac"
	case ".mp3":
		contentType = "audio/mpeg"
	default:
		contentType = "application/octet-stream"
	}

	// 设置响应头
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", contentType)
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Range, Authorization")
	r.Response.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range, Accept-Ranges")

	if r.Method == "OPTIONS" {
		r.Response.WriteStatus(200)
		return &v1.GetFilePlayRes{}, nil
	}

	rangeHeader := r.Header.Get("Range")
	if rangeHeader != "" {
		// 解析Range头部，格式为: bytes=start-end
		parts := strings.Split(strings.TrimPrefix(rangeHeader, "bytes="), "-")
		if len(parts) != 2 {
			return nil, gerror.New("无效的Range头部格式")
		}

		start := int64(0)
		end := fileInfo.Size() - 1

		if parts[0] != "" {
			start = gconv.Int64(parts[0])
		}

		if parts[1] != "" {
			end = gconv.Int64(parts[1])
			if end < 0 {
				return nil, gerror.New("无效的Range结束位置")
			}
			if end >= fileInfo.Size() {
				end = fileInfo.Size() - 1
			}
		}

		if start > end || start < 0 {
			return nil, gerror.New("Range范围无效")
		}

		ranges := []struct{ Start, End int64 }{{Start: start, End: end}}
		if err != nil {
			return nil, err
		}
		if len(ranges) > 0 {
			rang := ranges[0]
			r.Response.Status = 206
			r.Response.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", rang.Start, rang.End, fileInfo.Size()))
			r.Response.Header().Set("Content-Length", gconv.String(rang.End-rang.Start+1))
			_, err = file.Seek(rang.Start, 0)
			if err != nil {
				return nil, err
			}
			_, err = io.CopyN(r.Response.Writer, file, rang.End-rang.Start+1)
			if err != nil {
				return nil, err
			}
			return &v1.GetFilePlayRes{}, nil
		}
	} else {
		r.Response.Header().Set("Content-Length", gconv.String(fileInfo.Size()))
	}

	// 使用 bufio.NewReader 优化读取性能
	reader := bufio.NewReader(file)
	_, err = io.Copy(r.Response.Writer, reader)
	if err != nil {
		return nil, err
	}

	return &v1.GetFilePlayRes{}, nil
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
