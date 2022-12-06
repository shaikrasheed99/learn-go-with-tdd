package conditions

func WhichValueOfDice(value int) string {
	switch {
	case value <= 6:
		return "In range!"
	default:
		return "Not in range!"
	}
}
