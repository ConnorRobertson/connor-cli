package connorcli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GreetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "greet",
		Short: "A simple cli test",
		Long:  "A basic run through using Cobra to create a custom CLI",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s!\n", args[0])
		},
	}
}
