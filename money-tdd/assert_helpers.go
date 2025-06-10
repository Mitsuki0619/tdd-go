package money

import "testing"

func assertEqual(t *testing.T, got, want Money) {
	if got.Amount() != want.Amount() {
		t.Errorf("got %d want %d", got.Amount(), want.Amount())
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("it is false")
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Errorf("it is true")
	}
}