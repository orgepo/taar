package commands

import (
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func BuildUDPCommand(parentCmd *cobra.Command) {
	UDPCmd := &cobra.Command{
		Use:   "udp",
		Short: "UDP protocol utils",
	}
	buildListenUDPCommand(UDPCmd)

	parentCmd.AddCommand(UDPCmd)
}

func buildListenUDPCommand(parentCmd *cobra.Command) {
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "listen to UDP packets",
	}
	parentCmd.AddCommand(listenCmd)

	portOpt := listenCmd.Flags().IntP("port", "p", 3000, "listening port")

	listenCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErr("please provide an address to listen.")
		}
		address := &net.UDPAddr{
			IP:   net.ParseIP(args[0]),
			Port: *portOpt,
		}
		conn, err := net.ListenUDP("udp", address)
		if err != nil {
			cmd.PrintErrf("can't listen UDP, error:\n%v", err)
			os.Exit(1)
		}
		defer conn.Close()
		cmd.Printf("start UDP server on: %s\n", args[0])
		for {
			message := make([]byte, 20)
			n, remote, err := conn.ReadFromUDP(message[:])
			if err != nil {
				cmd.PrintErrf("packet loss: \nerror: %v\n remote: %v\n", err, remote)
			}
			data := strings.TrimSpace(string(message[:n]))
			cmd.Printf("new packet\n data: %s\n remote:%v\n", data, remote.IP)
		}
	}
}
