package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Sleeper interface {
	Response() time.Duration
	Timeout() time.Duration
}

type RealSleeper struct {
	duration time.Duration
	timeout  time.Duration
}

func (r RealSleeper) Response() time.Duration {
	return r.duration * time.Second
}
func (r RealSleeper) Timeout() time.Duration {
	return r.timeout * time.Second
}

func hello(s Sleeper) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		timeout, cancel := context.WithTimeout(req.Context(), s.Timeout())
		defer cancel()

		select {
		case <-timeout.Done():
			http.Error(w, timeout.Err().Error(), http.StatusInternalServerError)
		case <-time.After(s.Response()):
			fmt.Fprint(w, "Hello, World!\n")
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello(RealSleeper{duration: 4, timeout: 1}))
	http.ListenAndServe(":8080", nil)
}
