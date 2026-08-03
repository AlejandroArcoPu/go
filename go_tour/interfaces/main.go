package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Miau"
}

func (d Dog) Speak() string {
	return "Woff"
}

func main() {
	var s Speaker
	c := Cat{"Suse"}
	d := Dog{"Curro"}
	s = c
	fmt.Println(s.Speak())
	s = d
	fmt.Println(d.Speak())
}
