package deferstrings

import (
	"fmt"
	"io"
)

func WriteStrings(writer io.Writer) {
	defer fmt.Fprint(writer, "this is defer string!!")

	fmt.Fprint(writer, "This is first string, ")

	fmt.Fprint(writer, "this is second string, ")
}
