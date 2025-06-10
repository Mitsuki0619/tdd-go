package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCollectMessage := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("3 + 4 = 7", func(t *testing.T) {
		sum := Add(3, 4);
		expected := 7;

		assertCollectMessage(t, sum, expected)
	})

	t.Run("5 + 6 = 11", func(t *testing.T) {
		sum := Add(5, 6)
		expected := 11;

		assertCollectMessage(t, sum, expected)
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}