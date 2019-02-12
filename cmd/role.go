package cmd

import (
	"github.com/takaishi/vault-pki/cmd/role"
	"github.com/urfave/cli"
)

func RoleSubcommands() []cli.Command {
	commands := []cli.Command{
		{
			Name:        "list",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#list-roles",
			Flags:       role.ListRoleFlags(),
			Action:      role.ListRole,
		},
		{
			Name:        "create",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#create-update-role",
			Flags:       role.CreateRoleFlags(),
			Action:      role.CreateRole,
		},
		{
			Name:        "update",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#create-update-role",
			Flags:       role.CreateRoleFlags(),
			Action:      role.CreateRole,
		},
		{
			Name:        "create",
			Description: "see more info: https://www.vaultproject.io/api/secret/pki/index.html#delete-role",
			Flags:       role.DeleteRoleFlags(),
			Action:      role.DeleteRole,
		},
	}

	return commands
}
