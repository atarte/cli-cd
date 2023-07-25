package config

import (
	"log"
	"os"
	"path/filepath"
)

const CLI_CD_CONFIG_FILE string = "appdata.json"

// getAppDataFilePath
func GetCliCdFilePath() (string, error) {
	appDataDirPath, err := getCliCdDirPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(appDataDirPath, CLI_CD_CONFIG_FILE), nil
}

// isCliCdFileExist
func IsCliCdConfigFileExist() bool {
	filePath, err := GetCliCdFilePath()
	if err != nil {
		return false
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateCliCdFile
func CreateCliCdConfigFile() {
	if !IsCliCdConfigFileExist() {
		if err := CreateCliCdDir(); err != nil {
			log.Fatal(err)
		}

		// emptyData, err := json.MarshalIndent(data.AppData{}, "", "    ")
		// emptyData, err := json.MarshalIndent(, "", "    ")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		filePath, err := GetCliCdFilePath()
		if err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
