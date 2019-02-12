package cmd

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"os"
	"strings"
	"time"
)

func CertificateSubcommands() []cli.Command {
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

				secret, err := client.Logical().List(fmt.Sprintf("%s/certs", c.String("pki")))
				if err != nil {
					return err
				}
				data := [][]string{}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Organization", "CommonName", "Expire"})

				for _, key := range secret.Data["keys"].([]interface{}) {
					secret, err := client.Logical().Read(fmt.Sprintf("%s/cert/%s", c.String("pki"), key))
					if err != nil {
						return err
					}
					rawCert := secret.Data["certificate"].(string)
					block, _ := pem.Decode([]byte(rawCert))
					if block == nil {
						return errors.New("failed to parse certificate PEM")
					}

					cert, err := x509.ParseCertificate(block.Bytes)
					if err != nil {
						return errors.Wrapf(err, "failed to parse certificate")
					}

					jst := time.FixedZone("Asia/Tokyo", 9*60*60)
					notAfter := cert.NotAfter.In(jst)

					data = append(data, []string{strings.Join(cert.Subject.Organization, ","), cert.Subject.CommonName, notAfter.Format(time.RFC3339)})
				}

				for _, v := range data {
					table.Append(v)
				}
				table.Render()
				return nil
			},
		},
	}

	return commands
}
