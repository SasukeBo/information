package utils

import (
	"testing"
	"time"
)

func TestDaySub(t *testing.T) {
	t1 := time.Date(2019, time.January, 1, 16, 0, 0, 0, time.UTC)
	t2 := time.Date(2019, time.January, 2, 1, 11, 2, 0, time.UTC)
	r := DaySub(t2, t1)
	if r != 1 {
		t.Errorf("DaySub(t2, t1) = %d; want 1", r)
	}
}

func TestCalculateDurations(t *testing.T) {
	t1 := time.Date(2019, time.December, 29, 22, 0, 0, 0, time.UTC)
	t2 := time.Date(2019, time.December, 29, 23, 0, 0, 0, time.UTC)
	t3 := time.Date(2019, time.December, 30, 1, 0, 0, 0, time.UTC)
	t4 := time.Date(2019, time.December, 31, 1, 0, 0, 0, time.UTC)

	hours1 := CalculateDurations(t1, t2)
	hours2 := CalculateDurations(t1, t3)
	hours3 := CalculateDurations(t1, t4)

	if len(hours1) < 1 || hours1[0] != 1 {
		t.Errorf("CalculateDurations(t1, t2) = %v; want [1]", hours1)
	}

	if len(hours2) < 2 || hours2[0] != 2 || hours2[1] != 1 {
		t.Errorf("CalculateDurations(t1, t3) = %v; want [2, 1]", hours2)
	}

	if len(hours3) < 3 || hours3[0] != 2 || hours3[1] != 24 || hours3[2] != 1 {
		t.Errorf("CalculateDurations(t1, t4) = %v; want [2, 24, 1]", hours3)
	}
}

func TestDuration(t *testing.T) {
	t1 := time.Date(2019, time.October, 23, 11, 30, 25, 0, time.UTC)
	t2 := time.Date(2019, time.October, 25, 01, 30, 25, 0, time.UTC)
	hours := CalculateDurations(t1, t2)
	if len(hours) != 1 {
		t.Errorf("CalculateDurations(t1, t2) = %v", hours)
	} else {
		t.Logf("OK, result is %v", hours)
	}
}
