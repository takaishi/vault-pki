package intermediate

import (
	"fmt"
	"github.com/hashicorp/vault/helper/certutil"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"io/ioutil"
	"regexp"
	"strings"
)

func SignIntermediateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "csr",
			Usage: "PEM-encoded CSR.",
		},
		cli.StringFlag{
			Name:  "common_name",
			Usage: "CN for the certificate.",
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
	}

	return flag
}

func SignIntermediate(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	data := map[string]interface{}{}

	if c.String("csr") != "" {
		r := regexp.MustCompile("@(.*)")
		if r.MatchString(c.String("csr")) {
			group := r.FindSubmatch([]byte(c.String("csr")))
			path := group[1]
			b, err := ioutil.ReadFile(string(path))
			if err != nil {
				return err
			}
			data["csr"] = string(b)
		}
	}
	if c.String("common_name") != "" {
		data["common_name"] = c.String("common_name")
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

	path := fmt.Sprintf("%s/root/sign-intermediate", c.String("pki"))
	rawData, err := client.Logical().Write(path, data)
	if err != nil {
		return err
	}

	rawCertData := rawData.Data["certificate"]
	certData, err := certutil.ParsePEMBundle(rawCertData.(string))
	if err != nil {
		return err
	}
	cn := certData.Certificate.Subject.CommonName
	err = ioutil.WriteFile(fmt.Sprintf("./%s.pem", cn), []byte(rawCertData.(string)), 0644)
	if err != nil {
		return err
	}

	return nil
}
