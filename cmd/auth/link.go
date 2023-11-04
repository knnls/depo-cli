package cli_auth

import (
	"fmt"

	"github.com/knnls/depo-cli/api"
	"github.com/knnls/depo-cli/log"
	"github.com/knnls/depo-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LinkCmd() *cobra.Command {
	link := &cobra.Command{
		Use:     "link",
		Short:   "link with your depo account",
		Long:    "link with your depo account",
		Example: "depo link <access_token>",
		Aliases: []string{"l", "link"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Print("error", fmt.Sprintf("command failed, try `%s%sdepo link <access_token>%s`", utils.Bold, utils.Underline, utils.Reset))
			} else if len(args[0]) != 32 {
				log.Print("error", "access_token is too short")
			} else {
				var a api.API

				responseMap := a.MakeAuthRequest(args[0])
				fmt.Println(responseMap)

				viper.Set("token", responseMap["token"])
				viper.Set("api_key", responseMap["access_token"])
				viper.Set("user.id", responseMap["id"])
				viper.Set("user.username", responseMap["username"])

				if err := viper.WriteConfig(); err != nil {
					fmt.Printf("Error writing config file: %s\n", err)
				}

				metadata := utils.GetDeviceMetadata()

				a.SubmitMetadata(metadata)
			}
		},
	}

	return link
}
