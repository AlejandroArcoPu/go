package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{1, 2}
// 	Scale(&v, 2)
// 	fmt.Println(v)
// }

func (v Vertex) Abs() float64 {
	fmt.Printf("%p\n", &v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	fmt.Printf("%p\n", &v)
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Printf("%p\n", &v)
	fmt.Println(v.Abs())

	v.Scale(2)
	fmt.Println(v)
	fmt.Println(*p)
}
