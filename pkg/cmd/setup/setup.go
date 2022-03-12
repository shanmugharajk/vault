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
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			database.Connect(&database.Config{
				Automigrate: true,
				Recreate:    true,
			})

			return nil
		},
	}

	return cmd
}
