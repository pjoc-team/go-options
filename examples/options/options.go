package options

//go:generate go run ../../ -type eoptions
type eoptions struct {
	Name string
	// Duration time.Duration
	// Config *config.AConfig
}
