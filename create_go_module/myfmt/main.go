package myfmt

import "fmt"

type MyStringer interface {
	MyString() string
}

func MyPrintln(i interface{}) {
	if t, ok := i.(MyStringer); ok {
		fmt.Println(t.MyString())
		return
	}
	fmt.Println(i)
}
