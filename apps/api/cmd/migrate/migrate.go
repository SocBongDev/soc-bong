package migrate

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Database migration command",
		Long:  "Migrate the database up or down",
	}
	cmd.AddCommand(UpCmd())
	cmd.AddCommand(DownCmd())
	cmd.AddCommand(ForceCmd())

	return cmd
}
