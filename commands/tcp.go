package commands

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func BuildTCPCmd(parentCmd *cobra.Command) {
	tcpCmd := &cobra.Command{
		Use:   "tcp",
		Short: "listen to tcp",
	}
	parentCmd.AddCommand(tcpCmd)

	listenOpt := tcpCmd.Flags().StringP("listen", "l", ":3000", "listening port")

	tcpCmd.Run = func(cmd *cobra.Command, args []string) {
		l, err := net.Listen("tcp", *listenOpt)
		if err != nil {
			cmd.PrintErrf("can't listen tcp, error:\n%v", err)
		} else {
			cmd.Printf("start tcp server on: %s\n", *listenOpt)
			pch := make(chan string, 1024)
			go accept(l, pch)
			for msg := range pch {
				fmt.Printf("new packet data: %s\n", msg)
			}
		}
	}
}

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
