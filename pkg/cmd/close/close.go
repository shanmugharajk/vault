package open

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/shanmugharajk/vault/internal/secret"
	"github.com/spf13/cobra"
)

func NewCloseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close",
		Short: "Closes the vault",
		Long: heredoc.Doc(`
			This saves the passphrase into the shell session, so until we close or refresh the shell 
			session we don't need to enter the passphrase for subsequent operation.
		`),
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return secret.DelSecrets()
		},
	}

	return cmd
}
