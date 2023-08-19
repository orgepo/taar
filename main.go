package main

import (
	"github.com/kehiy/taar/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "taar",
		Version: "0.7.0",
	}

	cmd.BuildDNSCmd(rootCmd)
	cmd.BuildTCPCmd(rootCmd)
	cmd.BuildIPCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
