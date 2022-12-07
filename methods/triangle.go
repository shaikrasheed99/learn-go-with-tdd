package methods

type Triangle struct {
	height float64
	base   float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
