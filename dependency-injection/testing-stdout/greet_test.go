package testingstdout

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestGreet(t *testing.T) {
	message := "Hello world!"
	want := "Hello world!"
	buffer := bytes.Buffer{}

	Greet(&buffer, message)

	got := buffer.String()

	assert.Equal(t, want, got)
}
