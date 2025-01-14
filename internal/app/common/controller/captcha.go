package common

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/tool"
	adminService "github.com/shichen437/live-dog/internal/app/admin/service"
	"github.com/shichen437/live-dog/internal/app/common/service"

	"github.com/google/uuid"
)

var Captcha = captchaController{}

type captchaController struct {
}

// CaptchaImage 获取验证码
func (c *captchaController) CaptchaImage(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	guid := uuid.New()
	res = &v1.CaptchaRes{
		Key:            idKeyC,
		Img:            base64stringC,
		Uuid:           guid.String(),
		CaptchaEnabled: true,
	}
	return
}

// CaptchaImage 获取验证码
func (c *captchaController) Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	err = adminService.SysMenu().InitApiPath(ctx)
	return
}
