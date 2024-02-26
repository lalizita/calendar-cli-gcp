package cmd

import (
	"fmt"
	"os"

	agenda "github.com/lalizita/calendar-cli-gcp/cmd/agenda"
	events "github.com/lalizita/calendar-cli-gcp/cmd/events"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "calendar",
		Short:         "Your calendar CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(agenda.AgendaCmd)
	rootCmd.AddCommand(events.EventsCmd)

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
