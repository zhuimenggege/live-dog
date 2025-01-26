package sys_push

import (
	"context"

	v1 "github.com/shichen437/live-dog/api/v1/admin"
	"github.com/shichen437/live-dog/internal/app/admin/consts"
	"github.com/shichen437/live-dog/internal/app/admin/dao"
	"github.com/shichen437/live-dog/internal/app/admin/model"
	"github.com/shichen437/live-dog/internal/app/admin/model/do"
	"github.com/shichen437/live-dog/internal/app/admin/model/entity"
	"github.com/shichen437/live-dog/internal/app/admin/service"
	commonConst "github.com/shichen437/live-dog/internal/app/common/consts"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	service.RegisterPushChannel(New())
}

func New() *sPushChannel {
	return &sPushChannel{}
}

type sPushChannel struct {
}

func (s *sPushChannel) Add(ctx context.Context, req *v1.PostPushChannelReq) (res *v1.PostPushChannelRes, err error) {
	adminName := gconv.String(ctx.Value(commonConst.CtxAdminName))
	if req.Type == "" || !utils.InSliceString(req.Type, &consts.PushChannelType) {
		return
	}
	if req.Type == "email" && req.Email == nil {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		//添加角色信息
		lastInfo, err := dao.PushChannel.Ctx(ctx).Insert(do.PushChannel{
			Name:       req.Name,
			Type:       req.Type,
			Status:     req.Status,
			Remark:     req.Remark,
			CreateBy:   adminName,
			CreateTime: gtime.Now(),
		})
		utils.WriteErrLogT(ctx, err, commonConst.AddF)
		lastId, err := lastInfo.LastInsertId()
		if req.Type == "email" {
			dao.PushChannelEmail.Ctx(ctx).Insert(do.PushChannelEmail{
				ChannelId:  lastId,
				Server:     req.Email.Server,
				Port:       req.Email.Port,
				From:       req.Email.From,
				To:         req.Email.To,
				AuthCode:   req.Email.AuthCode,
				CreateTime: gtime.Now(),
			})
			utils.WriteErrLogT(ctx, err, commonConst.AddF)
		} else {
			dao.PushChannelWeb.Ctx(ctx).Insert(do.PushChannelWeb{
				ChannelId:    lastId,
				Url:          req.Web.Url,
				HttpMethod:   req.Web.HttpMethod,
				Secret:       req.Web.Secret,
				AppId:        req.Web.AppId,
				CorpId:       req.Web.CorpId,
				ReceiverId:   req.Web.ReceiverId,
				ReceiverType: req.Web.ReceiverType,
				ExtraParams:  req.Web.ExtraParams,
				CreateTime:   gtime.Now(),
			})
			utils.WriteErrLogT(ctx, err, commonConst.AddF)
		}

	})
	return
}

func (s *sPushChannel) Delete(ctx context.Context, req *v1.DeletePushChannelReq) (res *v1.DeletePushChannelRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		ids := utils.ParamStrToSlice(req.Id, ",")
		_, e := dao.PushChannel.Ctx(ctx).WhereIn(dao.PushChannel.Columns().Id, ids).Delete()
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
		_, e = dao.PushChannelEmail.Ctx(ctx).WhereIn(dao.PushChannelEmail.Columns().ChannelId, ids).Delete()
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
		_, e = dao.PushChannelWeb.Ctx(ctx).WhereIn(dao.PushChannelWeb.Columns().ChannelId, ids).Delete()
		utils.WriteErrLogT(ctx, e, commonConst.DeleteF)
	})
	return
}

