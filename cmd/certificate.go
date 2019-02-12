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
			Name:   "generate",
			Flags:  certificate.GenerateCertificateFlags(),
			Action: certificate.GenerateCertificateFlags,
		},
		{
			Name:   "revoke",
			Flags:  certificate.RevokeCertificateFlags(),
			Action: certificate.RevokeCertificateFlags,
		},
	}

	return commands
}
