package strings

import "errors"

var ErrLengthIsNotEqual = errors.New("both names did not have same length")

func UniqueNames(groupOne []string, groupTwo []string) []string {
	uniqueNames := []string{}
	uniqueSet := map[string]int{}

	for _, g1 := range groupOne {
		_, ok := uniqueSet[g1]
		if !ok {
			uniqueSet[g1] = 1
			uniqueNames = append(uniqueNames, g1)
		}
	}

	for _, g2 := range groupTwo {
		_, ok := uniqueSet[g2]
		if !ok {
			uniqueSet[g2] = 1
			uniqueNames = append(uniqueNames, g2)
		}
	}

	return uniqueNames
}

func UniqueCharacters(firstName string, secondName string) (int, error) {
	if len(firstName) != len(secondName) {
		return -1, ErrLengthIsNotEqual
	}

	count := 0

	for index := range firstName {
		if firstName[index] != secondName[index] {
			count += 2
		}
	}

	return count, nil
}
