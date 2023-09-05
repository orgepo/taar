package commands

import (
	"fmt"
	"net"
	"os"

	"github.com/kehiy/taar/utils"
	cobra "github.com/spf13/cobra"
)

func BuildDNSCommand(parentCmd *cobra.Command) {
	dnsCmd := &cobra.Command{
		Use:   "dns",
		Short: "change and manage DNS",
	}
	buildSetCommand(dnsCmd)
	buildShowResolveCommand(dnsCmd)
	buildAskCommand(dnsCmd)

	parentCmd.AddCommand(dnsCmd)
}

func buildShowResolveCommand(parentCmd *cobra.Command) {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "shows DNS setting",
	}
	parentCmd.AddCommand(showCmd)

	showCmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Println(utils.ShowResolve())
	}
}

func buildSetCommand(parentCmd *cobra.Command) {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "sets new DNS",
	}
	parentCmd.AddCommand(setCmd)

	setCmd.Run = func(cmd *cobra.Command, args []string) {
		err := changeDNS(args)
		if err != nil {
			cmd.PrintErrf("can't change DNS: error:\n%v\n", err)
		} else {
			cmd.Printf("DNS successfully changed, new config:\n%s\n", args)
		}
	}
}

func buildAskCommand(parentCmd *cobra.Command) {
	askCmd := &cobra.Command{
		Use:   "ask",
		Short: "makes a DNS query and show the result",
	}
	parentCmd.AddCommand(askCmd)

	askCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErr("please provide a domain")
		}

		ips, err := net.LookupIP(args[0])
		if err != nil {
			cmd.PrintErrf("can't lookup address:%v", err)
		}

		for i, ip := range ips {
			cmd.Printf("%d-%v\n", i, ip)
		}
	}
}

// ! not CMDs.
func changeDNS(DNSs []string) error {
	path := "/etc/resolv.conf"

	newContent := `
	# DO NOT CHANGE!
	# managed by taar network manager.
	`
	for _, dns := range DNSs {
		newContent += fmt.Sprintf("nameserver %s\n", dns)
	}

	err := os.WriteFile(path, []byte(newContent), 0o600)
	if err != nil {
		return err
	}

	return nil
}