func (s *sPushChannel) Update(ctx context.Context, req *v1.PutPushChannelReq) (res *v1.PutPushChannelRes, err error) {
	if req.Id == 0 {
		err = utils.TError(ctx, commonConst.IDEmpty)
		return
	}
	var source *entity.PushChannel
	err = g.Try(ctx, func(ctx context.Context) {
		dao.PushChannel.Ctx(ctx).WherePri(req.Id).Scan(&source)
	})
	if err != nil || source == nil || source.Type != req.Type {
		err = utils.TError(ctx, commonConst.UpdateF)
		return
	}
	g.Try(ctx, func(ctx context.Context) {
		_, e := dao.PushChannel.Ctx(ctx).WherePri(&req.Id).Update(do.PushChannel{
			Name:       req.Name,
			Type:       req.Type,
			Status:     req.Status,
			Remark:     req.Remark,
			ActionTime: gtime.Now(),
		})
		utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
	})
	if source.Type == "email" {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.PushChannelEmail.Ctx(ctx).Where(dao.PushChannelEmail.Columns().ChannelId, &req.Id).Update(do.PushChannelEmail{
				Server:     req.Email.Server,
				Port:       req.Email.Port,
				AuthCode:   req.Email.AuthCode,
				From:       req.Email.From,
				To:         req.Email.To,
				ActionTime: gtime.Now(),
			})
			utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
		})
	} else {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.PushChannelWeb.Ctx(ctx).Where(dao.PushChannelWeb.Columns().ChannelId, &req.Id).Update(do.PushChannelWeb{
				Url:          req.Web.Url,
				HttpMethod:   req.Web.HttpMethod,
				Secret:       req.Web.Secret,
				AppId:        req.Web.AppId,
				CorpId:       req.Web.CorpId,
				ReceiverId:   req.Web.ReceiverId,
				ReceiverType: req.Web.ReceiverType,
				ExtraParams:  req.Web.ExtraParams,
				ActionTime:   gtime.Now(),
			})
			utils.WriteErrLogT(ctx, e, commonConst.UpdateF)
		})
	}
	return
}

func (s *sPushChannel) List(ctx context.Context, req *v1.GetPushChannelListReq) (res *v1.GetPushChannelListRes, err error) {
	res = &v1.GetPushChannelListRes{}
	m := dao.PushChannel.Ctx(ctx)
	if req.Type != "" {
		m = m.Where(dao.PushChannel.Columns().Type, req.Type)
	}
	if req.Name != "" {
		m = m.WhereLike(dao.PushChannel.Columns().Name, "%"+req.Name+"%")
	}
	err = g.Try(ctx, func(ctx context.Context) {
		var result []*entity.PushChannel
		res.Total, err = m.Count()
		err = m.Page(req.PageNum, req.PageSize).Scan(&result)
		utils.WriteErrLogT(ctx, err, commonConst.ListF)
		res.Rows = result
	})
	return
}

func (s *sPushChannel) Get(ctx context.Context, req *v1.GetPushChannelReq) (res *v1.GetPushChannelRes, err error) {
	res = &v1.GetPushChannelRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.PushChannel.Ctx(ctx).WherePri(req.Id).Scan(&res)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
	})
	if res.Type == "email" {
		email := &entity.PushChannelEmail{}
		err = dao.PushChannelEmail.Ctx(ctx).Where(dao.PushChannelEmail.Columns().ChannelId, req.Id).Scan(&email)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		res.Email = email
	} else {
		web := &entity.PushChannelWeb{}
		err = dao.PushChannelWeb.Ctx(ctx).Where(dao.PushChannelWeb.Columns().ChannelId, req.Id).Scan(&web)
		utils.WriteErrLogT(ctx, err, commonConst.GetF)
		res.Web = web
	}
	return
}

func (s *sPushChannel) ListAll(ctx context.Context) (res []*model.PushChannel, err error) {
	m := dao.PushChannel.Ctx(ctx)
	err = g.Try(ctx, func(ctx context.Context) {
		m = m.Where(dao.PushChannel.Columns().Status, 1)
		err = m.Scan(&res)
		utils.WriteErrLogT(ctx, err, commonConst.ListF)
	})
	if len(res) == 0 {
		return
	}
	arr := []int{}
	cMap := make(map[int]*model.PushChannel)
	for _, v := range res {
		arr = append(arr, v.Id)
		cMap[v.Id] = v
	}
	emails := []*entity.PushChannelEmail{}
	err = dao.PushChannelEmail.Ctx(ctx).WhereIn(dao.PushChannelEmail.Columns().ChannelId, arr).Scan(&emails)
	utils.WriteErrLogT(ctx, err, commonConst.ListF)
	for _, v := range emails {
		cMap[v.ChannelId].Email = v
	}
	webs := []*entity.PushChannelWeb{}
	err = dao.PushChannelWeb.Ctx(ctx).WhereIn(dao.PushChannelWeb.Columns().ChannelId, arr).Scan(&webs)
	utils.WriteErrLogT(ctx, err, commonConst.ListF)
	for _, v := range webs {
		cMap[v.ChannelId].Web = v
	}
	return
}
