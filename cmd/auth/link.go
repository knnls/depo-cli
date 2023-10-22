package auth

import (
	"fmt"

	"github.com/knnls/depo-cli/log"
	"github.com/knnls/depo-cli/utils"
	"github.com/spf13/cobra"
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
			} else if len(args[0]) != 36 {
				log.Print("error", "access_token is too short")
			} else {
				log.Print("info", args[0])
			}
		},
	}

	return link
}
