package goroutines

import (
	"sync"
)

var wg sync.WaitGroup

func donateApples(donationList *[]string, numberOfApples int) {
	defer wg.Done()
	for apple := 0; apple < numberOfApples; apple++ {
		*donationList = append(*donationList, "apple")
	}
}
func donateBananas(donationList *[]string, numberOfBananas int) {
	for banana := 0; banana < numberOfBananas; banana++ {
		*donationList = append(*donationList, "banana")
	}
}

func DonateApplesAndBananas(numberOfApples int, numberOfBananas int) []string {
	donationList := []string{}
	wg.Add(1)

	go donateApples(&donationList, numberOfApples)
	donateBananas(&donationList, numberOfBananas)

	wg.Wait()

	return donationList
}
