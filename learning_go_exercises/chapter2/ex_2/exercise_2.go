package main

import "fmt"

func main() {
	const value = 20
	var i int
	var f float64
	i = value
	f = value
	fmt.Println(i, f)
}

// The trick here is use an untyped int that can be assigned to both int and float64
