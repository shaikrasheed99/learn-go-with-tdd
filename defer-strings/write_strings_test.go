package deferstrings

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestWriteStrings(t *testing.T) {
	t.Run("ShouldBeAbleToWriteStringsWithDefer", func(t *testing.T) {
		want := "This is first string, this is second string, this is defer string!!"
		buffer := bytes.Buffer{}

		WriteStrings(&buffer)

		got := buffer.String()

		assert.Equal(t, want, got)
	})
}
