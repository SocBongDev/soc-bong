package migrate

import (
	"log"
	"strconv"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/database"
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
			log.Println("Migrating force", version)
			cfg, err := config.New()
			if err != nil {
				log.Fatalln("config.New err:", err)
			}

			m, err := database.NewMigrator(&cfg.DatabaseSecret)
			if err != nil {
				log.Fatalln("database.NewMigrator err: ", err)
			}

			if err := m.Force(version); err != nil {
				log.Fatal("Force failed: ", err)
			}

			log.Println("Migrate force success!")
		},
	}

	return cmd
}
