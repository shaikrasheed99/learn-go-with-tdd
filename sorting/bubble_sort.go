package sorting

import "strings"

func BubbleSort(numbers []int, direction string) []int {
	d := strings.ToLower(direction)
	n := len(numbers)

	if d == "asc" {
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if numbers[j] > numbers[j+1] {
					temp := numbers[j]
					numbers[j] = numbers[j+1]
					numbers[j+1] = temp
				}
			}
		}
	} else if d == "desc" {
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if numbers[j] < numbers[j+1] {
					temp := numbers[j]
					numbers[j] = numbers[j+1]
					numbers[j+1] = temp
				}
			}
		}
	}

	return numbers
}
