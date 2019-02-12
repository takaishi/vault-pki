package cmd

import (
	"github.com/takaishi/vault-pki/cmd/certificate"
	"github.com/urfave/cli"
)

func CertificateSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:   "list",
			Flags:  certificate.ListCertificateFlags(),
			Action: certificate.ListCertificate,
		},
		{
			Name:   "read",
			Flags:  certificate.ReadCertificateFlag(),
			Action: certificate.ReadCertificate,
		},
		{
			Name:        "generate",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#generate-certificate",
			ArgsUsage:   "[pki] [role]",
			Flags:       certificate.GenerateCertificateFlags(),
			Action:      certificate.GenerateCertificate,
		},
		{
			Name:   "revoke",
			Flags:  certificate.RevokeCertificateFlags(),
			Action: certificate.RevokeCertificate,
		},
	}

	return commands
}
