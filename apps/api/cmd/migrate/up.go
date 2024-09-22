package migrate

import (
	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/spf13/cobra"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func UpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Migrate the database up",
		Long:  "Running all migrations to bring the database up to date",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.New()
			if err != nil {
				logger.Error("config.New err", "err", err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				logger.Error("database.NewMigrator err", "err", err)
			}

			if err := m.Up(); err != nil {
				logger.Error("Up failed", "err", err)
			}

			logger.Info("Migrate up success!")
		},
	}

	return cmd
}
