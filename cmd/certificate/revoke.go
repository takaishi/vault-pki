package certificate

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func RevokeCertificateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "serial",
			Usage: "Serial number of the certificate to revoke. hyphen-separated or colon-separated.",
		},
	}

	return flag
}

func RevokeCertificate(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	data := map[string]interface{}{}
	_, err = client.Logical().Write(fmt.Sprintf("%s/revoke", c.String("pki")), data)
	if err != nil {
		return err
	}

	return nil
}
