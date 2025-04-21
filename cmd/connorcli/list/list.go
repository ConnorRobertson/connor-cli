package list

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List falls in the current directory",
		Long:  "List all non-hidden files in the current directory",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := os.ReadDir("./")
			if err != nil {
				log.Fatal(err)
			}
			// OS print Colours
			var Reset = "\033[0m"
			var Cyan = "\033[36m"
			var Red = "\033[31m"
			var Green = "\033[32m"

			for _, file := range files {
				// Ignore hidden files
				var isHidden bool = file.Name()[0] == '.'
				if !isHidden && !file.IsDir() {
					fileInfo, err := os.Stat(file.Name())
					if err != nil {
						log.Fatal(err)
					}
					output := "- " + Red + fileInfo.Name() + Green + " Size: " + strconv.FormatInt(fileInfo.Size(), 10) + " Bytes " +
						Cyan + "Last Modified: " + fileInfo.ModTime().UTC().Format(time.UnixDate) + Reset
					fmt.Println(output)
				}
				if file.IsDir() && !isHidden {

					output := "📁 " + Cyan + file.Name() + Reset
					fmt.Println(output)
				}
			}
		},
	}
	// cmd.AddCommand(ListAllCommand())

	return cmd
}
