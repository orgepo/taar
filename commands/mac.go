package commands

import (
	"net"

	"github.com/spf13/cobra"
)

func BuildMACCommand(parentCmd *cobra.Command) {
	MACCmd := &cobra.Command{
		Use:   "mac",
		Short: "MAC address stuff",
	}
	buildShowMACCommand(MACCmd)

	parentCmd.AddCommand(MACCmd)
}

func buildShowMACCommand(parentCmd *cobra.Command) {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show your mac address list",
	}
	parentCmd.AddCommand(showCmd)

	showCmd.Run = func(cmd *cobra.Command, _ []string) {
		interfaces, err := net.Interfaces()
		if err != nil {
			cmd.PrintErrf("can't get device MAC address: %v", err)
		}

		for i, inter := range interfaces {
			cmd.Printf("%d-%v\n", i, inter.HardwareAddr)
		}
	}
}
