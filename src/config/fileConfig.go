package config

import (
	"log"
	"os"
	"path/filepath"
)

const CLI_CD_CONFIG_FILE string = "appdata.json"

// getAppDataFilePath
func getCliCdFilePath() (string, error) {
	appDataDirPath, err := getCliCdDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(appDataDirPath, CLI_CD_CONFIG_FILE), nil
}

// isCliCdFileExist
func isCliCdConfigFileExist() bool {
	filePath, err := getCliCdFilePath()
	if err != nil {
		return false
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateCliCdFile
func CreateCliCdConfigFile() error {
	if !isCliCdConfigFileExist() {
		if err := CreateCliCdDir(); err != nil {
			log.Fatal(err)
		}

		filePath, err := getCliCdFilePath()
		if err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
