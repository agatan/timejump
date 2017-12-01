package timejump_test

import (
	"fmt"
	"time"

	"github.com/agatan/timejump"
)

func Example() {
	timejump.Activate()
	defer timejump.Deactivate()

	future := time.Date(1985, 10, 26, 0, 0, 0, 0, time.UTC)
	timejump.Stop()
	timejump.Jump(future)

	fmt.Println(timejump.Now())

	// Output:
	// 1985-10-26 00:00:00 +0000 UTC
}

func ExampleScale() {
	timejump.Activate()
	defer timejump.Deactivate()

	now := timejump.Now()
	// Time passes 100 times faster.
	timejump.Scale(1000)
	// Sleep just a millisecond.
	time.Sleep(time.Millisecond)

	sub := timejump.Now().Sub(now)

	if sub.Seconds() < 1 {
		panic("1 second have passed while sleeping 1 millisecond in this world.")
	}
}

func ExampleMove() {
	timejump.Activate()
	defer timejump.Deactivate()

	timejump.Move(time.FixedZone("test/zone", 123))

	zone, offset := timejump.Now().Zone()

	fmt.Printf("zone: %v, offset: %v\n", zone, offset)

	// Output:
	// zone: test/zone, offset: 123
}
