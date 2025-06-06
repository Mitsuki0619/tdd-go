package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCollectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	repeated := Repeat("a", 10);
	expected := "aaaaaaaaaa";

	assertCollectMessage(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 10);
	}
}