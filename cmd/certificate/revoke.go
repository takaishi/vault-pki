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
	if c.String("pki") == "" || c.String("serial") == "" {
		cli.ShowCommandHelp(c, "revoke")
		os.Exit(1)
	}

	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	data := map[string]interface{}{}
	data["serial_number"] = c.String("serial")
	resp, err := client.Logical().Write(fmt.Sprintf("%s/revoke", c.String("pki")), data)
	if err != nil {
		return err
	}

	fmt.Printf("serial %s\n", c.String("serial"))
	for k, v := range resp.Data {
		fmt.Printf("%s %s\n", k, v)
	}

	return nil
}
