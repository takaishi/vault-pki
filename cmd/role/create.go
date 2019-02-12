package role

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func CreateRoleFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "Name of the role to create.",
		},
		cli.StringFlag{
			Name: "organization",
		},
	}

	return flags
}

func CreateRole(c *cli.Context) error {
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
}
