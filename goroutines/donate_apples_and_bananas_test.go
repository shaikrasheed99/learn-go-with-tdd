package goroutines

import (
	"testing"

	"gotest.tools/assert"
)

func TestDonateApplesAndBananas(t *testing.T) {
	t.Run("ShouldBeAbleToDonateFiveApplesAndThreeBananas", func(t *testing.T) {
		numberOfApples := 5
		numberOfBananas := 3
		want := 8

		got := DonateApplesAndBananas(numberOfApples, numberOfBananas)

		assert.Equal(t, want, len(got))
	})
}
