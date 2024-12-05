package main

import (
	"os"

	"github.com/dvan-sqsp/advent-of-code-2024/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("Failed to execute root command")
		os.Exit(1)
	}
}
