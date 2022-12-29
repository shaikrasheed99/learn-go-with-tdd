package channels

func getSum(ch chan int, slice []int) {
	sum := 0

	for _, value := range slice {
		sum += value
	}

	ch <- sum
}

func Sum(first []int, second []int) int {
	ch := make(chan int)

	go getSum(ch, first)
	go getSum(ch, second)

	return (<-ch) + (<-ch)
}
