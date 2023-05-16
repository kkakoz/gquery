package gquery

type Selector[T any] struct {
	table string
}

func NewSelect[T any]() *Selector[T] {
	return &Selector[T]{}
}
