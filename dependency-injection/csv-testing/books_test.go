package csvtesting

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestCSV(t *testing.T) {
	want := "[Fundamentals of Wavelets Goswami, Jaideva signal_processing Wiley]\n"
	buffer := bytes.Buffer{}
	ReadCSV(&buffer)

	got := buffer.String()

	assert.Equal(t, want, got)
}
