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
		{
			Name: "create",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "pki",
				},
				cli.StringFlag{
					Name: "name",
				},
				cli.StringFlag{
					Name: "organization",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := vault.NewClient()
				if err != nil {
					return err
				}

				data := map[string]interface{}{
					"key_bites":         "4096",
					"max_ttl":           "8760h",
					"allow_any_name":    true,
					"enforce_hostnames": false,
					"organization":      c.String("organization"),
				}

				_, err = client.Logical().Write(fmt.Sprintf("/%s/roles/%s", c.String("pki"), c.String("name")), data)
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	return commands
}
