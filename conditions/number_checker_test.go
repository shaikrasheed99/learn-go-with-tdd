package conditions

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumberChecker(test *testing.T) {
	test.Run("Should be able to return number as positive", func(t *testing.T) {
		want := "Positive"
		n := 10

		got := WhichNumber(n)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return number as negative", func(t *testing.T) {
		want := "Negative"
		n := -10

		got := WhichNumber(n)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return number as zero", func(t *testing.T) {
		want := "Zero"
		n := 0

		got := WhichNumber(n)

		assert.Equal(t, want, got)
	})
}
