package migrate

import (
	"log"
	"strconv"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
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
			log.Println("Migrating down", steps)
			cfg, err := config.New()
			if err != nil {
				log.Fatalln("config.New err:", err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				log.Fatalln("database.NewMigrator err: ", err)
			}

			if err := m.Down(steps); err != nil {
				log.Fatal("Down failed: ", err)
			}

			log.Println("Migrate down success!")
		},
	}

	return cmd
}
