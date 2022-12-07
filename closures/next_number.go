package closures

func NextFifthNumber(number int) func() int {
	return func() int {
		number += 5
		return number
	}
}
