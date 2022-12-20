package memorymanagement

func CreateSliceUsingNew() *[]int {
	return new([]int)
}

func CreateSliceUsingMake(size int) []int {
	return make([]int, size)
}
