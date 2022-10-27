package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	closeCmd "github.com/shanmugharajk/vault/pkg/cmd/close"
	deleteCmd "github.com/shanmugharajk/vault/pkg/cmd/delete"
	fetchCmd "github.com/shanmugharajk/vault/pkg/cmd/fetch"
	fetchAllCmd "github.com/shanmugharajk/vault/pkg/cmd/fetchall"
	openCmd "github.com/shanmugharajk/vault/pkg/cmd/open"
	saveCmd "github.com/shanmugharajk/vault/pkg/cmd/save"
	setupCmd "github.com/shanmugharajk/vault/pkg/cmd/setup"
)

func NewCmdRoot() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "vault",
		Short: "vault is command line tool to save passwords",
		Long: heredoc.Doc(`
		About:
		  vault is command line tool to save passwords / secret strings with a single master password 'vault help' in
		the terminal to learn more.
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(setupCmd.NewSetupCmd())
	rootCmd.AddCommand(fetchCmd.NewFetchCmd())
	rootCmd.AddCommand(saveCmd.NewSaveCmd())
	rootCmd.AddCommand(openCmd.NewOpenCmd())
	rootCmd.AddCommand(closeCmd.NewCloseCmd())
	rootCmd.AddCommand(fetchAllCmd.NewFetchAllCmd())
	rootCmd.AddCommand(deleteCmd.NewDeleteCmd())

	return rootCmd
}
