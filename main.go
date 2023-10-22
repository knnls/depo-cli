/*
Copyright © 2023 KNNLS <i.am@madebyknnls.com>
*/
package main

import (
	"github.com/knnls/depo-cli/cmd"
	"github.com/knnls/depo-cli/files"
)

func main() {
	var filesystem files.FS

	homeDir := filesystem.GetHomeDir()
	depoFolder := "depo"
	depoFolderPath := filesystem.CreateFolder(homeDir, depoFolder)
	filesystem.CreateFolder(depoFolderPath, "config")
	filesystem.CreateFolder(depoFolderPath, "logs")
	cmd.Execute()
}
