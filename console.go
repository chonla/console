package console

import (
	"fmt"
)

var writeFn = fmt.Print

func setWriter(w func(a ...interface{}) (int, error)) {
	writeFn = w
}

func wrap(s string, c string) string {
	return fmt.Sprintf("%s%s%s", c, s, ColorReset)
}

// Print to write text with color
func Print(s string, c string) {
	buf := wrap(s, c)
	writeFn(buf)
}

// Println to write text and new line with color
func Println(s string, c string) {
	buf := fmt.Sprintf("%s\n", wrap(s, c))
	writeFn(buf)
}
