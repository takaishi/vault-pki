package cmd

import (
	"github.com/takaishi/vault-pki/cmd/intermediate"
	"github.com/urfave/cli"
)

func IntermediateSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:        "set-signed",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#set-signed-intermediate",
			Flags:       intermediate.SetSignedIntermediateFlags(),
			Action:      intermediate.SetSignedIntermediate,
		},
		{
			Name:        "generate",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#generate-intermediate",
			Flags:       intermediate.GenerateIntermediateFlags(),
			Action:      intermediate.GenerateIntermediate,
		},
		{
			Name:        "sign",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#sign-intermediate",
			Flags:       intermediate.SignIntermediateFlags(),
			Action:      intermediate.SignIntermediate,
		},
	}

	return commands
}
