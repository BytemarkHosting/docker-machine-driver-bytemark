package main

import (
	bytemark "github.com/BytemarkHosting/docker-machine-driver-bytemark/driver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(bytemark.NewDriver("", ""))
}
