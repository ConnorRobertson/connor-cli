package create

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	return &cobra.Command{
		Use: "cre",
		Short: "Create a new file",
		Long: "Create a new file similar to touch",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string){
			file, err := os.Create(args[0])
			
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close() // Closes file when function closes

			fmt.Fprintf(cmd.OutOrStdout(), "%v is successfully created!\n", file.Name())
		},
	}
}