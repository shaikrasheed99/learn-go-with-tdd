package variadicfunctions

func Total(temperatures ...float64) float64 {
	total := 0.0

	for _, temperature := range temperatures {
		total += temperature
	}

	return total
}
