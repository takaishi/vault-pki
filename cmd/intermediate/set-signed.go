package intermediate

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"io/ioutil"
	"regexp"
)

func SetSignedIntermediateFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
		cli.StringFlag{
			Name:  "certificate",
			Usage: "certificate in the PEM format.",
		},
	}

	return flag
}

func SetSignedIntermediate(c *cli.Context) error {
	data := map[string]interface{}{}
	if c.String("certificate") != "" {
		r := regexp.MustCompile("@(.*)")
		if r.MatchString(c.String("certificate")) {
			group := r.FindSubmatch([]byte(c.String("certificate")))
			path := group[1]
			b, err := ioutil.ReadFile(string(path))
			if err != nil {
				return err
			}
			data["certificate"] = string(b)
		}
	}
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	_, err = client.Logical().Write(fmt.Sprintf("%s/intermediate/set-signed", c.String("pki")), data)
	if err != nil {
		return err
	}

	return nil
}
