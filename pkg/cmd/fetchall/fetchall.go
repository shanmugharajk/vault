package fetchall

import (
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/shanmugharajk/vault/internal/utils"
	"github.com/spf13/cobra"
)

func NewFetchAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fetchall",
		Short:   "Fetches all the saved secrets",
		Long:    "Fetches all the saved secrets.",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var passphrase string
			var saltkey string

			passphrase, saltkey = secret.GetSecrets()
			if len(passphrase) == 0 || len(saltkey) == 0 {
				passphrase, saltkey = utils.ReadSecrets()
			}

			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)

			var secrets []models.Secret
			database.Db.Find(&secrets)

			for _, v := range secrets {
				value := crypt.Decrypt([]byte(v.Value), saltedPassphrase)
				fmt.Println(string(value))
			}

			return nil
		},
	}

	return cmd
}
