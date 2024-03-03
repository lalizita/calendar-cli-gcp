/* cmd: root cmd for events week. */
package cmd

import (
	"fmt"

	"github.com/lalizita/calendar-cli-gcp/internal/calendar"
	"github.com/spf13/cobra"
)

// EventListCmd represents all events in a week.
var EventListCmd = &cobra.Command{
	Use:   "week",
	Short: "list the current events for this week",
	Long:  `List all the events in the week in the selected calendar`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := calendar.NewClient()

		id, err := c.GetAgendaID()
		if err != nil {
			return err
		}

		events, err := c.ListWeekEvents(id)
		if err != nil {
			return err
		}
		fmt.Printf("%s", events)

		return nil
	},
}
