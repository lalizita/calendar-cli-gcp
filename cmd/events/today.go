package cmd

import (
	"fmt"

	"github.com/lalizita/calendar-cli-gcp/internal/calendar"
	"github.com/spf13/cobra"
)

var EventTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "list all events for today",
	Long:  "List all event you have",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := calendar.NewClient()

		id, err := c.GetAgendaID()
		if err != nil {
			return err
		}

		events, err := c.ListTodayEvents(id)
		if err != nil {
			return err
		}
		fmt.Printf("%s", events)

		return nil
	},
}
