// cmd/greeter/root.go

package connorcli

import (
	"fmt"
	// Custom cmds
	"connorcli/cmd/connorcli/connorcli"
	"connorcli/cmd/connorcli/create"
	del "connorcli/cmd/connorcli/delete"
	"connorcli/cmd/connorcli/list"
	"connorcli/cmd/connorcli/rename"

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

	// Add Commands
	cmd.AddCommand(connorcli.GreetCommand())
	cmd.AddCommand(list.ListCommand())
	cmd.AddCommand(list.ListAllCommand())
	cmd.AddCommand(rename.RenameCommand())
	cmd.AddCommand(del.DeleteCommand()) // Delete is a keyword I shouldn't use it
	cmd.AddCommand(create.CreateCommand())

	return cmd
}
