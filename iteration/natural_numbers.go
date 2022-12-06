package iteration

func SumOfNaturalNumbersUsingTraditionalForLoop(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func SumOfNaturalNumbersUsingForAsAWhileLoop(n int) int {
	sum := 0
	i := 1

	for i <= n {
		sum += i
		i += 1
	}

	return sum
}

func StopIterationsAt(n int) int {
	counter := 0

	for {
		counter++
		if counter == n {
			break
		}
	}

	return counter
}
