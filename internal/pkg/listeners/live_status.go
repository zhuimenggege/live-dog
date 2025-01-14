package listeners

const (
	statusToTrueEvt uint8 = 1 << iota
	statusToFalseEvt
	nameChangedEvt
)

type status struct {
	roomName   string
	anchor     string
	roomStatus bool
}

func (s status) Diff(that status) (res uint8) {
	if !s.roomStatus && that.roomStatus {
		res |= statusToTrueEvt
	}
	if s.roomStatus && !that.roomStatus {
		res |= statusToFalseEvt
	}
	if s.roomStatus && that.roomStatus && (s.roomName != that.roomName || s.anchor != that.anchor) {
		res |= nameChangedEvt
	}
	return res
}
