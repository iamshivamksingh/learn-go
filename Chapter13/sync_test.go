package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		//counter := NewCounter()
		counter := Counter{}

		// WaitGroup are similar to countdown latches in Java
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		// Creating `wantedCount` numbers goroutines
		// which will calls counter.Inc() and
		// then marking themself Done when finished
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		// Wait() will block until all goroutines have finished
		wg.Wait()

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), 3)
	}

}
