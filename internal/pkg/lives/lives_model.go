package lives

import (
	"net/url"
	"sync"

	"github.com/shichen437/live-dog/internal/pkg/interfaces"
)

type GLiveModel struct {
	Lives           map[int]Live
	ModelsMap       map[int]*LiveModel
	CookieMap       map[string]string
	EventDispatcher interfaces.Module
	ListenerManager interfaces.Module
	RecorderManager interfaces.Module
	WaitGroup       sync.WaitGroup
}

type LiveModel struct {
	LiveManage LiveManage
	RoomInfo   RoomInfo
}

type LiveManage struct {
	Id           int
	RoomUrl      string
	Interval     int
	Format       string
	EnableNotice int
	MonitorType  int
	MonitorStart string
	MonitorStop  string
}

type RoomInfo struct {
	LiveId      int
	RoomName    string
	Anchor      string
	Platform    string
	Status      int
	LiveStatus  bool
	StreamInfos []*StreamUrlInfo
}

type StreamUrlInfo struct {
	Url                  *url.URL
	Name                 string
	Description          string
	Resolution           int
	Vbitrate             int
	HeadersForDownloader map[string]string
}

func (l *LiveModel) ParseUrl() (sUrl *url.URL, err error) {
	sUrl, err = url.Parse(l.LiveManage.RoomUrl)
	if err != nil {
		return
	}
	return sUrl, nil
}

func (m *LiveModel) GetLiveID() int {
	return m.LiveManage.Id
}
