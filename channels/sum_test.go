package channels

import (
	"testing"

	"gotest.tools/assert"
)

func TestSum(t *testing.T) {
	t.Run("ShouldBeAbleToReturnSumOfTwoSlices", func(t *testing.T) {
		first := []int{1, 2, 3, 4, 5}
		second := []int{6, 7, 8, 9, 10}

		want := 55

		got := Sum(first, second)

		assert.Equal(t, want, got)
	})
}
