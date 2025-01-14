package live

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/live"
	"github.com/shichen437/live-dog/internal/app/live/service"
)

type statController struct {
}

var StatDaily = statController{}

func (c *statController) GetStatDailyById(ctx context.Context, req *v1.GetStatDailyReq) (res *v1.GetStatDailyRes, err error) {
	res, err = service.StatDaily().GetForUpdate(ctx, req)
	return
}

func (c *statController) GetStatDailyList(ctx context.Context, req *v1.GetStatDailyListReq) (res *v1.GetStatDailyListRes, err error) {
	res, err = service.StatDaily().GetStatDailyList(ctx, req)
	return
}

func (c *statController) Add(ctx context.Context, req *v1.PostStatDailyReq) (res *v1.PostStatDailyRes, err error) {
	res, err = service.StatDaily().Add(ctx, req)
	return
}

func (c *statController) Update(ctx context.Context, req *v1.PutStatDailyReq) (res *v1.PutStatDailyRes, err error) {
	res, err = service.StatDaily().Update(ctx, req)
	return
}
func (c *statController) Delete(ctx context.Context, req *v1.DeleteStatDailyReq) (res *v1.DeleteStatDailyRes, err error) {
	res, err = service.StatDaily().Delete(ctx, req)
	return
}
