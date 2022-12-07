package tabledriventests

func WhichNumber(number int) string {
	if number > 0 {
		return "Positive"
	} else if number < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}
