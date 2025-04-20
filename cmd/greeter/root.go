// cmd/greeter/root.go

package greeter

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greeter",
		Short: "A simple cli test",
		Long:  "A basic run through using Cobra to create a custom CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), "Welcome to Connor Cli!\n")
		},
	}
	cmd.AddCommand(GreetCommand())

	return cmd
}
