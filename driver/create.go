package bytemark

import "fmt"

// PreCreateCheck is called to enforce pre-creation steps
func (d *Driver) PreCreateCheck() error {
	return fmt.Errorf("not implemented yet")
}

// Create creates a Bytemark Cloud server acting as a docker host.
func (d *Driver) Create() error {
	return fmt.Errorf("not implemented yet")
}
