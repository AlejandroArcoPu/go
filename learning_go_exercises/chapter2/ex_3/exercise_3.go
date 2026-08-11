package main

import "fmt"

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	b = b + 1
	smallI = 2_147_483_647
	smallI = smallI + 1
	bigI = 18_446_744_073_709_551_615
	bigI = bigI + 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}

// They are reset to the minimum value because it causes an overflow
