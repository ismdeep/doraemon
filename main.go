package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ismdeep/doraemon/command"
)

func main() {
	cmd := cobra.Command{
		Use:          "doraemon",
		Short:        "doraemon",
		SilenceUsage: true,
	}

	cmd.AddCommand(command.DB())
	cmd.AddCommand(command.Version())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
