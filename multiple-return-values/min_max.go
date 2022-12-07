package multiplereturnvalues

import "math"

func MinimumAndMaximumNumbers(numbers []int) (int, int) {
	min, max := math.MaxInt, math.MinInt

	for _, number := range numbers {
		if number > max {
			max = number
		}

		if number < min {
			min = number
		}
	}

	return min, max
}
