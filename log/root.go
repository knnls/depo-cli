package log

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"github.com/knnls/depo-cli/files"
// 	log "github.com/sirupsen/logrus"
// )

// type ILog interface {
// 	Logger(logLevel string, message string, fields interface{})
// }

// type Error struct {
// 	ILog
// }

// type Warn struct {
// 	ILog
// }

// type Info struct {
// 	ILog
// }

// type Hook struct{}

// func getLogLvl(logLevel string) log.Level {
// 	lvl, _ := log.ParseLevel(logLevel)

// 	return lvl
// }

// func createAppLogger() *log.Logger {
// 	logger := log.New()
// 	logger.SetFormatter(&log.TextFormatter{
// 		FullTimestamp: true,
// 		DisableColors: false,
// 	})
// 	return logger
// }

// func CreateErrorLogFile() {
// 	var filesystem files.FS
// 	homeDir := filesystem.GetHomeDir()
// 	cwdFolderName := filesystem.GetCwdFolderName()
// 	logFolderPath := filepath.Join(homeDir, ".depocli", "logs")
// 	currentProjectFolderPath := filesystem.CreateFolder(logFolderPath, cwdFolderName)
// 	filesystem.CreateFile(currentProjectFolderPath, fmt.Sprintf("%s.error.log", ".depocli"))
// }

// func getLogfile() string {
// 	var filesystem files.FS
// 	homeDir := filesystem.GetHomeDir()
// 	cwdFolderName := filesystem.GetCwdFolderName()
// 	logFolderPath := filepath.Join(homeDir, ".depocli", "logs")
// 	currentProjectFolderPath := filesystem.CreateFolder(logFolderPath, cwdFolderName)
// 	errorFilePath := filepath.Join(currentProjectFolderPath, ".depocli.error.log")
// 	return errorFilePath
// }

// func Log(logLevel string, message string, fields *log.Fields) {
// 	logger := createAppLogger()
// 	lvl := getLogLvl(logLevel)

// 	logger.Out = os.Stdout

// 	// if logLevel == "error" {
// 	// 	errorFilePath := getLogfile()
// 	// 	file, _ := os.OpenFile(errorFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	// 	logger.Out = file
// 	// }

// 	logger.SetLevel(lvl)

// 	if fields != nil {
// 		logger.WithFields(*fields)
// 	}

// 	logger.Error(message)
// }
