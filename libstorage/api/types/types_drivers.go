package types

import gofig "github.com/akutz/gofig/types"

// Driver is the base interface for a libStorage driver.
type Driver interface {

	// Name returns the name of the driver
	Name() string

	// Init initializes the driver.
	Init(ctx Context, config gofig.Config) error
}
