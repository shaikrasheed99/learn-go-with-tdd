package weather

func Sum(report []float64) float64 {
	sum := 0.0

	for i := 0; i < len(report); i++ {
		sum += report[i]
	}

	return sum
}
