package sorting

import (
	"testing"

	"gotest.tools/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Should be able to sort numbers in ascending order", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}
		direction := "aSc"
		want := []int{1, 2, 3, 4, 5}

		got := BubbleSort(numbers, direction)

		assert.DeepEqual(t, want, got)
	})

	t.Run("Should be able to sort numbers in descending order", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		direction := "dEsC"
		want := []int{5, 4, 3, 2, 1}

		got := BubbleSort(numbers, direction)

		assert.DeepEqual(t, want, got)
	})
}
