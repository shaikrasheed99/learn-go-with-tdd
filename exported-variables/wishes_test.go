package exportedvariables

import (
	"learn-go-with-tdd/wishes"
	"testing"

	"gotest.tools/assert"
)

func TestGreetingCongratulations(t *testing.T) {
	want := "Congratulations"

	got := wishes.Congratulations

	assert.Equal(t, want, got)
}
