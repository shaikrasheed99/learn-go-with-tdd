package typeinference

type Ordered interface {
	~int | ~float64 | ~string
}

func Sum[T Ordered](first T, second T) T {
	return first + second
}
