package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ismdeep/doraemon/command"
)

func main() {
	cmd := cobra.Command{
		Use:   "doraemon",
		Short: "doraemon",
	}

	cmd.AddCommand(command.DB())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
