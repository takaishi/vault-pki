package role

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func ListRoleFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
	}
	return flags
}

func ListRole(c *cli.Context) error {
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
}
