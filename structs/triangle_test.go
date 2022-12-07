package structs

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToGetAreaOfATriangle(t *testing.T) {
	trianlge := Triangle{height: 10, base: 10}
	want := 50.0

	got := Area(trianlge)

	assert.Equal(t, want, got)
}
