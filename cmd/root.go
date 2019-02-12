package cmd

import (
	"github.com/takaishi/vault-pki/cmd/root"
	"github.com/urfave/cli"
)

func RootSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "generate",
			Flags:  root.GenerateFlags(),
			Action: root.Generate,
		},
		{
			Name:   "delete",
			Flags:  root.DeleteFlags(),
			Action: root.Delete,
		},
	}

	return commands
}
