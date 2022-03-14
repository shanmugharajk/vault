package setup

import (
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/spf13/cobra"
)

func NewSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setup",
		Short:   "Initial setup",
		Long:    "This command performs the initial setup required for the first time run",
		Aliases: []string{"setup"},
		RunE: func(cmd *cobra.Command, args []string) error {
			database.Recreate()
			return nil
		},
	}

	return cmd
}
