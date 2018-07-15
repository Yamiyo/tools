package timetool

import (
	"fmt"
	"strings"
	"time"
)

func GetIntervalBefore(srcTime time.Time, interval int, intervalType string) (time.Time, error) {
	return GetInterval(srcTime, -interval, intervalType)
}

func GetIntervalAfter(srcTime time.Time, interval int, intervalType string) (time.Time, error) {
	return GetInterval(srcTime, interval, intervalType)
}

// GetInterval : AddDate normalizes its result in the same way that Date does, so, for example, adding one month to October 31 yields December 1, the normalized form for November 31.
func GetInterval(srcTime time.Time, interval int, intervalType string) (time.Time, error) {
	target := time.Time{}

	switch strings.ToLower(intervalType) {
	case "s", "second":
		target = srcTime.Add(time.Duration(interval) * time.Second)
	case "m", "minute":
		target = srcTime.Add(time.Duration(interval) * time.Minute)
	case "h", "hour":
		target = srcTime.Add(time.Duration(interval) * time.Hour)
	case "d", "day":
		target = srcTime.AddDate(0, 0, interval)
	case "w", "week":
		target = srcTime.AddDate(0, 0, interval*7)
	case "mth", "month":
		target = srcTime.AddDate(0, interval, 0)
	case "y", "year":
		target = srcTime.AddDate(interval, 0, 0)
	default:
		return srcTime, fmt.Errorf("[ %v ] is not interval type.", intervalType)
	}

	return target, nil
}
