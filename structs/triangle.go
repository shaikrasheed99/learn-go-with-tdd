package structs

type Triangle struct {
	height float64
	base   float64
}

func Area(triangle Triangle) float64 {
	return (0.5 * triangle.base * triangle.height)
}
