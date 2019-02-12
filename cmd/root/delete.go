package root

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func DeleteFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "pki",
		},
	}

	return flags
}

func Delete(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	_, err = client.Logical().Delete(fmt.Sprintf("/%s/root", c.String("pki")))
	if err != nil {
		return err
	}

	return nil
}
