package methods

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToGetAreaOfATriangle(t *testing.T) {
	triangle := Triangle{height: 10, base: 20}
	want := 100.0

	got := triangle.Area()

	assert.Equal(t, want, got)
}
