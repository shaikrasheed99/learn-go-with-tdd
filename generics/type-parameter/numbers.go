package typeparameter

type Ordered interface {
	int | float64 | string
}

func Sum[T Ordered](values ...T) T {
	var sum T

	for _, value := range values {
		sum += value
	}

	return sum
}
