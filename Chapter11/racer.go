package racer

import (
	"net/http"
)

func Racer(a, b string) string {
	// select lets you wait on multiple channels.
	// The first one to send the value will wins and the case underneath the case is executed.
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

/*
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/

func ping(url string) chan struct{} {
	// chan struct{} is the smallest data type available from the memory perspective
	ch := make(chan struct{})
	// We don't care what type is sent to the channel,
	// we just want to signal we are done and closing the channel works perfectly
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
