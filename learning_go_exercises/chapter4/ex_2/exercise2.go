package ex2

import (
	"fmt"
	"io"
)

func Divisible(number int, out io.Writer) {
	switch {
	case number%2 == 0 && number%3 == 0:
		fmt.Fprintf(out, "Six!")
	case number%2 == 0:
		fmt.Fprintf(out, "Two!")
	case number%3 == 0:
		fmt.Fprintf(out, "Three!")
	default:
		fmt.Fprintf(out, "Never mind")
	}
}
