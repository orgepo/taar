package commands

import (
	"encoding/json"
	"net"

	cobra "github.com/spf13/cobra"

	"github.com/kehiy/taar/types"
	"github.com/kehiy/taar/utils"
)

func BuildIPCommand(parentCmd *cobra.Command) {
	IPCmd := &cobra.Command{
		Use:   "ip",
		Short: "IP utils",
	}
	buildTrackCommand(IPCmd)
	buildIPShowCommand(IPCmd)

	parentCmd.AddCommand(IPCmd)
}

func buildTrackCommand(parentCmd *cobra.Command) {
	trackCmd := &cobra.Command{
		Use:   "track",
		Short: "tracks an IP address data",
	}
	parentCmd.AddCommand(trackCmd)

	trackCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, IP := range args {
				showData(IP, cmd)
			}
		} else {
			cmd.Println("please provide an IP address")
		}
	}
}

func buildIPShowCommand(parentCmd *cobra.Command) {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "shows device IP address",
	}
	parentCmd.AddCommand(showCmd)

	showCmd.Run = func(cmd *cobra.Command, args []string) {
		interfaces, err := net.InterfaceAddrs()
		if err != nil {
			cmd.PrintErrf("can't get device IP address: %v", err)
		}

		for i, inter := range interfaces {
			cmd.Printf("%d-%v\n", i, inter.String())
		}
	}
}

// ! not CMDs.
func showData(ip string, cmd *cobra.Command) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := utils.GetDataHTTP(url)

	data := types.IP{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		cmd.Println("Unable to unmarshal the response")
	}

	cmd.Println("DATA FOUND :")

	cmd.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE:%s\nPOSTAL :%s\n",
		data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}
