package dates

import (
	"strings"
	"time"
)

func Year(current time.Time) string {
	return strings.Split(current.Format("01-02-2006"), "-")[2]
}

func Day(current time.Time) string {
	return current.Format("Monday")
}

func Month(current time.Time) time.Month {
	return current.Month()
}
