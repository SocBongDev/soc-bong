package migrate

import (
	"strconv"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/spf13/cobra"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var steps int

func DownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "Migrate the database down - given the number of steps",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			var err error
			if steps, err = strconv.Atoi(args[0]); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Migrating down", steps)
			cfg, err := config.New()
			if err != nil {
				logger.Error("config.New err", "err", err)
				panic(err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				logger.Error("database.NewMigrator err", "err", err)
				panic(err)
			}

			if err := m.Down(steps); err != nil {
				logger.Error("Down failed", "err", err)
				panic(err)
			}

			logger.Info("Migrate down success!")
		},
	}

	return cmd
}
