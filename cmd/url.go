package cmd

import (
	"github.com/takaishi/vault-pki/cmd/url"
	"github.com/urfave/cli"
)

func URLSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "read",
			Flags:  url.ReadFlags(),
			Action: url.Read,
		},
		{
			Name:   "set",
			Flags:  url.SetFlags(),
			Action: url.Set,
		},
	}

	return commands
}
