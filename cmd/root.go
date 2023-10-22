package cmd

import (
	"context"
	"fmt"

	"github.com/knnls/depo-cli/cmd/auth"
	"github.com/knnls/depo-cli/utils"
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Version: utils.VERSION,
		Use:     "depo",
		Long:    fmt.Sprintf("%s\n%s%s%sdepo-cli [%s%s]", utils.ASCII_TEXT, utils.Bold, utils.Italics, utils.Underline, utils.VERSION, utils.Reset),
		Short:   fmt.Sprintf("%s\n\n%s%s%sdepo-cli [%s%s]", utils.ASCII_TEXT, utils.Bold, utils.Italics, utils.Underline, utils.VERSION, utils.Reset),
		Example: "depo",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	rootCmd.AddCommand(auth.LinkCmd())
	return rootCmd.ExecuteContext(context.Background())
}
