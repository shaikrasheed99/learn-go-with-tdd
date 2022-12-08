package typeparameter

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumber(t *testing.T) {
	t.Run("Should be able to return sum of integer values", func(t *testing.T) {
		want := 10

		got := Sum(1, 2, 3, 4)

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to return sum of float values", func(t *testing.T) {
		want := 10.1

		got := Sum(1.0, 2.0, 3.0, 4.1)

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to return concatenation of string values", func(t *testing.T) {
		want := "Hello World"

		got := Sum("Hello", " ", "World")

		assert.Equal(t, want, got)
	})
}
