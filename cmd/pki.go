package cmd

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func PkiSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name: "list",
			Action: func(c *cli.Context) error {
				vaultAddr := os.Getenv("VAULT_ADDR")
				vaultToken := ""
				vaultTokenFromEnv := os.Getenv("VAULT_TOKEN")

				if vaultTokenFromEnv != "" {
					vaultToken = vaultTokenFromEnv
				} else {
					vaultTokenFromCache, err := ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", os.Getenv("HOME")))
					if err != nil {
						return err
					}
					vaultToken = string(vaultTokenFromCache)
				}

				var httpClient = &http.Client{
					Timeout: 10 * time.Second,
				}

				client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
				if err != nil {
					return err
				}
				client.SetToken(vaultToken)

				mounts, err := client.Sys().ListMounts()
				if err != nil {
					return err
				}

				pkiList := filterPkiOnly(mounts)

				for name, pki := range pkiList {
					fmt.Printf("%s %s %s\n", name, pki.Accessor, pki.Description)
				}

				return nil
			},
		},
	}

	return commands
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
