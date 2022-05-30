package iteration

import (
	"testing"
)

func TestRepeating(t *testing.T) {
	t.Run("it should validate older implementation", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("it should validate newer implementation", func(t *testing.T) {
		repeated := RepeatModified("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}

}
