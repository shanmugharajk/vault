package fetch

import (
	"errors"
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/utils"
	"github.com/spf13/cobra"
)

func NewFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fetch",
		Short:   "Fetch the secret password",
		Long:    "Fetch the secret password with the key",
		Aliases: []string{"f"},
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase, saltkey := utils.ReadSecrets()
			key := utils.ReadPassword("\nenter the key to fetch\n", 0)

			keyToFetch := crypt.CreateHashKey(key, saltkey)
			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)

			var secret models.Secret
			database.Db.First(&secret, "key = ?", keyToFetch)

			if len(secret.Value) == 0 {
				return errors.New("sorry, unable to find the matching key")
			}

			value := crypt.Decrypt([]byte(secret.Value), saltedPassphrase)
			fmt.Println(string(value))

			return nil
		},
	}

	return cmd
}
