package crons

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shichen437/live-dog/internal/pkg/listeners"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func AddCron(liveId int) {
	global := utils.GetGlobal(gctx.GetInitCtx())
	m, ok := global.ModelsMap[liveId]
	if !ok {
		g.Log().Error(gctx.GetInitCtx(), "未找到该直播间信息")
		return
	}
	if m.LiveManage.MonitorType != 1 {
		g.Log().Error(gctx.GetInitCtx(), "非定时监控直播间")
		return
	}
	cronStart, err := parseCronTime(m.LiveManage.MonitorStart)
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "解析开始时间失败")
		return
	}
	cronStop, err := parseCronTime(m.LiveManage.MonitorStop)
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "解析停止时间失败")
		return
	}
	l, ok := global.Lives[liveId]
	if !ok {
		g.Log().Error(gctx.GetInitCtx(), "未找到该直播间映射信息")
		return
	}
	gcron.Add(strings.Trim(cronStart, " "), func() {
		g.Log().Info(gctx.GetInitCtx(), "启动定时任务-", liveId)
		global.ListenerManager.(listeners.Manager).AddListener(gctx.GetInitCtx(), l)
	}, "Cron-Start-"+strconv.Itoa(m.LiveManage.Id))
	gcron.Add(strings.Trim(cronStop, " "), func() {
		g.Log().Info(gctx.GetInitCtx(), "停止定时任务-", liveId)
		global.ListenerManager.(listeners.Manager).RemoveListener(gctx.GetInitCtx(), l.GetLiveId())
	}, "Cron-Stop-"+strconv.Itoa(m.LiveManage.Id))
	if utils.IsTimeRange(m.LiveManage.MonitorStart, m.LiveManage.MonitorStop) {
		global.ListenerManager.(listeners.Manager).AddListener(gctx.GetInitCtx(), l)
	}
}

func RemoveCron(liveId int) {
	search := gcron.Search("Cron-Start-" + strconv.Itoa(liveId))
	if search != nil {
		gcron.Remove(search.Name)
	}
	search = gcron.Search("Cron-Stop-" + strconv.Itoa(liveId))
	if search != nil {
		gcron.Remove(search.Name)
	}
	utils.GetGlobal(gctx.GetInitCtx()).ListenerManager.(listeners.Manager).RemoveListener(gctx.GetInitCtx(), liveId)
	g.Log().Info(gctx.GetInitCtx(), "移除定时任务-", liveId)
}

func parseCronTime(t string) (string, error) {
	if len(t) == 0 {
		return "", gerror.New("传入时间不能为空")
	}
	s := strings.Split(t, ":")
	if len(s) != 2 {
		return "", gerror.New("传入时间不能为空")
	}
	return fmt.Sprintf("0 %s %s * * *", s[1], s[0]), nil
}
