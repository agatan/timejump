// Package timejump allows you to mock `time.Now`.
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

// Activate enables timejump's mocking functionality.
// Call this at the top of test functions.
func Activate() {
	activeMu.Lock()
	active = true
	traveledTime = nil
	traveledAt = time.Time{}
	timeScale = 1
	location = nil
}

// Deactivate disables timejump's mocking functionality.
// Call this at the bottom of test functions (or use `defer` at the top).
func Deactivate() {
	active = false
	traveledTime = nil
	traveledAt = time.Time{}
	timeScale = 1
	location = nil
	activeMu.Unlock()
}

// Move sets the location of the time generated by `Now`.
func Move(loc *time.Location) {
	checkActive()
	location = loc
}

// Jump jumps to the time t.
func Jump(t time.Time) {
	checkActive()
	traveledAt = time.Now()
	traveledTime = &t
}

// Scale sets the scale of the world speed.
func Scale(n int) {
	checkActive()
	now := Now()
	timeScale = n
	traveledTime = &now
	traveledAt = now
}

// Stop stops the world (equal to `Scale(0)`).
func Stop() {
	Scale(0)
}

// Now returns the current time.
// If timejump is activated, it returns fake time calculated based on fake values set by `Jump`, `Move`, and `Scale`.
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
