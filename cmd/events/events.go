package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	EventsCmd.AddCommand(EventListCmd)
	EventsCmd.AddCommand(EventTodayCmd)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Check events in calendar",
	Long:  `Choose a calendar ID and check events`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Choose a command: \nlist or create")
	},
}
