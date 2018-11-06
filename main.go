package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BytemarkHosting/bytemark-client/cmd/bytemark/app/auth"
	"github.com/BytemarkHosting/bytemark-client/cmd/bytemark/config"
	bmcli "github.com/BytemarkHosting/bytemark-client/lib"

	bytemark "github.com/BytemarkHosting/docker-machine-driver-bytemark/driver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	doAuthenticate := flag.Bool("authenticate", false, "generate a BYTEMARK_AUTH_TOKEN")
	flag.Parse()
	if *doAuthenticate {
		err := authenticate()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
	plugin.RegisterDriver(bytemark.NewDriver("", ""))
}

func authenticate() error {
	tmpDir, err := ioutil.TempDir("", "docker-machine-driver-bytemark")
	if err != nil {
		return fmt.Errorf("Couldn't create a temporary directory: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	conf, err := config.New(tmpDir)
	if err != nil {
		return fmt.Errorf("Couldn't setup config: %s", err)
	}

	client, err := bmcli.New()
	if err != nil {
		return fmt.Errorf("Couldn't create a bytemark client: %s", err)
	}

	err = auth.NewAuthenticator(client, conf).Authenticate()
	if err != nil {
		return fmt.Errorf("Couldn't authenticate: %s", err)
	}
	fmt.Println(conf.GetIgnoreErr("token"))
	return nil
}
