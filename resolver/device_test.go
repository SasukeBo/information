package resolver

import (
	"testing"
	"time"
)

func TestCalculateDuration(t *testing.T) {
	begin := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(2019, time.January, 2, 16, 0, 0, 0, time.UTC)
	t2 := time.Date(2019, time.January, 3, 2, 0, 0, 0, time.UTC)
	index, hours := calculateDuration(begin, t1, t2)
	if index != 1 || len(hours) != 2 {
		t.Errorf("calculateDuration(begin, t1, t2) = %d, %v; want 1, []float64{16, 2}", index, hours)
	}
}
