package cmd

import (
	"context"
	"fmt"

	cli_auth "github.com/knnls/depo-cli/cmd/auth"
	"github.com/knnls/depo-cli/files"
	"github.com/knnls/depo-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	var filesystem files.FS
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
	homeDir := filesystem.GetHomeDir()
	depoFolder := "depo"
	depoFolderPath := filesystem.CreateFolder(homeDir, depoFolder)
	filesystem.CreateFolder(depoFolderPath, "logs")

	// Use depoFolderPath as the config path
	viper.AddConfigPath(depoFolderPath)
	viper.SetConfigName("config") // Specify the config file name here
	viper.SetConfigType("json")

	// Set default values before reading the configuration file
	viper.SetDefault("api_key", "")
	viper.SetDefault("user.username", "")
	viper.SetDefault("user.id", "")

	// Read the configuration file after setting default values
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Creating a new one with default values...")
			createDefaultConfigFile(depoFolderPath) // Create a new configuration file
		} else {
			fmt.Printf("Error reading config file: %s\n", err)
		}
	}
	rootCmd.AddCommand(cli_auth.LinkCmd())
	return rootCmd.ExecuteContext(context.Background())
}

func createDefaultConfigFile(path string) {
	// Set default values directly using viper.SetDefault
	viper.SetDefault("api_key", "")
	viper.SetDefault("user.username", "")
	viper.SetDefault("user.id", "")

	viper.SetConfigFile(path + "/config.json")
	viper.SetConfigType("json")
	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		fmt.Printf("Error creating config file: %s\n", err)
	}
}
