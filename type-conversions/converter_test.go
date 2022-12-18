package typeconversions

import (
	"testing"

	"gotest.tools/assert"
)

func TestConverter(t *testing.T) {
	t.Run("ShouldBeAbleToConvertFloatToInteger", func(t *testing.T) {
		want := 10

		got := FloatToInt(10)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToConvertFloat64ToFloat32", func(t *testing.T) {
		want := float32(10.2)

		got := Float64ToFloat32(10.2)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToConvertIntegerToBool", func(t *testing.T) {
		want := true

		got := IntToBool(10)

		assert.Equal(t, want, got)
	})
}
