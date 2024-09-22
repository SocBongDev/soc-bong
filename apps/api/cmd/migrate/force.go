package migrate

import (
	"strconv"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/spf13/cobra"
)

var version int

func ForceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "force",
		Short: "Force migration to the version number",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			var err error
			if version, err = strconv.Atoi(args[0]); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Migrating force", version)
			cfg, err := config.New()
			if err != nil {
				logger.Error("config.New err", "err", err)
				panic(err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				logger.Error("database.NewMigrator err", "err", err)
			}

			if err := m.Force(version); err != nil {
				logger.Error("Force failed", "err", err)
			}

			logger.Info("Migrate force success!")
		},
	}

	return cmd
}
