package cmd

import (
	"github.com/takaishi/vault-pki/cmd/root"
	"github.com/urfave/cli"
)

func RootSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "create",
			Flags:  root.GenerateFlags(),
			Action: root.Generate,
		},
	}

	return commands
}
