package main

import "testing"

func TestHello(t *testing.T) {
	assertCollectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Hello関数に名前を入れて実行すると「Hello, {名前}」が返る", func(t *testing.T) {
		got := Hello("Mitsuki")
		want := "Hello, Mitsuki"

		assertCollectMessage(t, got, want)
	})
	t.Run("Hello関数に空文字列を渡して実行すると「Hello, World」が返る", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCollectMessage(t, got, want)
	})
}
