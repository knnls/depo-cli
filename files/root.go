package files

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/knnls/depo-cli/utils"
)

type FS struct{}

func (fs *FS) GetCwd() string {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("error getting current directory %v\n", err)
	}

	return cwd
}

func (fs *FS) GetHomeDir() string {
	currentUser, err := user.Current()

	if err != nil {
		log.Fatalf("error getting current user %v\n", err)
	}

	homeDir := currentUser.HomeDir

	return homeDir
}

func (fs *FS) GetCwdFolderName() string {
	cwd := fs.GetCwd()
	folderName := filepath.Base(cwd)
	return folderName
}

func (fs *FS) CreateFolder(dirPath string, folderName string) string {
	folderPath := filepath.Join(dirPath, folderName)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0755); err != nil {
			log.Fatalf(
				"error creating folder `%s%s%s%s` in directory: `%s%s%s%s` %v\n",
				utils.Bold,
				utils.Underline,
				folderName,
				utils.Reset,
				utils.Bold,
				utils.Underline,
				dirPath,
				utils.Reset,
				err,
			)
		}
		return folderPath
	} else {
		return folderPath
	}
}

func (fs *FS) CreateFile(dirPath string, fileName string) string {
	filePath := filepath.Join(dirPath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)

		if err != nil {
			log.Fatalf(
				"error creating file in `%s%s%s%s` in directory: `%s%s%s%s` %v\n",
				utils.Bold,
				utils.Underline,
				fileName,
				utils.Reset,
				utils.Bold,
				utils.Underline,
				dirPath,
				utils.Reset,
				err,
			)
		}
		defer file.Close()
		return filePath
	} else {
		return filePath
	}
}
