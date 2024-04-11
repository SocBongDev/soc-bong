package serve

import (
	"log"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Args:  cobra.ArbitraryArgs,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Serving dinner...")

			config, err := config.New()
			if err != nil {
				log.Panicln("config.New err: ", err)
			}

			app := NewApp(config)
			app.RunHttpServer()
		},
	}

	return cmd
}
