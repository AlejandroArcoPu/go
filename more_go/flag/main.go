package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar int
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.Parse()
	fmt.Println(flagvar)
}
