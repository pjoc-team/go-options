package examples

//go:generate go run ../ -type eoptions
type eoptions[T any] struct {
	Name       string
	FilterFunc FilterFunc[T]
}

type FilterFunc[T any] func(t T) bool
