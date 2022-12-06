package ranges

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnTotalPrice(test *testing.T) {
	want := 10

	got := Sum(4, 3, 1, 2)

	assert.Equal(test, want, got)
}
