package bytemark

import (
	"strings"

	"github.com/BytemarkHosting/bytemark-client/lib"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	auth3 "gitlab.bytemark.co.uk/auth/client"
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
		mcnflag.StringSliceFlag{
			Name:  "bytemark-disk",
			Usage: "Spec for a disk in [label:]<size>:<grade> e.g. 25:sata",
			Value: strings.Split(defaultDiscSpec, ","),
		},
		mcnflag.StringFlag{
			Name:   "bytemark-server",
			Usage:  "Name for the Bytemark Cloud server that will be created - can be a name.group.account",
			EnvVar: "BYTEMARK_NAME",
			Value:  defaultName,
		},
		mcnflag.StringFlag{
			Name:   "bytemark-ssh-key",
			Usage:  "Path to your SSH private key (will be created if not specified)",
			EnvVar: "BYTEMARK_SSH_KEY",
			Value:  defaultSSHKey,
		},
		mcnflag.StringFlag{
			Name:   "bytemark-user",
			Usage:  "Username to log in as",
			Value:  defaultUser,
			EnvVar: "BYTEMARK_USER",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-pass",
			Usage:  "Password to log in with",
			Value:  defaultPass,
			EnvVar: "BYTEMARK_PASS",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-2fa-token",
			Usage:  "2fa token to log in with",
			Value:  default2FAToken,
			EnvVar: "BYTEMARK_2FA_TOKEN",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-yubikey",
			Usage:  "Yubikey to log in with",
			Value:  defaultYubikey,
			EnvVar: "BYTEMARK_YUBIKEY",
		},
		mcnflag.StringFlag{
			Name:   "bytemark-auth-token",
			Usage:  "Auth token to log in with bytemark-user, pass, 2fa-token and yubikey",
			Value:  defaultAuthToken,
			EnvVar: "BYTEMARK_AUTH_TOKEN",
		},
	}
}

func (d *Driver) setClientFromFlags(flags drivers.DriverOptions) {
	token := flags.String("bytemark-auth-token")
	d.client, err = lib.New()
	if err != nil {
		return
	}

	if token != "" {
		err = d.client.AuthWithToken(token)
	} else {
		twoFactorToken := flags.String("bytemark-2fa-token")
		yubikey := flags.String("bytemark-yubikey")

		credents := auth3.Credentials{
			"username": flags.String("bytemark-user"),
			"password": flags.String("bytemark-password"),
		}
		if twoFactorToken != "" {
			credents["2fa"] = twoFactorToken
		}
		if yubikey != "" {
			credentis["yubikey"] = yubikey
		}
		err = d.client.AuthWithCredentials(credents)
	}
	return
}

func (d *Driver) setServerSpecFromFlags(flags drivers.DriverOptions) {

}

// SetConfigFromFlags initializes the driver based on the command line flags.
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) (err error) {

	return nil
}
