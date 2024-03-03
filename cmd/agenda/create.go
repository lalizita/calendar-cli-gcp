package cmd

import (
	"github.com/lalizita/calendar-cli-gcp/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "add a agenda id",
	Long:  `calendar agenda add <id>`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		agendaId := args[0]

		c := calendar.NewClient()
		err := c.AddAgenda(agendaId)
		if err != nil {
			return err
		}

		return nil
	},
}
