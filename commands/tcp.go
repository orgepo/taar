package commands

import (
	"net"

	"github.com/spf13/cobra"
)

func BuildTCPCommand(parentCmd *cobra.Command) {
	TCPCmd := &cobra.Command{
		Use:   "tcp",
		Short: "TCP utils",
	}
	buildListenTCPCommand(TCPCmd)
	buildSendPacketTCPCommand(TCPCmd)

	parentCmd.AddCommand(TCPCmd)
}

func buildListenTCPCommand(parentCmd *cobra.Command) {
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "listens to TCP packets",
	}
	parentCmd.AddCommand(listenCmd)

	listenCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErr("please provide an address to listen.")
		}

		l, err := net.Listen("tcp", args[0])
		if err != nil {
			cmd.PrintErrf("can't listen tcp, error:\n%v", err)
		} else {
			cmd.Printf("start tcp server on: %s\n", args[0])

			pch := make(chan string, 1024)
			go accept(l, pch)

			for msg := range pch {
				cmd.Printf("new packet\n data: %s\n", msg)
			}
		}
	}
}

func buildSendPacketTCPCommand(parentCmd *cobra.Command) {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "sends TCP packet to a server",
	}
	parentCmd.AddCommand(sendCmd)

	toOpt := sendCmd.Flags().String("to", "", "IP address of server")
	portOpt := sendCmd.Flags().IntP("port", "p", 0, "IP address of server")

	sendCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErrln("please provide data to send")
		}

		addr := &net.TCPAddr{
			IP:   net.ParseIP(*toOpt),
			Port: *portOpt,
		}

		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			cmd.PrintErrf("can not Dial TCP server\n error:%v\n", err)
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

// ! not CMDs.
func accept(l net.Listener, pch chan string) {
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go read(conn, pch)
	}
}

func read(conn net.Conn, pch chan string) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}
		pch <- string(buf[:n])
	}
}
