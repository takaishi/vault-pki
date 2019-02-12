package cmd

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/vault/helper/certutil"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"io/ioutil"
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
		{
			Name: "create",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "pki",
				},
				cli.StringFlag{
					Name: "role",
				},
				cli.StringFlag{
					Name: "common_name",
				},
				cli.StringFlag{
					Name: "organization",
				},
				cli.StringSliceFlag{
					Name: "alt_names",
				},
				cli.StringSliceFlag{
					Name: "ip_sans",
				},
				cli.StringFlag{
					Name: "ttl",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := vault.NewClient()
				if err != nil {
					return err
				}

				data := map[string]interface{}{}

				if c.String("common_name") != "" {
					data["common_name"] = c.String("common_name")
				}
				if c.String("organization") != "" {
					data["organization"] = c.String("organization")
				}
				if len(c.StringSlice("alt_names")) != 0 {
					data["alt_names"] = strings.Join(c.StringSlice("alt_names"), ",")
				}
				if len(c.StringSlice("ip_sans")) != 0 {
					data["ip_sans"] = strings.Join(c.StringSlice("ip_sans"), ",")
				}
				if c.String("ttl") != "" {
					data["ttl"] = c.String("ttl")
				}

				rawCertData, err := client.Logical().Write(fmt.Sprintf("%s/issue/%s", c.String("pki"), c.String("role")), data)
				if err != nil {
					return err
				}
				certData, err := certutil.ParsePKIMap(rawCertData.Data)
				if err != nil {
					return err
				}

				out := &bytes.Buffer{}
				pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: certData.Certificate.Raw})
				err = ioutil.WriteFile(fmt.Sprintf("./%d.pem", certData.Certificate.SerialNumber), out.Bytes(), 0644)
				if err != nil {
					return err
				}

				out2 := &bytes.Buffer{}
				b := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: certData.PrivateKeyBytes}
				pem.Encode(out2, b)
				err = ioutil.WriteFile(fmt.Sprintf("./%d-key.pem", certData.Certificate.SerialNumber), out2.Bytes(), 0644)
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	return commands
}
