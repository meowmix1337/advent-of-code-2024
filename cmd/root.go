package cmd

import (
	"github.com/dvan-sqsp/advent-of-code-2024/internal/runner"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var day int
	cmd := &cobra.Command{
		Use:   "advent",
		Short: "Advent of Code 2024 solution runner",
		Run: func(cmd *cobra.Command, args []string) {
			runner.Run(day)
		},
	}

	cmd.Flags().IntVar(&day, "day", 0, "Day to run")

	return cmd
}
