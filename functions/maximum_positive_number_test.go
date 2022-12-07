package functions

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnMaximumPositiveNumberInASlice(test *testing.T) {
	numbers := []int{1, 2, -3, -4, -5, 6, 7, 8, -8, 10, 9, -10}
	want := 10

	got := MaximumPositiveNumber(numbers)

	assert.Equal(test, want, got)
}
