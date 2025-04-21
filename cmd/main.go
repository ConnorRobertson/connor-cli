// Package provides an entrypoint for a range of supported CLI
// Commands
package main

import (
	"fmt"
	"os"

	// Custom CLI commands
	"connorcli/cmd/connorcli"
)

func main() {
	rootCmd := connorcli.RootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
