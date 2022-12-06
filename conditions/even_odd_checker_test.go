package conditions

import (
	"testing"

	"gotest.tools/assert"
)

func TestEvenOddChecker(test *testing.T) {

	want := true
	n := 10

	got := IsEven(n)

	assert.Equal(test, want, got)
}
