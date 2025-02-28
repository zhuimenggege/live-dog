package listeners

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/shichen437/live-dog/internal/pkg/events"
	"github.com/shichen437/live-dog/internal/pkg/lives"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/lthibault/jitterbug"
)

type Listener interface {
	Start() error
	Close()
}

type listener struct {
	Live   lives.Live
	status status
	ed     events.Dispatcher
	state  uint32
	stop   chan struct{}
}

func NewListener(ctx context.Context, live lives.Live) Listener {
	global := utils.GetGlobal(ctx)
	return &listener{
		Live:   live,
		status: status{},
		ed:     global.EventDispatcher.(events.Dispatcher),
		state:  begin,
		stop:   make(chan struct{}),
	}
}

func (l *listener) Start() error {
	if !atomic.CompareAndSwapUint32(&l.state, begin, pending) {
		return nil
	}
	defer atomic.CompareAndSwapUint32(&l.state, pending, running)

	l.ed.DispatchEvent(events.NewEvent("ListenStart", l.Live))
	l.refresh()
	go l.run()
	return nil
}

func (l *listener) Close() {
	if !atomic.CompareAndSwapUint32(&l.state, running, stopped) {
		return
	}
	l.ed.DispatchEvent(events.NewEvent("ListenStop", l.Live))
	close(l.stop)
}

func (l *listener) refresh() {
	info, err := l.Live.GetInfo()
	if err != nil {
		g.Log().Error(gctx.GetInitCtx(), "failed to get live info")
		return
	}
	var evtTyp events.EventType
	latestStatus := l.getLatestStatus(info)
	defer func() { l.status = latestStatus }()
	isStatusChanged := true
	switch l.status.Diff(latestStatus) {
	case 0:
		isStatusChanged = false
	case statusToTrueEvt:
		evtTyp = "LiveStart"
	case statusToFalseEvt:
		evtTyp = "LiveEnd"
	}
	if isStatusChanged {
		l.ed.DispatchEvent(events.NewEvent(evtTyp, l.Live))
	}
	evtTyp = "NameChanged"
	if !info.LiveStatus && l.nameChanged(info) {
		l.ed.DispatchEvent(events.NewEvent(evtTyp, l.Live))
	}
}

func (l *listener) getLatestStatus(info *lives.RoomInfo) status {
	return status{
		roomName:   info.RoomName,
		anchor:     info.Anchor,
		roomStatus: info.LiveStatus,
	}
}

func (l *listener) nameChanged(info *lives.RoomInfo) bool {
	global := utils.GetGlobal(gctx.GetInitCtx())
	source, ok := global.ModelsMap[l.Live.GetLiveId()]
	if ok {
		return source.RoomInfo.Anchor != info.Anchor || source.RoomInfo.RoomName != info.RoomName
	}
	return false
}

func (l *listener) run() {
	ticker := jitterbug.New(
		time.Duration(30)*time.Second,
		jitterbug.Norm{
			Stdev: time.Second * 3,
		},
	)
	defer ticker.Stop()

	for {
		select {
		case <-l.stop:
			return
		case <-ticker.C:
			l.refresh()
		}
	}
}
