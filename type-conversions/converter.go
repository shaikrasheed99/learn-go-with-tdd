package typeconversions

func FloatToInt(number float32) int {
	return int(number)
}

func Float64ToFloat32(number float64) float32 {
	return float32(number)
}

func IntToBool(number int) bool {
	if number == 0 {
		return false
	} else {
		return true
	}
}
