package commands

import (
	"fmt"
	"os"

	"github.com/kehiy/taar/utils"
	cobra "github.com/spf13/cobra"
)

func BuildDNSCmd(parentCmd *cobra.Command) {
	dnsCmd := &cobra.Command{
		Use:   "dns",
		Short: "change and manage dns",
	}
	parentCmd.AddCommand(dnsCmd)

	showOpt := dnsCmd.Flags().BoolP("show", "s", false, "show dns config")
	setOpt := dnsCmd.Flags().Bool("set", false, "set new dns")

	dnsCmd.Run = func(cmd *cobra.Command, args []string) {
		if *showOpt {
			cmd.Println(utils.ShowResolve())
		}
		if *setOpt {
			err := changeDNS(args)
			if err != nil {
				cmd.PrintErrf("can't change dns: error:\n%v\n", err)
			} else {
				cmd.Printf("dns successfully changed, new config:\n%s\n", args)
			}
		}
	}
}

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
