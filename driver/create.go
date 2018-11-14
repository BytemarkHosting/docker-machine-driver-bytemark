package bytemark

import (
	"io/ioutil"

	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/ssh"
)

// PreCreateCheck is called to enforce pre-creation steps
func (d *Driver) PreCreateCheck() error {
	return nil
}

// Create creates a Bytemark Cloud server acting as a docker host.
// TODO: add a first-boot script to create a non-root user in the docker group and populate their authorized_keys
func (d *Driver) Create() error {
	client, err := d.getClient()
	if err != nil {
		return err
	}
	log.Info("Generating SSH Key")

	if err = ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return err
	}

	sshKey, err := ioutil.ReadFile(d.GetSSHKeyPath() + ".pub")
	if err != nil {
		return err
	}

	log.Info("Creating host...")

	d.Spec.Reimage.PublicKeys = string(sshKey)
	vmn := d.vmName()

	_, err = client.CreateVirtualMachine(vmn.GroupName(), d.Spec)

	return err
}
