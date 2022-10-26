package save

import (
	"errors"
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/spf13/cobra"
)

func NewSaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "save",
		Short:   "Save the secret",
		Long:    "Saves the secret password with the key, value encrypted with the passphrase and saltkey given",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var passphrase string
			var saltkey string

			passphrase, saltkey = secret.GetSecrets()
			if len(passphrase) == 0 || len(saltkey) == 0 {
				passphrase, saltkey = secret.ReadSecrets()
			}

			key := secret.ReadPassword("\nenter the key to save with minimum length 2\n", 2)
			value := secret.ReadPassword("\nenter the value to save\n", 0)

			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)
			keyToSave := crypt.Encrypt([]byte(key), saltedPassphrase)
			valueToSave := crypt.Encrypt([]byte(value), saltedPassphrase)

			result := database.Db.Create(&models.Secret{Key: string(keyToSave), Value: string(valueToSave)})
			if result.RowsAffected <= 0 || result.Error != nil {
				return errors.New("sorry, unable to save the data")
			}

			fmt.Println("Saved successfully!")

			return nil
		},
	}

	return cmd
}
