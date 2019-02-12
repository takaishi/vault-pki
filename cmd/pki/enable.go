package pki

import (
	"github.com/hashicorp/vault/api"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func EnableFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "path",
		},
		cli.StringFlag{
			Name: "description",
		},
	}

	return flags
}

func Enable(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	mountInput := &api.MountInput{
		Type:        "pki",
		Description: c.String("description"),
		Config: api.MountConfigInput{
			DefaultLeaseTTL: "87600h",
			MaxLeaseTTL:     "87600h",
		},
	}
	return client.Sys().Mount(c.String("path"), mountInput)
}
