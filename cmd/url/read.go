package url

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func ReadFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "pki",
		},
	}

	return flags
}

func Read(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	secret, err := client.Logical().Read(fmt.Sprintf("%s/config/urls", c.String("pki")))
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", secret.Data)

	return nil
}
