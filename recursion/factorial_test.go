package recursion

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnFactorialOfANumber(t *testing.T) {
	number := 5
	want := 120

	got := Factorial(number)

	assert.Equal(t, want, got)
}
