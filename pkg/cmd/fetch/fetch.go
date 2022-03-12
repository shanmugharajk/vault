package fetch

import (
	"fmt"

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
			key := args[0]
			fmt.Println("The key to find is ", key)
			utils.PromptSecret("Please enter passphrase\n")

			// TODO: Add the logic to get passphrase, salt key. Then fetch the value from
			// db, decrypt and set it.

			return nil
		},
	}

	return cmd
}
