package shortifblock

func NearestPowerOfTwo(limit int) int {
	power := 1
	for index := 0; index < 100; index++ {
		if pow := power * 2; pow > limit {
			return power
		} else {
			power = pow
		}
	}
	return -1
}
