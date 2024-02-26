package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	AgendaCmd.AddCommand(AgendaCreateCmd)
}

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "agenda configuration",
	Long:  `config for agendas in service account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a command: \nlist or create")
	},
}
