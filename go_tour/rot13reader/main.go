package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err != nil {
		return 0, io.EOF
	}
	for i := 0; i < n; i++ {
		p[i] = p[i] + 13
	}

	return len(p), nil
}

func main() {
	s := strings.NewReader("Hello")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
