package bytemark

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BytemarkHosting/bytemark-client/lib"
	"github.com/BytemarkHosting/bytemark-client/lib/brain"
	"github.com/docker/machine/libmachine/log"
)

const (
	defaultZone      = "york"
	defaultMemory    = 1024
	defaultCores     = 1
	defaultDiskGrade = "sata"
	defaultDiskSize  = 25
	defaultName      = "docker-machine"
	defaultSSHKey    = ""

	defaultUser     = ""
	defaultPass     = ""
	default2FAToken = ""
	defaultYubikey  = ""
)

func (d *Driver) vmName() lib.VirtualMachineName {
	if d.vmNameCached != nil {
		return *d.vmNameCached
	}
	vmn, _ := lib.ParseVirtualMachineName(d.MachineName)
	if vmn.VirtualMachine == "" {
		vmn.VirtualMachine = defaultName
	}
	d.vmNameCached = &vmn
	return vmn
}

func (d *Driver) getToken() (string, error) {
	if d.token != "" {
		return d.token, nil
	}
	d.token = os.Getenv("BYTEMARK_AUTH_TOKEN")
	if d.token != "" {
		return d.token, nil
	}

	log.Debug("Trying to read bytemark auth token from docker-machine store")

	tokenBytes, err := ioutil.ReadFile(d.ResolveStorePath("token"))

	if err == nil {
		d.token = string(tokenBytes)
		return d.token, nil
	}
	path := filepath.Join(os.Getenv("HOME"), ".bytemark", "token")
	log.Debugf("Trying to read bytemark auth token from %s", path)
	tokenBytes, err = ioutil.ReadFile(path)

	if err == nil {
		d.token = string(tokenBytes)
		return d.token, nil
	}

	err = fmt.Errorf("Couldn't find a bytemark auth token. Please set BYTEMARK_AUTH_TOKEN")
	return "", err
}

func (d *Driver) getClient() (lib.Client, error) {
	if d.client != nil {
		return d.client, nil
	}
	token, err := d.getToken()
	if err != nil {
		return nil, fmt.Errorf("Could not find a token: %s. Please set BYTEMARK_AUTH_TOKEN environment variable", err)
	}
	client, err := lib.New()
	if err != nil {
		return nil, fmt.Errorf("Could not create client: %s", err)
	}

	err = client.AuthWithToken(token)
	if err != nil {
		return nil, fmt.Errorf("Could not authenticate: %s", err)
	}
	d.client = client
	return client, nil
}

func (d *Driver) getVirtualMachine(useCache bool) (vm brain.VirtualMachine, err error) {
	if useCache && d.vmCached != nil {
		return *d.vmCached, nil
	}
	vmName := d.vmName()
	client, err := d.getClient()
	if err != nil {
		return
	}
	vm, err = client.GetVirtualMachine(vmName)
	d.vmCached = &vm
	return
}
