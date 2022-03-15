package open

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/shanmugharajk/vault/internal/utils"
	"github.com/spf13/cobra"
)

func NewOpenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Opens the vault",
		Long: heredoc.Doc(`
			This saves the passphrase into the shell session, so until we close or refresh the shell 
			session we don't need to enter the passphrase for subsequent operation.
		`),
		Aliases: []string{"o"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return secret.SetSecrets(utils.ReadSecrets())
		},
	}

	return cmd
}
