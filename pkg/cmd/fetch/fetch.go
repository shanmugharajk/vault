package fetch

import (
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/spf13/cobra"
)

func NewFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fetch",
		Short:   "Fetch the secret password",
		Long:    "Fetch the secret password with the key",
		Aliases: []string{"f"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var passphrase string
			var saltkey string

			passphrase, saltkey = secret.GetSecrets()
			if len(passphrase) == 0 || len(saltkey) == 0 {
				passphrase, saltkey = secret.ReadSecrets()
			}

			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)
			keyToFetch := secret.ReadPassword("\nenter the key to fetch\n", 0)

			var secrets []models.Secret
			database.Db.Find(&secrets)

			for _, v := range secrets {
				key := string(crypt.Decrypt([]byte(v.Key), saltedPassphrase))
				if key == keyToFetch {
					fmt.Println(string(crypt.Decrypt([]byte(v.Value), saltedPassphrase)))
					return nil
				}
			}

			return fmt.Errorf("no matching records found for %s", string(keyToFetch))
		},
	}

	return cmd
}
