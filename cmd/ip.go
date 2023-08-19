package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	cobra "github.com/spf13/cobra"

	"github.com/kehiy/taar/types"
	"github.com/kehiy/taar/utils"
)

func BuildIPCmd(parentCmd *cobra.Command) {
	IPCmd := &cobra.Command{
		Use:   "ip",
		Short: "ip utils",
	}
	parentCmd.AddCommand(IPCmd)

	trackOpt := IPCmd.Flags().BoolP("track", "t", false, "track IP address info")

	IPCmd.Run = func(cmd *cobra.Command, args []string) {
		if *trackOpt {
			if len(args) > 0 {
				for _, IP := range args {
					showData(IP)
				}
			} else {
				cmd.Println("please provide an IP address")
			}
		}
	}
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := utils.GetDataHTTP(url)

	data := types.IP{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	fmt.Println("DATA FOUND :")

	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE:%s\nPOSTAL :%s\n",
		data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}
