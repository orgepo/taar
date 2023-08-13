package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "taar",
		Version: "1.0.0",
	}

	buildDNSCmd(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
