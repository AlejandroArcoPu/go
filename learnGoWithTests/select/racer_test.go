package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		_, err := ConfigurableRacer(slowUrl, fastUrl, 5*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error and didn't get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
