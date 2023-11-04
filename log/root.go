package log

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/knnls/depo-cli/files"
)

func createErrorLogFile() string {
	var filesystem files.FS
	homeDir := filesystem.GetHomeDir()
	cwdFolderName := filesystem.GetCwdFolderName()
	logFolderPath := filepath.Join(homeDir, "depo", "logs")
	currentProjectFolderPath := filesystem.CreateFolder(logFolderPath, cwdFolderName)
	logFilePath := filesystem.CreateFile(currentProjectFolderPath, "error.log")
	return logFilePath
}

func Print(lvl string, msg interface{}) {
	logger := log.Default().With()
	logger.SetPrefix("[📦 depo]")
	logger.SetReportTimestamp(true)
	logger.SetReportCaller(false)
	logLvl := log.ParseLevel(lvl)
	logger.SetLevel(logLvl)

	switch logLvl {
	case log.DebugLevel:
		logger.Debug(msg)
	case log.InfoLevel:
		logger.Info(msg)
	case log.WarnLevel:
		logger.Warn(msg)
	case log.ErrorLevel:
		errorLogFilePath := createErrorLogFile()
		f, _ := os.OpenFile(errorLogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		logger.SetOutput(f)
		logger.SetFormatter(log.JSONFormatter)
		defer f.Close()
		logger.Error(msg)
		logger2 := log.Default().With()
		logger2.SetPrefix("[📦 depo]")
		logger2.SetReportTimestamp(true)
		logger2.SetReportCaller(false)
		logger2.SetLevel(log.ErrorLevel)
		log.ErrorLevelStyle = lipgloss.NewStyle().
			SetString("[ERROR]").
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#fff"))
		log.KeyStyles["err_msg"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
		log.ValueStyles["err_msg"] = lipgloss.NewStyle().Bold(true)
		logger2.Error(msg)

	case log.FatalLevel:
		logger.Fatal(msg)
	}
}
