package dates

import (
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestDates(t *testing.T) {
	t.Run("ShouldBeAbleToReturnCurrentYearFromTimeString", func(t *testing.T) {
		current := time.Now()

		want := strings.Split(current.Format("01-02-2006"), "-")[2]

		got := Year(current)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnNextYearFromTimeString", func(t *testing.T) {
		current := time.Now().AddDate(1, 0, 0)

		want := strings.Split(current.Format("01-02-2006"), "-")[2]

		got := Year(current)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnCurrentDayFromTimeString", func(t *testing.T) {
		current := time.Now()

		want := current.Format("Monday")

		got := Day(current)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnTomorrowDayFromTimeString", func(t *testing.T) {
		current := time.Now().AddDate(0, 0, 1)

		want := current.Format("Monday")

		got := Day(current)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnCurrentMonthFromTimeString", func(t *testing.T) {
		current := time.Now()

		want := current.Month()

		got := Month(current)

		assert.Equal(t, want, got)
	})
}
