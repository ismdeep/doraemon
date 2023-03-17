package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ismdeep/doraemon/interval/version"
)

// Version command
func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get doraemon version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version:", version.Version)
		},
	}
}
