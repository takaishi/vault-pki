package cmd

import (
	"github.com/takaishi/vault-pki/cmd/pki"
	"github.com/urfave/cli"
)

func PkiSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:        "list",
			Description: "see more info: https://www.vaultproject.io/api/system/mounts.html#list-mounted-secrets-engines",
			Action:      pki.List,
		},
		{
			Name:        "enable",
			Description: "see more info: https://www.vaultproject.io/api/system/mounts.html#enable-secrets-engine",
			Flags:       pki.EnableFlags(),
			Action:      pki.Enable,
		},
		{
			Name:        "disable",
			Description: "see more info: https://www.vaultproject.io/api/system/mounts.html#disable-secrets-engine",
			Flags:       pki.DisableFlags(),
			Action:      pki.Disable,
		},
	}

	return commands
}
