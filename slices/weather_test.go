package slices

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnTotalSumOfWeather(test *testing.T) {
	want := 27.2
	report := []float64{5, 1.1, 1.1, 5, 5, 10}

	got := Sum(report)

	assert.Equal(test, want, got)
}
