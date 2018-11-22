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
			Name:   "bytemark-zone",
			Usage:  "Zone",
			Value:  defaultZone,
			EnvVar: "BYTEMARK_ZONE",
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
	}
}

func (d *Driver) setServerSpecFromFlags(flags drivers.DriverOptions) {
	d.Spec.VirtualMachine.Cores = flags.Int("bytemark-cores")
	d.Spec.VirtualMachine.Memory = flags.Int("bytemark-memory")
	d.Spec.VirtualMachine.ZoneName = flags.String("bytemark-zone")
	d.Spec.Discs = brain.Discs{brain.Disc{
		Size:         flags.Int("bytemark-disk-size") * 1024,
		StorageGrade: flags.String("bytemark-disk-grade"),
	}}

	// kind of unrelated, but now is as good a time as any
	d.Spec.VirtualMachine.Name = d.vmName().VirtualMachine
}

// SetConfigFromFlags initializes the driver based on the command line flags.
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) (err error) {
	d.setServerSpecFromFlags(flags)

	return nil
}
