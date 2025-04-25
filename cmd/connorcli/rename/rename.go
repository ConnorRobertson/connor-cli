package rename

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)


func RenameCommand () *cobra.Command {
	return &cobra.Command {
		Use: "rename",
		Short: "Rename specified file and directory",
		Long: "Rename file or directory args[0] to args[1]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string){
			err := os.Rename(args[0], args[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully renamed: %v\n", args[0])
		},
	}

}
