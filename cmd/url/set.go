package url

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func SetFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "pki",
		},
		cli.StringSliceFlag{
			Name: "issuing_certificates",
		},
		cli.StringSliceFlag{
			Name: "crl_distribution_points",
		},
		cli.StringSliceFlag{
			Name: "ocsp_servers",
		},
	}
	return flags
}

func Set(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"issuing_certificates":    c.String("issuing_certificates"),
		"crl_distribution_points": c.StringSlice("crl_distribution_points"),
		"ocsp_servers":            c.StringSlice("ocsp_servers"),
	}
	secret, err := client.Logical().Write(fmt.Sprintf("%s/config/urls", c.String("pki")), data)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", secret)
	return nil
}
