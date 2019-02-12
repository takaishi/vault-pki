package pki

import (
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func DisableFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "path",
		},
	}

	return flags
}

func Disable(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	return client.Sys().Unmount(c.String("path"))
}
