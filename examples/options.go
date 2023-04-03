package examples

//go:generate go run ../ -type eoptions -option EOption
type eoptions struct {
	Name       string
	FilterFunc FilterFunc
}

type FilterFunc func(t interface{}) bool
