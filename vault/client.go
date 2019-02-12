package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func NewClient() (*api.Client, error) {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := ""
	vaultTokenFromEnv := os.Getenv("VAULT_TOKEN")

	if vaultTokenFromEnv != "" {
		vaultToken = vaultTokenFromEnv
	} else {
		vaultTokenFromCache, err := ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", os.Getenv("HOME")))
		if err != nil {
			return nil, err
		}
		vaultToken = string(vaultTokenFromCache)
	}

	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		return nil, err
	}
	client.SetToken(vaultToken)

	return client, nil
}
