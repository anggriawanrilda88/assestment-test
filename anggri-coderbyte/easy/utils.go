package easy

import "time"

func MakeTimestampMilli() int64 {
	return UnixMilli(time.Now())
}
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
