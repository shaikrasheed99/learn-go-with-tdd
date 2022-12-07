package multiplereturnvalues

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnMinimaxAndMaximumNumbersInASlice(test *testing.T) {
	numbers := []int{1, 2, -3, -4, -5, 6, 7, 8, -8, 10, 9, -10}
	wantMin := -10
	wantMax := 10

	gotMin, gotMax := MinimumAndMaximumNumbers(numbers)

	assert.Equal(test, wantMin, gotMin)
	assert.Equal(test, wantMax, gotMax)
}
