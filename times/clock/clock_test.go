package clock

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestClock(t *testing.T) {
	t.Run("ShouldBeAbleToReturnCurrentHourIn12HoursFormatFromTimeString", func(t *testing.T) {
		current := time.Now()

		want := current.Format("3")

		got := Hour(current)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnDayAfterTwoDaysFromTimeString", func(t *testing.T) {
		current := time.Now().Add(time.Hour * 48)

		want := current.Format("Monday")

		got := NextDay(current)

		assert.Equal(t, want, got)
	})
}
