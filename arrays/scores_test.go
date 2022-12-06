package arrays

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnTotalOfScores(test *testing.T) {
	want := 20
	scores := [5]int{10, 0, 2, 3, 5}

	got := Sum(scores)

	assert.Equal(test, want, got)
}
