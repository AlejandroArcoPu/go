package main

import "fmt"

func main() {
	var i int
	var f float64
	i = 20
	f = float64(i)
	fmt.Println(i)
	fmt.Println(f)
}

// This will fail because they need to be of the same type a solution is to use float64() cast
