package functions

import (
	"math"
)

func MaximumPositiveNumber(numbers []int) int {
	maximumPositiveNumber := math.MinInt

	for _, number := range numbers {
		if isPositive(number) && number > maximumPositiveNumber {
			maximumPositiveNumber = number
		}
	}

	return maximumPositiveNumber
}

func isPositive(number int) bool {
	return number > 0
}
