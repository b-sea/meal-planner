// Package cli defines the command line interface for the meal planner.
package cli

import (
	"github.com/spf13/cobra"
)

// New creates a new CLI root command.
func New(version string) *cobra.Command {
	root := &cobra.Command{
		Version: version,
		Use:     "mealplan [OPTIONS] [COMMAND]",
		Short:   "The meal planner backend tools",
	}

	root.AddCommand(startCmd(version))

	return root
}
