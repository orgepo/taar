package commands

import (
	"net"

	"github.com/spf13/cobra"
)

func BuildTCPCommand(parentCmd *cobra.Command) {
	TCPCmd := &cobra.Command{
		Use:   "tcp",
		Short: "listen to tcp",
	}
	buildListenTCPCommand(TCPCmd)

	parentCmd.AddCommand(TCPCmd)
}

func buildListenTCPCommand(parentCmd *cobra.Command) {
	listenCmd := &cobra.Command{
		Use:   "listen",
		Short: "listen to TCP packets",
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
