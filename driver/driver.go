package bytemark

import (
	"fmt"
	"net"

	"github.com/BytemarkHosting/bytemark-client/lib"
	"github.com/BytemarkHosting/bytemark-client/lib/brain"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/state"
)

// Driver is a struct compatible with the docker.hosts.drivers.Driver interface.
type Driver struct {
	*drivers.BaseDriver

	Spec brain.VirtualMachineSpec

	Token string

	client       lib.Client
	vmCached     *brain.VirtualMachine
	vmNameCached *lib.VirtualMachineName
}

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
				Distribution: "stretch",
				// TODO: generate a random root password.
				RootPassword: "Shohshu9mi9aephahnaigi5l",
			},
		},
		BaseDriver: &drivers.BaseDriver{
			SSHUser:     defaultUser,
			MachineName: machineName,
			StorePath:   storePath,
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "bytemark"
}

// GetURL returns the URL of the remote docker daemon.
func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
}

// GetIP returns the IP address of the Bytemark Cloud server
func (d *Driver) GetIP() (string, error) {
	vm, err := d.getVirtualMachine(true)
	if err != nil {
		return "", err
	}
	return vm.PrimaryIP().String(), nil
}

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {
	return "root"
}

// GetState returns a docker.hosts.state.State value representing the current state of the host.
func (d *Driver) GetState() (state.State, error) {
	vm, err := d.getVirtualMachine(false)
	if err != nil {
		return state.None, err
	}
	// TODO: use the appliance config API - if the server hasn't finished installing call it Starting, if it has call it Started
	// practically speaking it shouldn't make any difference cause docker-machine waits 3 mins (ish) for SSH to come up anyway
	if vm.PowerOn {
		return state.Running, nil
	} else {
		return state.Stopped, nil
	}
}

// Start starts an existing Bytemark Cloud server
func (d *Driver) Start() error {
	client, err := d.getClient()
	if err != nil {
		return err
	}
	return client.StartVirtualMachine(d.vmName())
}

// Stop stops an existing Bytemark Cloud server.
func (d *Driver) Stop() error {
	client, err := d.getClient()
	if err != nil {
		return err
	}
	return client.ShutdownVirtualMachine(d.vmName(), true)
}

// Restart restarts a machine which is known to be running.
func (d *Driver) Restart() error {
	err := d.Stop()
	if err != nil {
		return err
	}
	return d.Start()
}

// Kill stops an existing Bytemark Cloud server.
func (d *Driver) Kill() error {
	client, err := d.getClient()
	if err != nil {
		return err
	}
	return client.StopVirtualMachine(d.vmName())
}

// Remove deletes the Bytemark Cloud server and its disks.
func (d *Driver) Remove() error {
	client, err := d.getClient()
	if err != nil {
		return err
	}
	return client.DeleteVirtualMachine(d.vmName(), true)
}
