package ranges

func Sum(prices ...int) int {
	sum := 0

	for _, price := range prices {
		sum += price
	}

	return sum
}
