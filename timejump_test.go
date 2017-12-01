package timejump

import (
	"math/rand"
	"testing"
	"time"
)

func TestJump(t *testing.T) {
	t.Parallel()

	Activate()
	defer Deactivate()

	now := time.Now()
	future := now.AddDate(1, 0, 0)

	Jump(future)

	if Now().Before(future) {
		t.Error("time travel is failed")
	}
}

func TestStop(t *testing.T) {
	t.Parallel()

	Activate()
	defer Deactivate()

	Stop()
	stopped := Now()
	for i := 0; i < 10; i++ {
		if Now() != stopped {
			t.Error("time of the world is not stopped")
		}
	}
}

func TestScale(t *testing.T) {
	t.Parallel()

	Activate()
	defer Deactivate()

	start := Now()
	Scale(int(time.Second / time.Millisecond))
	time.Sleep(1 * time.Millisecond)

	if Now().Before(start.Add(time.Second)) {
		t.Errorf("time scale doesn't work well: start is %v, now is %v", start, Now())
	}
}

func TestMove(t *testing.T) {
	t.Parallel()

	Activate()
	defer Deactivate()

	moveZone, moveOffset := "test/zone", rand.Int()

	Move(time.FixedZone(moveZone, moveOffset))
	zone, offset := Now().Zone()
	if zone != moveZone {
		t.Errorf("zone move is failed: expected zone is %v, but got %v", moveZone, zone)
	}
	if offset != moveOffset {
		t.Errorf("zone move is failed: expected offset is %v, but got %v", moveOffset, offset)
	}
}
