package memorymanagement

import (
	"testing"

	"gotest.tools/assert"
)

func TestCreateSlice(t *testing.T) {
	t.Run("ShouldBeAbleToCreateASliceWithNewKeyword", func(t *testing.T) {
		want := new([]int)

		got := CreateSliceUsingNew()

		assert.DeepEqual(t, want, got)
	})

	t.Run("ShouldBeAbleToCreateASliceWithMakeKeyword", func(t *testing.T) {
		size := 5
		want := make([]int, size)

		got := CreateSliceUsingMake(size)

		assert.Equal(t, want[0], got[0])
	})
}
