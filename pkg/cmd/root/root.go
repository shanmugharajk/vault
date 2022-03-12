package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	fetchCmd "github.com/shanmugharajk/vault/pkg/cmd/fetch"
	setupCmd "github.com/shanmugharajk/vault/pkg/cmd/setup"
)

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "vault",
		Aliases: []string{"va"},
		Short:   "vault is command line tool to save passwords",
		Long: heredoc.Doc(`
		vault is command line tool to save passwords / secret strings with a single master password

			Run
				$ vault help

			to learn more.
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(setupCmd.NewSetupCmd())
	rootCmd.AddCommand(fetchCmd.NewFetchCmd())

	return rootCmd
}
