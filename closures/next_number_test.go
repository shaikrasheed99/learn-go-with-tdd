package closures

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnNextFifthNumber(test *testing.T) {
	number := 10
	want := 15

	next := NextFifthNumber(number)
	got := next()

	assert.Equal(test, want, got)
}
