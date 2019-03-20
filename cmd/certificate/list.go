package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"math/big"
	"os"
	"strings"
	"time"
)

func ListCertificateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
	}

	return flag
}

func ListCertificate(c *cli.Context) error {
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
	table.SetHeader([]string{"Organization", "CommonName", "Expire", "SerialNumber"})

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

		serial := serialToString(cert.SerialNumber)

		data = append(data, []string{strings.Join(cert.Subject.Organization, ","), cert.Subject.CommonName, notAfter.Format(time.RFC3339), serial})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}

func serialToString(serial *big.Int) string {
	r := []string{}
	splitLen := 2
	runes := []rune(fmt.Sprintf("%x", serial))
	for i := 0; i < len(runes); i += splitLen {
		if i+splitLen < len(runes) {
			r = append(r, string(runes[i:(i+splitLen)]))
		} else {
			r = append(r, string(runes[i:]))
		}
	}
	return strings.Join(r, ":")
}
