// Code generated by github.com/pjoc-team/go-options.  DO NOT EDIT.

package examples

// ApplyOptionFunc the func of eoptions
type ApplyOptionFunc func(c *eoptions) error

func (f ApplyOptionFunc) apply(c *eoptions) error {
	return f(c)
}

// NewEoptions create Eoptions with options
func NewEoptions(opts ...Option) (eoptions, error) {
	var c eoptions
	err := applyEoptionsOptions(&c, opts...)
	return c, err
}

func newEoptions(opts ...Option) (eoptions, error) {
	var c eoptions
	err := applyEoptionsOptions(&c, opts...)
	return c, err
}

func applyEoptionsOptions(c *eoptions, options ...Option) error {
	for _, o := range options {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

// Option interface Option
type Option interface {
	apply(*eoptions) error
}

// WithName option func
func WithName(o string) ApplyOptionFunc {
	return func(c *eoptions) error {
		c.Name = o
		return nil
	}
}
