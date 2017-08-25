package console

import (
	"errors"
	"fmt"
)

var writeFn = fmt.Print

func setWriter(w func(a ...interface{}) (int, error)) {
	writeFn = w
}

func wrap(s interface{}, c string) string {
	return fmt.Sprintf("%s%v%s", c, s, ColorReset)
}

// Print to write text with color
func Print(s interface{}, c string) error {
	buf := wrap(s, c)
	_, e := writeFn(buf)
	return e
}

// Println to write text and new line with color
func Println(s interface{}, c string) error {
	buf := fmt.Sprintf("%s\n", wrap(s, c))
	_, e := writeFn(buf)
	return e
}

// Printf to write a formatted text with color
func Printf(a ...interface{}) error {
	switch len(a) {
	case 0:
		return errors.New("no format specified")
	case 1:
		return errors.New("no color specified")
	}
	f := a[0].(string)
	c := a[len(a)-1].(string)
	return Print(fmt.Sprintf(f, a[1:len(a)-1]...), c)
}

// Printfln to write a formatted text and newline with color
func Printfln(a ...interface{}) error {
	switch len(a) {
	case 0:
		return errors.New("no format specified")
	case 1:
		return errors.New("no color specified")
	}
	f := a[0].(string)
	c := a[len(a)-1].(string)
	return Println(fmt.Sprintf(f, a[1:len(a)-1]...), c)
}
