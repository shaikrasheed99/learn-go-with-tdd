package maps

import (
	"errors"
)

var ErrPlayerNotFound = errors.New("player is not present in score board")

func TotalScore(scoreBoard map[string]int) int {
	sum := 0

	for key := range scoreBoard {
		sum += scoreBoard[key]
	}

	return sum
}

func HighestScored(scoreBoard map[string]int) string {
	player, highestScore := "", 0

	for key, value := range scoreBoard {
		if value > highestScore {
			highestScore = value
			player = key
		}
	}

	return player
}

func IsPresentInScoreBoard(scoreBoards map[string]int, player string) (bool, error) {
	_, isPresent := scoreBoards[player]

	if !isPresent {
		return false, ErrPlayerNotFound
	}

	return true, nil
}
