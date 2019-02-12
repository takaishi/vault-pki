package cmd

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func PkiSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name: "list",
			Action: func(c *cli.Context) error {
				client, err := vault.NewClient()
				if err != nil {
					return err
				}

				mounts, err := client.Sys().ListMounts()
				if err != nil {
					return err
				}

				pkiList := filterPkiOnly(mounts)

				for name, pki := range pkiList {
					fmt.Printf("%s %s %s\n", name, pki.Accessor, pki.Description)
				}

				return nil
			},
		},
		{
			Name: "enable",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "path",
				},
				cli.StringFlag{
					Name: "description",
				},
			},
			Action: func(c *cli.Context) error {
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
			},
		},
		{
			Name: "disable",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "path",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := vault.NewClient()
				if err != nil {
					return err
				}

				return client.Sys().Unmount(c.String("path"))
			},
		},
	}

	return commands
}

func filterPkiOnly(mounts map[string]*api.MountOutput) map[string]*api.MountOutput {
	filtered := map[string]*api.MountOutput{}

	for name, mount := range mounts {
		if mount.Type == "pki" {
			filtered[name] = mount
		}
	}

	return filtered
}
