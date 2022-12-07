package recursion

func Factorial(number int) int {
	if number <= 0 {
		return 1
	}

	return number * Factorial(number-1)
}
