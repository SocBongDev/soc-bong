package cmd

import (
	"log"

	"github.com/SocBongDev/soc-bong/cmd/migrate"
	"github.com/SocBongDev/soc-bong/cmd/serve"
	"github.com/spf13/cobra"
)

func new() *cobra.Command {
	cmd := &cobra.Command{Use: "server"}
	cmd.AddCommand(serve.New())
	cmd.AddCommand(migrate.New())

	return cmd
}

func Exec() {
	cmd := new()
	if err := cmd.Execute(); err != nil {
		log.Fatalln("Exec err: ", err)
	}
}
