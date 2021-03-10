package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Sequential racer", func(t *testing.T) {
		fastServer := makeDelayedServer(0 * time.Millisecond)
		slowServer := makeDelayedServer(75 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got := SequentialRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Concurrent racer", func(t *testing.T) {
		{
			fastServer := makeDelayedServer(time.Millisecond * 10)
			slowServer := makeDelayedServer(time.Millisecond * 20)

			defer fastServer.Close()
			defer slowServer.Close()

			got, err := Racer(fastServer.URL, slowServer.URL)

			if err != nil {
				t.Fatalf("did not expect an error but got one %v", err)
			}
			if got != fastServer.URL {
				t.Errorf("got %q, want %q", got, fastServer.URL)
			}
		}
	})

	t.Run("expect error when server doesn't respond within 10s",
		func(t *testing.T) {
			server := makeDelayedServer(time.Millisecond * 25)
			defer server.Close()

			_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond*20)

			if err == nil {
				t.Error("expected an error but did not receive one")
			}
		})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
