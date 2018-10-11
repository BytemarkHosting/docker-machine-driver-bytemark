package bytemark

import (
	"fmt"

	"github.com/BytemarkHosting/bytemark-client/lib"
	"github.com/BytemarkHosting/bytemark-client/lib/brain"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/state"
)

// Driver is a struct compatible with the docker.hosts.drivers.Driver interface.
type Driver struct {
	*drivers.BaseDriver

	Spec      brain.VirtualMachineSpec
	GroupName lib.GroupName
	Server    *brain.VirtualMachine

	client lib.Client
}

const (
	defaultZone     = "york"
	defaultMemory   = 1024
	defaultCores    = 1
	defaultDiscSpec = "25:sata"
	defaultName     = "docker"
	defaultSSHKey   = ""

	defaultUser      = ""
	defaultPass      = ""
	default2FAToken  = ""
	defaultYubikey   = ""
	defaultAuthToken = ""
)

// NewDriver creates a Driver with the specified storePath.
func NewDriver(machineName string, storePath string) *Driver {
	defaultDiscs := brain.Discs{
		brain.Disc{
			Label:        "disk-1",
			StorageGrade: "sata",
			Size:         25600,
		},
	}

	return &Driver{
		Spec: brain.VirtualMachineSpec{
			VirtualMachine: brain.VirtualMachine{
				Cores: defaultCores,
				Discs: defaultDiscs,

				Memory:   defaultMemory,
				ZoneName: defaultZone,
				Name:     defaultName,
			},
			Reimage: &brain.ImageInstall{
				Distribution: "docker",
				RootPassword: "Shohshu9mi9aephahnaigi5l",
			},
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "bytemark"
}

// GetURL returns the URL of the remote docker daemon.
func (d *Driver) GetURL() (string, error) {
	//return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
	return "", fmt.Errorf("not implemented yet")
}

// GetIP returns the IP address of the Bytemark Cloud server
func (d *Driver) GetIP() (string, error) {
	//return ip, nil
	return "", fmt.Errorf("not implemented yet")
}

// GetState returns a docker.hosts.state.State value representing the current state of the host.
func (d *Driver) GetState() (state.State, error) {
	return state.None, fmt.Errorf("not implemented yet")
}

// Start starts or creates an existing Bytemark Cloud server.
func (d *Driver) Start() error {
	return fmt.Errorf("not implemented yet")
}

// Stop stops an existing Bytemark Cloud server.
func (d *Driver) Stop() error {
	return fmt.Errorf("not implemented yet")
}

// Restart restarts a machine which is known to be running.
func (d *Driver) Restart() error {
	return fmt.Errorf("not implemented yet")
}

// Kill stops an existing Bytemark Cloud server.
func (d *Driver) Kill() error {
	return fmt.Errorf("not implemented yet")
}

// Remove deletes the Bytemark Cloud server and its disks.
func (d *Driver) Remove() error {
	return fmt.Errorf("not implemented yet")
}
