package delete

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func DeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use: "del",
		Short: "Delete specified file or directory",
		Long: "Given a file or directory remove file or directory similar to rm",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string){
			err := os.Remove(args[0])
			
			if err != nil {
				log.Fatal(err)
			}


			fmt.Fprintf(cmd.OutOrStdout(), "%v is successfully deleted!\n", args[0])
		},
	}
}