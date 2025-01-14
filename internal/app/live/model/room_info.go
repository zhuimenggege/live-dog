package model

import "github.com/shichen437/live-dog/internal/app/live/model/entity"

type RoomInfo struct {
	*entity.RoomInfo
	Recording bool `json:"recording" description:"录制状态"`
}
