package maps

import (
	"testing"

	"gotest.tools/assert"
)

func TestCricketTeam(test *testing.T) {
	scoreBoard := map[string]int{
		"Sachin": 200,
		"Virat":  183,
		"Dhoni":  183,
		"Yuvraj": 150,
		"Rohit":  264,
		"Sehwag": 219,
	}

	test.Run("Should be able to return total score of the team", func(t *testing.T) {
		want := 1199

		got := TotalScore(scoreBoard)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return name of highest scored player", func(t *testing.T) {
		want := "Rohit"

		got := HighestScored(scoreBoard)

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to check for a player existance", func(t *testing.T) {
		player := "Virat"
		want := true

		got, error := IsPresentInScoreBoard(scoreBoard, player)

		if error != nil {
			t.Error()
		}

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to return error when we check for not existed player in score board", func(t *testing.T) {
		player := "Jadeja"
		want := false

		got, error := IsPresentInScoreBoard(scoreBoard, player)

		if error == nil {
			t.Error()
		}

		assert.Equal(t, want, got)
		assert.Equal(t, "player is not present in score board", error.Error())
	})
}
