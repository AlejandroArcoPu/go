package myexample

import (
	"fmt"
	"net/http"
	"time"
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (r RealSleeper) Sleep() {
	time.Sleep(5 * time.Second)
}

func hello(s Sleeper) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s.Sleep()
		fmt.Fprintf(w, "Hello, World")
	})
}

// func main() {
// 	http.ListenAndServe(":8080", hello(RealSleeper{}))
// }
