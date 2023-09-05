package commands

import (
	"net"
	"strings"

	"github.com/spf13/cobra"
)

func BuildUDPCommand(parentCmd *cobra.Command) {
	UDPCmd := &cobra.Command{
		Use:   "udp",
		Short: "UDP utils",
	}
	buildListenUDPCommand(UDPCmd)
	buildSendPacketUDPCommand(UDPCmd)

	parentCmd.AddCommand(UDPCmd)
}

func buildListenUDPCommand(parentCmd *cobra.Command) {
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "listens to UDP packets",
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

func buildSendPacketUDPCommand(parentCmd *cobra.Command) {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "sends UDP packet to a server",
	}
	parentCmd.AddCommand(sendCmd)

	toOpt := sendCmd.Flags().String("to", "", "IP address of server")
	portOpt := sendCmd.Flags().IntP("port", "p", 0, "IP address of server")

	sendCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErrln("please provide data to send")
		}
		addr := &net.UDPAddr{
			IP:   net.ParseIP(*toOpt),
			Port: *portOpt,
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			cmd.PrintErrf("can not Dial UDP server\n error:%v\n", err)
		}
		defer conn.Close()

		n, err := conn.Write([]byte(args[0]))
		if err != nil {
			cmd.PrintErrf("write to connection failed\n error:%v\n", err)
		}

		cmd.Printf("%d bytes send\n waiting for respond\n", n)

		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			cmd.PrintErrf("read from connection failed\n error:%v\n", err)
		}

		cmd.Printf("%v\n", string(buf))
	}
}
