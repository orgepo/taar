package main

import (
	commands "github.com/kehiy/taar/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "taar",
		Version: "0.8.0",
	}

	commands.BuildDNSCommand(rootCmd)
	commands.BuildTCPCommand(rootCmd)
	commands.BuildIPCommand(rootCmd)
	commands.BuildMACCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
