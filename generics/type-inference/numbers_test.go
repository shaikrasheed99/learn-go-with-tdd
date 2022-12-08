package typeinference

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumber(t *testing.T) {
	t.Run("Should be able to return sum of two extended int values", func(t *testing.T) {
		type extended int

		first := extended(5)
		second := extended(5)
		want := extended(10)

		got := Sum(first, second)

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to return sum of two extended float values", func(t *testing.T) {
		type extended float64

		first := extended(5.1)
		second := extended(5.2)
		want := extended(10.3)

		got := Sum(first, second)

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to return concatenation of two extended string values", func(t *testing.T) {
		type extended string

		first := extended("Hello")
		second := extended("World")
		want := extended("HelloWorld")

		got := Sum(first, second)

		assert.Equal(t, want, got)
	})
}
