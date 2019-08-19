package intermediate

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"strings"
)

func GenerateIntermediateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "(Required) PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "type",
			Usage: "(Required) type of the intermediate to create.",
		},
		cli.StringFlag{
			Name:  "common_name",
			Usage: "(Required) CN for the certificate.",
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

func GenerateIntermediate(c *cli.Context) error {
	err := validateGenerateIntermediateFlags(c)
	if err != nil {
		return err
	}

	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	payload := createGenerateIntermediatePayload(c)
	path := fmt.Sprintf("%s/intermediate/generate/%s", c.String("pki"), c.String("type"))

	rawData, err := client.Logical().Write(path, payload)
	if err != nil {
		return err
	}

	csrData := rawData.Data["csr"]
	err = ioutil.WriteFile(fmt.Sprintf("./%s_intermediate_csr.pem", c.String("pki")), []byte(csrData.(string)), 0644)
	if err != nil {
		return err
	}

	privKeyData := rawData.Data["private_key"]
	log.Println(privKeyData.(string))
	err = ioutil.WriteFile(fmt.Sprintf("./%s_intermediate_private_key.pem", c.String("pki")), []byte(privKeyData.(string)), 0644)
	if err != nil {
		return err
	}

	return nil
}

func validateGenerateIntermediateFlags(c *cli.Context) error {
	requiredFlags := []string{"pki", "type", "common_name"}
	for _, item := range requiredFlags {
		if c.String(item) == "" {
			cli.ShowCommandHelp(c, "generate")
			return errors.Errorf("Require flags %s", item)
		}
	}

	return nil
}

func createGenerateIntermediatePayload(c *cli.Context) map[string]interface{} {
	data := map[string]interface{}{}

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

	return data
}
