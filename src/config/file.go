package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/atarte/cli-cd/data"
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
func CreateCliCdConfigFile() {
	if !isCliCdConfigFileExist() {
		if err := CreateCliCdDir(); err != nil {
			log.Fatal(err)
		}

		emptyData, err := json.MarshalIndent(data.AppData{}, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		filePath, err := getCliCdFilePath()
		if err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile(filePath, emptyData, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// LoadCkiCdConfigFile
func LoadCliCdConfigFile() data.AppData {
	if !isCliCdConfigFileExist() {
		log.Fatal("Config file should exist at this point. Cannot 'LoadCliCdConfigFile'. Relauch the program.")
	}

	filePath, err := getCliCdFilePath()
	if err != nil {
		log.Fatal(err)
	}

	byteFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var appdata data.AppData
	if err := json.Unmarshal(byteFile, &appdata); err != nil {
		log.Fatal(err)
	}

	return appdata
}

// UpdateCliCdConfigFile
func UpdateCliCdConfigFile(data data.AppData) {
	if !isCliCdConfigFileExist() {
		log.Fatal("Config file should exist at this point. Cannot 'UpdateCliCdConfigFile'. Relauch the program.")
	}

	dataJson, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	filePath, err := getCliCdFilePath()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filePath, dataJson, 0644); err != nil {
		log.Fatal(err)
	}
}
