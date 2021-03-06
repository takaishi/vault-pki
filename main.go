package main

import (
	"github.com/takaishi/vault-pki/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:        "ca",
			Usage:       "ca",
			Subcommands: cmd.CASubcommands(),
		},
		{
			Name:        "pki",
			Usage:       "pki",
			Subcommands: cmd.PkiSubcommands(),
		},
		{
			Name:        "root",
			Usage:       "operate root CA",
			Subcommands: cmd.RootSubcommands(),
		},
		{
			Name:        "url",
			Usage:       "operate URLs",
			Subcommands: cmd.URLSubcommands(),
		},
		{
			Name:        "role",
			Usage:       "operate Role",
			Subcommands: cmd.RoleSubcommands(),
		},
		{
			Name:        "intermediate",
			Usage:       "operate intermediate",
			Subcommands: cmd.IntermediateSubcommands(),
		},
		{
			Name:        "certificate",
			Usage:       "operate Certificate",
			Subcommands: cmd.CertificateSubcommands(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
