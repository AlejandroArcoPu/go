package main

import (
	"fmt"
	"net/http"
)

type Greeting string

func (g Greeting) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, g)
}

func main() {
	http.ListenAndServe(":4000", Greeting("Hello, PETER"))
}
