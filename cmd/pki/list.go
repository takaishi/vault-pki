package pki

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/takaishi/vault-pki/vault"
	"github.com/urfave/cli"
)

func List(c *cli.Context) error {
	client, err := vault.NewClient()
	if err != nil {
		return err
	}

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		return err
	}

	pkiList := filterPkiOnly(mounts)

	for name, pki := range pkiList {
		fmt.Printf("%s %s %s\n", name, pki.Accessor, pki.Description)
	}

	return nil

}

func filterPkiOnly(mounts map[string]*api.MountOutput) map[string]*api.MountOutput {
	filtered := map[string]*api.MountOutput{}

	for name, mount := range mounts {
		if mount.Type == "pki" {
			filtered[name] = mount
		}
	}

	return filtered
}
