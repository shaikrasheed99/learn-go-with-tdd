package testingstdout

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, message string) {
	fmt.Fprint(writer, message)
}
