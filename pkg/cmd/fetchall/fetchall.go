package fetchall

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/secret"
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
				passphrase, saltkey = secret.ReadSecrets()
			}

			saltedPassphrase := crypt.CreateHashKey(passphrase, saltkey)

			var secrets []models.Secret
			database.Db.Find(&secrets)

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Key", "Value"})
			table.SetAutoWrapText(false)
			table.SetRowLine(true)

			for _, v := range secrets {
				key := crypt.Decrypt([]byte(v.Key), saltedPassphrase)
				value := crypt.Decrypt([]byte(v.Value), saltedPassphrase)
				table.Append([]string{string(key), string(value)})
			}

			table.Render()

			return nil
		},
	}

	return cmd
}
