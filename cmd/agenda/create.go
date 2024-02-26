package cmd

import (
	"fmt"

	"github.com/lalizita/calendar-cli-gcp/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "add a agenda id",
	Long:  `calendar agenda add <id>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := calendar.NewClient()
		err := c.AddAgenda("2h8ldqm21pcgr5r0h4dj2gr42k@group.calendar.google.com")
		if err != nil {
			return err
		}
		fmt.Println("SUCESSO!!!")

		return nil
	},
}
