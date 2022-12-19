package clock

import (
	"time"
)

func Hour(current time.Time) string {
	return current.Format("3")
}

func NextDay(current time.Time) string {
	return current.Format("Monday")
}
