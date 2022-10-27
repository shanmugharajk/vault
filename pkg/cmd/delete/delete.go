package delete

import (
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete the given key",
		Long:    "Find the matching key and deletes it.",
		Aliases: []string{"d"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var passphrase string
			var saltkey string

			passphrase, saltkey = secret.GetSecrets()
			if len(passphrase) == 0 || len(saltkey) == 0 {
				passphrase, saltkey = secret.ReadSecrets()
			}

			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)

			itemToDelete := secret.ReadPassword("\nenter the key to delete\n", 0)

			var secrets []models.Secret
			database.Db.Find(&secrets)

			for _, v := range secrets {
				key := string(crypt.Decrypt([]byte(v.Key), saltedPassphrase))
				if key != itemToDelete {
					continue
				}

				res := database.Db.Delete(&models.Secret{Key: v.Key})

				if res.RowsAffected > 0 {
					fmt.Printf("Successfully deleted the '%s'\n", string(itemToDelete))
				}

				return res.Error
			}

			fmt.Printf("No matching records found for '%s'\n", string(itemToDelete))

			return nil
		},
	}

	return cmd
}
