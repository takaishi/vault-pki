package cmd

import (
	"github.com/takaishi/vault-pki/cmd/certificate"
	"github.com/urfave/cli"
)

func CertificateSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:        "list",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#list-certificates , https://www.vaultproject.io/api/secret/pki/index.html#read-certificate",
			Flags:       certificate.ListCertificateFlags(),
			Action:      certificate.ListCertificate,
		},
		{
			Name:        "read",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#read-certificate",
			Flags:       certificate.ReadCertificateFlag(),
			Action:      certificate.ReadCertificate,
		},
		{
			Name:        "generate",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#generate-certificate",
			Flags:       certificate.GenerateCertificateFlags(),
			Action:      certificate.GenerateCertificate,
		},
		{
			Name:        "revoke",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#revoke-certificate",
			Flags:       certificate.RevokeCertificateFlags(),
			Action:      certificate.RevokeCertificate,
		},
	}

	return commands
}
