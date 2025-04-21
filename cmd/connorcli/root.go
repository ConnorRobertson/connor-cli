// cmd/greeter/root.go

package connorcli

import (
	"fmt"
	// Custom cmds
	"connorcli/cmd/connorcli/list"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple cli test",
		Long:  "A basic run through using Cobra to create a custom CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), "Welcome to Connor CLI!\n")
		},
	}
	cmd.AddCommand(GreetCommand())
	cmd.AddCommand(list.ListCommand())

	return cmd
}
