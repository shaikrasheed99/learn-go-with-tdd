package variables

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalculator(test *testing.T) {
	test.Run("Should be able to return sum of two numbers", func(t *testing.T) {
		want := 20

		got := Sum(10, 10)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return subtract of two numbers", func(t *testing.T) {
		want := -1

		got := Subtract(10, 11)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return multiply of two numbers", func(t *testing.T) {
		want := 10.5

		got := Multiply(5, 2.1)

		assert.Equal(t, want, got)
	})
}
