package shortifblock

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrimeChecker(t *testing.T) {
	t.Run("ShouldBeAbleToReturnNearestPowerOfTwo", func(t *testing.T) {
		want := 16
		limit := 20

		got := NearestPowerOfTwo(limit)

		assert.Equal(t, want, got)
	})
}
