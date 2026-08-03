package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func Sum(numbers []int) (sum int) {
	for _, s := range numbers {
		sum += s
	}
	return
}

// this is an unreal example i can use len but I want to try named return
func Counter(persons []Person) (count int) {
	for _, p := range persons {
		fmt.Println(p.Name)
		count++
	}
	return // the named return variable is used if no arguments
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(s))
}
