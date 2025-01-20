// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiveHistory is the golang structure for table live_history.
type LiveHistory struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	LiveId    int         `json:"liveId"    orm:"live_id"    description:"直播ID"`
	StartTime *gtime.Time `json:"startTime" orm:"start_time" description:"直播开始时间"`
	EndTime   *gtime.Time `json:"endTime"   orm:"end_time"   description:"直播结束时间"`
	Duration  float64     `json:"duration"  orm:"duration"   description:"直播时长"`
}
