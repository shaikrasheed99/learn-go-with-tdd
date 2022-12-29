package selectstatement

import (
	"testing"

	"gotest.tools/assert"
)

func TestPlayers(t *testing.T) {
	t.Run("ShouldBeAbleToReturnNameOfPlayerWhoIsReady", func(t *testing.T) {
		virat := "Virat"
		sachin := "Sachin"

		want := "Virat"

		got := WhoIsReady(virat, sachin)

		assert.Equal(t, want, got)
	})
}
