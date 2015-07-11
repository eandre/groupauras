package dtgraph

import (
	"fmt"
	"time"
)

const (
	keepPrevious = 5 * time.Second
	guessFuture  = 30 * time.Second
	resolution   = 1 * time.Second
	recordCount  = int((keepPrevious + guessFuture) / resolution)
)

type Recorder struct {
	points []int
	now    int
}

func NewRecorder() *Recorder {
	// Pre-allocate points with 0 damage
	points := make([]int, recordCount)
	now := getIdx(0, keepPrevious)

	return &Recorder{
		points: points,
		now:    now,
	}
}

func (r *Recorder) Record(dt time.Duration, value int) {
	if dt > guessFuture {
		fmt.Println("Tried to exceed future threshold")
		return
	}

	idx := getIdx(r.now, dt)
	r.points[idx] = value
}

func getIdx(now int, dt time.Duration) int {
	idx := int(dt / resolution)
	return (now + idx) % recordCount
}
