package model

import "github.com/gogf/gf/os/gtime"

type LiveHistory struct {
	Id        int64       `json:"id"`
	LiveId    int64       `json:"liveId" description:"直播ID"`
	Anchor    string      `json:"anchor" description:"主播名称"`
	StartTime *gtime.Time `json:"startTime" description:"开始时间"`
	EndTime   *gtime.Time `json:"endTime" description:"结束时间"`
	Duration  float64     `json:"duration" description:"时长"`
}
