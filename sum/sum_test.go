package sum

import (
	"testing"
)

func TestSumI(t *testing.T) {
	assertCollectMessage := func(t testing.TB, result, expected int, numbers []int) {
		t.Helper()
		if result != expected {
			t.Errorf("expected %d but result %d given %v", expected, result, numbers)
		}
	}

	t.Run("[1,2,3,4,5]の合計は15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15
		assertCollectMessage(t, result, expected, numbers)
	})

	t.Run("[]の合計は0", func(t *testing.T) {
		empty := []int{}
		result := Sum(empty)
		expected := 0

		assertCollectMessage(t, result, expected, empty)
	})
}
