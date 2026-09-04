package myexample

import (
	"io"
	"net/http/httptest"
	"testing"
)

type FakeSleep struct{}

func (f FakeSleep) Sleep() {}

func TestServer(t *testing.T) {
	t.Run("hello returns data", func(t *testing.T) {
		want := "Hello, World"
		handler := hello(FakeSleep{})

		server := httptest.NewServer(handler)
		response, _ := server.Client().Get(server.URL + "/hello")
		body, _ := io.ReadAll(response.Body)

		if string(body) != want {
			t.Errorf("got %s, want %s", string(body), want)
		}
	})
}
