// Package provides an entrypoint for a range of supported CLI
// Commands
package main

import (
	"fmt"
	"os"

	// Custom CLI commands
	connorcli "connorcli/cmd/connorcli/root"
)

func main() {
	rootCmd := connorcli.RootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
