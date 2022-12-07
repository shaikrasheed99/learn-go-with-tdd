package variadicfunctions

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnTotalSumOfTemperatures(test *testing.T) {
	want := 14.100000000000001

	got := Total(1.5, 2, -3, -4, -5.2, 6, 7, 8.8, -8, 10, 9, -10)

	assert.Equal(test, want, got)
}
