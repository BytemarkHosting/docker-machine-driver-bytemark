package bytemark

import (
	"fmt"
	"os"

	"github.com/docker/machine/libmachine/mcnutils"
)

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

func (d *Driver) GetSSHKeyPath() string {
	d.SSHKeyPath = d.ResolveStorePath("id_rsa")
	return d.SSHKeyPath
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {
	// i am sorry i think
	return "root"
}

func copySSHKey(src, dst string) error {
	if err := mcnutils.CopyFile(src, dst); err != nil {
		return fmt.Errorf("unable to copy ssh key: %s", err)
	}

	if err := os.Chmod(dst, 0600); err != nil {
		return fmt.Errorf("unable to set permissions on the ssh key: %s", err)
	}

	return nil
}
