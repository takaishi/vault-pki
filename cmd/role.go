package cmd

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func RoleSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name: "list",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "pki",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := vault.NewClient()
				if err != nil {
					return err
				}

				secret, err := client.Logical().List(fmt.Sprintf("%s/roles", c.String("pki")))
				if err != nil {
					return err
				}
				for _, role := range secret.Data["keys"].([]interface{}) {
					fmt.Println(role)
				}

				return nil
			},
		},
	}

	return commands
}
