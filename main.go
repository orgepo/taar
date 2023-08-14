package main

import (
	"taar/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "taar",
		Version: "1.0.0",
	}

	cmd.BuildDNSCmd(rootCmd)
	cmd.BuildTCPCmd(rootCmd)
	
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
