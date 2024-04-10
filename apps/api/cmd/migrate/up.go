package migrate

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
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
				log.Fatal("config.New err: ", err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				log.Fatal("database.NewMigrator err: ", err)
			}

			if err := m.Up(); err != nil {
				log.Fatal("Up failed: ", err)
			}

			log.Println("Migrate up success!")
		},
	}

	return cmd
}
