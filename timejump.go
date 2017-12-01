package timejump

import (
	"sync"
	"time"
)

var (
	active       bool
	activeMu     sync.Mutex
	traveledTime *time.Time
	traveledAt   time.Time
	timeScale    int
	location     *time.Location
)

func Activate() {
	activeMu.Lock()
	active = true
	traveledTime = nil
	traveledAt = time.Time{}
	timeScale = 1
	location = nil
}

func Deactivate() {
	active = false
	traveledTime = nil
	traveledAt = time.Time{}
	timeScale = 1
	location = nil
	activeMu.Unlock()
}

func Move(loc *time.Location) {
	checkActive()
	location = loc
}

func Jump(t time.Time) {
	checkActive()
	traveledAt = time.Now()
	traveledTime = &t
}

func Scale(n int) {
	checkActive()
	now := Now()
	timeScale = n
	traveledTime = &now
	traveledAt = now
}

func Stop() {
	Scale(0)
}

func Now() time.Time {
	if !active {
		return time.Now()
	}
	var t time.Time
	if traveledTime != nil {
		t = traveledTime.Add(time.Now().Sub(traveledAt) * time.Duration(timeScale))
	} else {
		t = time.Now()
	}
	if location != nil {
		t = t.In(location)
	}
	return t
}

func checkActive() {
	if !active {
		panic("timegop is not activated; call timegop.Activate() at first")
	}
}
