package size

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func SizeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sz",
		Short: "Finds the size of the current directory in Bytes",
		Long:  "Finds the size of the files and current directory of all subdirectories measured in Bytes, can specify a directory with arguments",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				if strings.HasPrefix(args[0], "./") && strings.HasSuffix(args[0], "/") {
					fmt.Fprint(cmd.OutOrStdout(), getSize(args[0], 0), " Bytes\n")
				} else {
					dirName := args[0]
					if !strings.HasPrefix(args[0], "./") {
						dirName = "./" + dirName
					}
					if !strings.HasSuffix(args[0], "/") {
						dirName += "/"
					}
					fmt.Fprint(cmd.OutOrStdout(), getSize(dirName, 0), " Bytes\n")
				}
			} else {
				fmt.Fprint(cmd.OutOrStdout(), getSize("./", 0), " Bytes\n")
			}
		},
	}

}

func getSize(dirName string, currSize int64) int64 {
	// Recursively enter directories add file size to running total
	files, err := os.ReadDir(dirName)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var isHidden bool = strings.HasPrefix(file.Name(), ".")
		if !isHidden && !file.IsDir() {
			// Add file size
			fileInfo, err := os.Stat(dirName + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			currSize += fileInfo.Size()
		} else if !isHidden && file.IsDir() {
			currSize += getSize(dirName+file.Name()+"/", currSize)
		}
	}
	return currSize
}
