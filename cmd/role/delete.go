package role

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func DeleteRoleFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "Name of the role to delete.",
		},
	}
	return flags
}

func DeleteRole(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	_, err = client.Logical().Delete(fmt.Sprintf("%s/role/%s", c.String("pki"), c.String("name")))
	if err != nil {
		return err
	}

	return nil
}
