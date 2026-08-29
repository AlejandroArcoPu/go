package main

import (
	"os"
	"time"

	"github.com/AlejandroArcoPu/go/learnGoWithTests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
