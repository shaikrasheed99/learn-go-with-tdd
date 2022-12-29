package selectstatement

import "time"

func isPlayerOneReady(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Player one is ready!!"
}

func isPlayerTwoReady(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "Player two is ready!!"
}

func WhoIsReady(playerOne string, playerTwo string) string {
	playerOneChannel := make(chan string)
	playerTwoChannel := make(chan string)

	go isPlayerOneReady(playerOneChannel)
	go isPlayerTwoReady(playerTwoChannel)

	select {
	case <-playerOneChannel:
		return playerOne
	case <-playerTwoChannel:
		return playerTwo
	}
}
