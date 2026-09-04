package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	t.Run("should return data", func(t *testing.T) {
		want := "Hello, World!\n"
		handler := hello(RealSleeper{duration: 1, timeout: 3})

		request := httptest.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		body, _ := io.ReadAll(response.Body)

		if string(body) != want {
			t.Errorf("got %s, want %s", string(body), want)
		}
	})
	t.Run("should cancel failure if request is cancelled", func(t *testing.T) {
		handler := hello(RealSleeper{duration: 2, timeout: 1})

		request := httptest.NewRequest(http.MethodGet, "/hello", nil)
		ctx, ctxFunc := context.WithCancel(request.Context())
		time.AfterFunc(time.Millisecond, ctxFunc)
		request = request.WithContext(ctx)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		if response.Code != 500 {
			t.Errorf("an 500 error should be thrown, but got: %d", response.Code)
		}
	})
}
