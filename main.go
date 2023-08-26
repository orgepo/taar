package main

import (
	commands "github.com/kehiy/taar/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "taar",
		Version: "0.7.0",
	}

	commands.BuildDNSCommand(rootCmd)
	commands.BuildTCPCommand(rootCmd)
	commands.BuildIPCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
