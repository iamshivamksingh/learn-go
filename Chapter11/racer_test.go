package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// net/http to make HTTP calls
// net/http/httptest to help us test them
// goroutines
// select to synchronise processes.

func TestRaceer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// Creating a server which is listing on a open port and providing the response after 20 ms
		slowServer := makeDelayedServer(20 * time.Microsecond)
		// Creating a server which is listing on a open port & provides response quickly
		fastServer := makeDelayedServer(0 * time.Microsecond)

		// defer will call the close function at the end of the containing function
		// basically used for cleaning up resources like closing the server connection, closing a file etc
		defer slowServer.Close()
		defer fastServer.Close()

		// This will provide the url with port without trailing slash
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("didn't expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't repsond within the specific time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Microsecond)
		w.WriteHeader(http.StatusOK)
	}))
}
