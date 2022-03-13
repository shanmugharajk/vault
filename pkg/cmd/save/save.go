package save

import (
	"errors"
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/utils"
	"github.com/spf13/cobra"
)

func NewSaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "save",
		Short:   "Save the secret",
		Long:    "Saves the secret password with the key, value encrypted with the passphrase and saltkey given",
		Aliases: []string{"save"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("pass the right no of arguments")
			}

			key := args[0]
			value := args[1]

			if len(key) < 1 || len(value) < 1 {
				return errors.New("pass the key, value with minimum length 1")
			}

			passphrase := utils.PromptSecret("please enter passphrase\n")
			saltkey := utils.PromptSecret("please enter key\n")

			keyToSave := crypt.CreateHashKey(key, saltkey)
			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)
			valueToSave := crypt.Encrypt([]byte(value), saltedPassphrase)

			result := database.Db.Create(&models.Secret{Key: keyToSave, Value: string(valueToSave)})
			if result.RowsAffected <= 0 || result.Error != nil {
				return errors.New("sorry, unable to save the data")
			}

			fmt.Println("Saved successfully!")

			return nil
		},
	}

	return cmd
}
