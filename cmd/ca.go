package cmd

import (
	"github.com/takaishi/vault-pki/cmd/ca"
	"github.com/urfave/cli"
)

func CASubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:        "read",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#read-ca-certificate",
			Flags:       ca.ReadFlags(),
			Action:      ca.Read,
		},
	}

	return commands
}
