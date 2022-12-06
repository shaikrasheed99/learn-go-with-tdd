package iteration

import (
	"testing"

	"gotest.tools/assert"
)

func TestNaturalNumbers(test *testing.T) {
	test.Run("Should be able to return sum of n natural numbers using traditional for loop", func(t *testing.T) {
		want := 15
		n := 5

		got := SumOfNaturalNumbersUsingTraditionalForLoop(n)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return sum of n natural numbers using for as a while loop", func(t *testing.T) {
		want := 21
		n := 6

		got := SumOfNaturalNumbersUsingForAsAWhileLoop(n)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to stop the infinite loop with a break condition", func(t *testing.T) {
		want := 100
		n := 100

		got := StopIterationsAt(n)

		assert.Equal(t, want, got)
	})
}
