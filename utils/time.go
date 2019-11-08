package utils

import (
	// "fmt"
	"time"
)

// TruncateDay 将时间戳截取为日期，输出时间为UTC时间
// t 时间戳
// h 用于补时差
func TruncateDay(t time.Time) time.Time {
	y, m, d := t.Date()
	truncated := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return truncated
}

// TruncateTime 将时间戳截取为时间，输出时间为UTC时间
func TruncateTime(t time.Time) time.Time {
	now := time.Now()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	ns := t.Nanosecond()
	return time.Date(now.Year(), now.Month(), now.Day(), h, m, s, ns, time.UTC)
}

// DaySub 计算时间t1与t2之间日期差
func DaySub(t1, t2 time.Time) int {
	t1Date := TruncateDay(t1)
	t2Date := TruncateDay(t2)
	return int(t1Date.Sub(t2Date).Truncate(time.Hour).Hours() / 24)
}

// LoadBeijingTimeZone 获取北京时区
func LoadBeijingTimeZone() *time.Location {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	return time.FixedZone("Beijing Time", secondsEastOfUTC)
}

// CalculateDurations 计算两个时间点在跨越日期中占的天数
// 例如:
// 		CalculateDuration("2019-12-29 22:00:00", "2019-12-30 01:00:00") = [2, 1]
// 		CalculateDuration("2019-12-29 22:00:00", "2019-12-31 01:00:00") = [2, 24, 1]
// 		CalculateDuration("2019-12-29 22:00:00", "2019-12-29 23:00:00") = [1]
func CalculateDurations(begin, end time.Time) []float64 {
	days := DaySub(end, begin)
	hours := []float64{}
	if days == 0 {
		hours = append(hours, end.Sub(begin).Hours())
		return hours
	}

	for i := 0; i <= days; i++ {
		if i == 0 {
			hours = append(hours, float64(24)-begin.Sub(TruncateDay(begin)).Hours())
		} else if i == days {
			hours = append(hours, end.Sub(TruncateDay(end)).Hours())
		} else {
			hours = append(hours, float64(24))
		}
	}

	return hours
}
