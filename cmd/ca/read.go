package ca

import (
	"bytes"
	"context"
	"fmt"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func ReadFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.StringFlag{
			Name:  "pki",
			Usage: "PKI secret engine name.",
		},
	}

	return flag
}

func Read(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	r := client.NewRequest("GET", fmt.Sprintf("/v1/%s/ca/pem", c.String("pki")))
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	resp, err := client.RawRequestWithContext(ctx, r)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())

	return nil
}
