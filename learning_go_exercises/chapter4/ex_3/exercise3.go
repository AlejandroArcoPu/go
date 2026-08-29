package main

import "fmt"

func main() {
	var total int
	for i := range 10 {
		total := total + i // the reason that this doesn't work is that := is defining again and again the variable
		fmt.Println(total)
	}

}
