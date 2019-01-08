package bytemark

import (
	"github.com/BytemarkHosting/bytemark-client/lib/brain"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
)

// GetCreateFlags registers the flags this driver adds to
// "docker hosts create"
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:   "bytemark-cluster",
			Usage:  "HTTP(S) root URL of cluster endpoint",
			Value:  "https://uk0.bigv.io",
			EnvVar: "BYTEMARK_CLUSTER_URL",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-zone",
			Usage:  "Zone",
			Value:  defaultZone,
			EnvVar: "BYTEMARK_ZONE",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-token",
			Usage:  "Token for authentication",
			Value:  "",
			EnvVar: "BYTEMARK_TOKEN",
		},
		mcnflag.IntFlag{
			Name:   "bytemark-memory",
			Usage:  "Memory in MiB",
			Value:  defaultMemory,
			EnvVar: "BYTEMARK_MEMORY",
		},
		mcnflag.IntFlag{
			Name:   "bytemark-cores",
			Usage:  "Number of CPU cores",
			Value:  defaultCores,
			EnvVar: "BYTEMARK_CORES",
		},
		mcnflag.IntFlag{
			Name:  "bytemark-disk-size",
			Usage: "Disk size in GiB",
			Value: defaultDiskSize,
		},
		mcnflag.StringFlag{
			Name:  "bytemark-disk-grade",
			Usage: "Disk storage grade (sata / archive)",
			Value: defaultDiskGrade,
		},
		mcnflag.StringFlag{
			Name:  "bytemark-image",
			Usage: "Base image to use",
			Value: "stretch",
		},
	}
}

func (d *Driver) setServerSpecFromFlags(flags drivers.DriverOptions) {
	d.spec.VirtualMachine.Cores = flags.Int("bytemark-cores")
	d.spec.VirtualMachine.Memory = flags.Int("bytemark-memory")
	d.spec.VirtualMachine.ZoneName = flags.String("bytemark-zone")
	d.spec.Discs = brain.Discs{brain.Disc{
		Size:         flags.Int("bytemark-disk-size") * 1024,
		StorageGrade: flags.String("bytemark-disk-grade"),
	}}
	d.spec.Reimage.Distribution = flags.String("bytemark-image")

	// kind of unrelated, but now is as good a time as any
	d.spec.VirtualMachine.Name = d.vmName().VirtualMachine
}

// SetConfigFromFlags initializes the driver based on the command line flags.
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) (err error) {
	d.ClusterURL = flags.String("bytemark-cluster")
	d.setServerSpecFromFlags(flags)
	d.Token = flags.String("bytemark-token")

	return nil
}
