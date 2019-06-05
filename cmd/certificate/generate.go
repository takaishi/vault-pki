package certificate

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/vault/helper/certutil"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
)

func GenerateCertificateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "role",
			Usage: "Name of role to create the certificate.",
		},
		cli.StringFlag{
			Name:  "common_name",
			Usage: "CN for the certificate.",
		},
		cli.StringFlag{
			Name:  "organization",
			Usage: "O for the certificate.",
		},
		cli.StringSliceFlag{
			Name:  "alt_names",
			Usage: "Subject Alternaive Names.",
		},
		cli.StringSliceFlag{
			Name:  "ip_sans",
			Usage: "IP Subject Alternative Names.",
		},
		cli.StringSliceFlag{
			Name:  "uri_sans",
			Usage: "URI Subject Alternative Names.",
		},
		cli.StringFlag{
			Name:  "ttl",
			Usage: "TIme To Live.",
		},
		cli.BoolFlag{
			Name: "exclude_cn_from_sans",
		},
	}

	return flag
}

func GenerateCertificate(c *cli.Context) error {
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

	if len(c.StringSlice("uri_sans")) != 0 {
		data["uri_sans"] = strings.Join(c.StringSlice("uri_sans"), ",")
	}

	if c.String("ttl") != "" {
		data["ttl"] = c.String("ttl")
	}

	data["exclude_cn_from_sans"] = c.Bool("exclude_cn_from_sans")

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

}
