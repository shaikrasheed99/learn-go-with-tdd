package conditions

import (
	"math/rand"
	"testing"

	"gotest.tools/assert"
)

func TestDiceValue(test *testing.T) {
	test.Run("Should be able to return in range of dice values", func(t *testing.T) {
		want := "In range!"
		value := rand.Intn(6)

		got := WhichValueOfDice(value)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return not in range of dice values", func(t *testing.T) {
		want := "Not in range!"
		max := 100
		min := 6
		value := rand.Intn(max-min) + min

		got := WhichValueOfDice(value)

		assert.Equal(t, want, got)
	})
}
