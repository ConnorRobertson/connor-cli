package list

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List files in the current or specified directory",
		Long:  "List all non-hidden files in the current or specified directory",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				if strings.HasPrefix(args[0], "./") && strings.HasSuffix(args[0], "/") {
					print(readDirectory(args[0], false))
				} else {
					dirName := args[0]
					if !strings.HasPrefix(args[0], "./") {
						dirName = "./" + dirName
					}
					if !strings.HasSuffix(args[0], "/") {
						dirName += "/"
					}

					print(readDirectory(dirName, false))
				}

			} else {
				print(readDirectory("./", false))
			}
		},
	}
}

func ListAllCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "lsall",
		Short: "List files in the all directories",
		Long:  "List all non-hidden files in the current directory and its subdirectories",
		Run: func(cmd *cobra.Command, args []string) {
			print(readDirectory("./", true))
		},
	}
}

func readDirectory(dirName string, isRecursive bool) string {
	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	// OS print Colours
	var Reset = "\033[0m"
	var Cyan = "\033[36m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var output = ""

	for _, file := range files {
		// Ignore hidden files
		var isHidden bool = strings.HasPrefix(file.Name(), ".")

		if !isHidden && !file.IsDir() {
			fileInfo, err := os.Stat(dirName + file.Name())

			// Error Handling
			if err != nil {
				log.Fatal(err)
			}
			// Output in the form of dirName/fileName
			output += Red + dirName + fileInfo.Name() + Green + " Size: " + strconv.FormatInt(fileInfo.Size(), 10) + " Bytes " +
				Cyan + "Last Modified: " + fileInfo.ModTime().UTC().Format(time.UnixDate) + Reset + "\n"
		} else if file.IsDir() && !isHidden {
			if isRecursive {
				// Recursively checks directories and outputs all the files in a given directory
				output += readDirectory(dirName+file.Name()+"/", true)
			} else {
				output += "📁 " + Cyan + file.Name() + Reset + "\n"
			}

		}
	}
	return output
}
