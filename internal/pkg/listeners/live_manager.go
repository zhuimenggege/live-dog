package listeners

import (
	"context"
	"sync"

	"github.com/shichen437/live-dog/internal/pkg/interfaces"
	"github.com/shichen437/live-dog/internal/pkg/lives"
	"github.com/shichen437/live-dog/internal/pkg/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	begin uint32 = iota
	pending
	running
	stopped
)

type manager struct {
	lock   sync.RWMutex
	savers map[int]Listener
}

type Manager interface {
	interfaces.Module
	AddListener(ctx context.Context, live lives.Live) error
	RemoveListener(ctx context.Context, liveId int) error
	GetListener(ctx context.Context, liveId int) (Listener, error)
	HasListener(ctx context.Context, liveId int) bool
}

func NewManager(ctx context.Context) Manager {
	lm := &manager{
		savers: make(map[int]Listener),
	}
	utils.GetGlobal(ctx).ListenerManager = lm
	return lm
}

func (m *manager) Start(ctx context.Context) error {
	global := utils.GetGlobal(ctx)
	if len(global.Lives) > 0 {
		global.WaitGroup.Add(1)
	}
	g.Log().Info(ctx, "ListenerManager Started!")
	return nil
}

func (m *manager) Close(ctx context.Context) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for id, listener := range m.savers {
		listener.Close()
		delete(m.savers, id)
	}
	global := utils.GetGlobal(ctx)
	global.WaitGroup.Done()
	g.Log().Info(ctx, "ListenerManager Closed!")
}

func (m *manager) AddListener(ctx context.Context, live lives.Live) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.savers[live.GetLiveId()]; ok {
		g.Log().Warning(ctx, "this live has a listener")
		return nil
	}
	listener := NewListener(ctx, live)
	m.savers[live.GetLiveId()] = listener
	return listener.Start()
}

func (m *manager) RemoveListener(ctx context.Context, liveId int) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	listener, ok := m.savers[liveId]
	if !ok {
		g.Log().Warning(ctx, "this live has not a listener")
		return nil
	}
	listener.Close()
	delete(m.savers, liveId)
	return nil
}

func (m *manager) GetListener(ctx context.Context, liveId int) (Listener, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	listener, ok := m.savers[liveId]
	if !ok {
		return nil, gerror.Newf("this live has a listener")
	}
	return listener, nil
}

func (m *manager) HasListener(ctx context.Context, liveId int) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.savers[liveId]
	return ok
}
