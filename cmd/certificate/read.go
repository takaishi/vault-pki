package certificate

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func ReadCertificateFlag() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "serial",
			Usage: "Srial of the key to read.",
		},
	}

	return flag
}

func ReadCertificate(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	secret, err := client.Logical().Read(fmt.Sprintf("%s/cert/%s", c.String("pki"), c.String("serial")))
	if err != nil {
		return err
	}
	fmt.Println(secret.Data["certificate"].(string))

	return nil

}
