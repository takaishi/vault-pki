package root

import (
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
	"strings"
)

func GenerateFlags() []cli.Flag {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "pki",
		},
		cli.StringFlag{
			Name: "common_name",
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
	}

	return flags
}

func Generate(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"common_name":          c.String("common_name"),
		"alt_names":            strings.Join(c.StringSlice("alt_names"), ","),
		"ip_sans":              strings.Join(c.StringSlice("ip_sans"), ","),
		"ttl":                  c.String("ttl"),
		"key_bites":            "4096",
		"exclude_cn_from_sans": true,
	}
	_, err = client.Logical().Write(fmt.Sprintf("/%s/root/generate/internal", c.String("pki")), data)
	if err != nil {
		return err
	}

	return nil
}
