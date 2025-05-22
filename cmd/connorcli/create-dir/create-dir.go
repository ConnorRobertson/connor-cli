package createdir

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func CreateDirCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "credir",
		Short: "Use this similar to mkdir but lets you make subdirectories as well",
		Long:  "Create a directory from the current directory at a given location unless it already exists",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			if file, err := os.Stat(name); err == nil && file.IsDir() {
				fmt.Fprintf(cmd.OutOrStdout(), "This directory %v already exists.\n", name)
			}

			err := os.MkdirAll(name, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Create a new directory %v!\n", name)
		},
	}

}
